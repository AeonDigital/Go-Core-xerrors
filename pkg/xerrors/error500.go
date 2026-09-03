package xerrors

// IError500 standardizes server-side, operational diagnostic payloads.
type IError500 interface {
	// CTX extracts the entry point operational flow boundary metadata.
	CTX() ErrorCode

	// Code returns the categorical failure domain classification.
	Code() ErrorCode

	// Component pinpoints the reflective architectural package or execution function path.
	Component() string

	// WithCallerSkip dynamically shifts the stack frame runtime collection depth.
	WithCallerSkip(skip int) IError500

	// WithArgs appends dynamic contextual payloads to map sequentially into metadata extraTags.
	WithArgs(args ...any) IError500

	// Message extracts the human-readable summary detailing the specific failure.
	Message() string

	// Info returns secondary raw debugging contextual payloads.
	Info() string

	// error ensures native integration with Go standard library error semantics.
	error
}

// error500Adapter wraps the internal private engine to satisfy the public IError500 contract.
type error500Adapter struct {
	*xError
}

// NewError500 creates an operational IError500 instance with underlying error context.
//
// It preserves the supplied context, error code, cause, summary message, and
// additional debugging payload while enabling component tracing for the error.
func NewError500(
	errCTX ErrorCode,
	errCode ErrorCode,
	err error,
	message string,
	info string,
) IError500 {
	return &error500Adapter{
		xError: &xError{
			contextCode:   errCTX,
			errorCode:     errCode,
			underlyingErr: err,
			message:       message,
			info:          info,
			isOperational: true, // Forces system telemetry behavior
		},
	}
}

// CTX returns the assigned operational context tracking identifier.
//
// It exposes the package-level boundary that classifies where the failure was
// produced inside the execution flow.
func (a *error500Adapter) CTX() ErrorCode {
	return a.contextCode
}

// Code returns the structured domain classification code assigned to the error.
//
// It exposes the specific failure code that complements the operational context.
func (a *error500Adapter) Code() ErrorCode {
	return a.errorCode
}

// Component returns the reflected runtime namespace where the failure originated.
//
// It resolves the caller component through the internal stack inspection logic
// and exposes it as a human-readable execution path.
func (a *error500Adapter) Component() string {
	return a.getComponent()
}

// WithCallerSkip returns a cloned error with an adjusted stack inspection depth.
//
// It is useful when the error is wrapped by helper functions and the caller
// frame must be shifted to preserve the correct component attribution.
func (a *error500Adapter) WithCallerSkip(skip int) IError500 {
	// Calls internal deep-copy mechanism to shield concurrent operations
	clonedEngine := a.withCallerSkip(skip)
	return &error500Adapter{xError: clonedEngine}
}

// WithArgs clones the error and injects additional operational context arguments.
//
// It preserves the original state and appends the supplied values so they can
// be rendered into the formatted output when metadata tags are present.
func (a *error500Adapter) WithArgs(args ...any) IError500 {
	if len(args) == 0 {
		return a
	}

	a.mu.RLock()
	// Explicitly copy only data fields to prevent copying the sync.RWMutex value
	clonedEngine := &xError{
		contextCode:   a.contextCode,
		errorCode:     a.errorCode,
		underlyingErr: a.underlyingErr,
		message:       a.message,
		info:          a.info,
		isOperational: a.isOperational,
	}
	a.mu.RUnlock()

	// Deep copy arguments payload to guarantee concurrency safety
	clonedEngine.arguments = make([]any, len(args))
	copy(clonedEngine.arguments, args)

	return &error500Adapter{xError: clonedEngine}
}

// Message returns the human-readable operational summary associated with the error.
//
// It resolves the stored message, falling back to any registered metadata entry
// when no explicit message was provided when the error was created.
func (a *error500Adapter) Message() string {
	return a.resolveMessage()
}

// Info returns the unstructured secondary metadata payload attached to the error.
//
// It exposes the free-form debugging context that can be used for logs, traces,
// or operational troubleshooting.
func (a *error500Adapter) Info() string {
	return a.info
}

// Error returns the final formatted error string for display and logging.
//
// It delegates to the shared formatting engine so the rendered output remains
// consistent across operational diagnostics.
func (a *error500Adapter) Error() string {
	return a.format()
}

// Unwrap exposes the wrapped underlying error for standard library compatibility.
//
// It enables callers to use errors.Is and errors.As with the operational error.
func (a *error500Adapter) Unwrap() error {
	return a.underlyingErr
}
