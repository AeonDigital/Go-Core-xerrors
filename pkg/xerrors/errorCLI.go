package xerrors

import "fmt"

// IErrorCLI defines the structural behavioral contract for light, terminal-friendly errors.
type IErrorCLI interface {
	// error implements the native Go error interface.
	error

	// SetMessage overrides the current technical diagnostic text and the current end-user friendly instruction.
	SetMessage(format string, args ...any) IErrorCLI

	// SetDevMessage overrides the current technical diagnostic text.
	SetDevMessage(format string, args ...any) IErrorCLI

	// SetUserMessage overrides the current end-user friendly instruction.
	SetUserMessage(format string, args ...any) IErrorCLI

	// AppendDevMessage concatenates a text payload into the current developer message.
	AppendDevMessage(format string, args ...any) IErrorCLI

	// AppendUserMessage concatenates a text payload into the current end-user message.
	AppendUserMessage(format string, args ...any) IErrorCLI

	// AppendLNDevMessage concatenates a text payload appending an explicit newline character at the end.
	AppendLNDevMessage(format string, args ...any) IErrorCLI

	// AppendLNUserMessage concatenates a text payload appending an explicit newline character at the end.
	AppendLNUserMessage(format string, args ...any) IErrorCLI

	// ClearDevMessage purges the developer message content, resetting it to empty string.
	ClearDevMessage() IErrorCLI

	// ClearUserMessage purges the end-user message content, resetting it to empty string.
	ClearUserMessage() IErrorCLI

	// WithDepth forces the engine to recalculate the runtime trace stack location using an offset.
	WithDepth(additionalDepth int) IErrorCLI

	// GetFunction returns the qualified target name tracking the package and functional scope.
	GetFunction() string

	// GetDevMessage returns the technical diagnostics text string.
	GetDevMessage() string

	// GetUserMessage returns the actionable text instruction designed for human end-users.
	GetUserMessage() string

	// HasDevMessage verifies if a technical message payload has been populated.
	HasDevMessage() bool

	// HasUserMessage verifies if a human-friendly instruction payload has been populated.
	HasUserMessage() bool

	// HasErrors checks if any descriptive message fields contain active content tracking failures.
	HasErrors() bool
}

//
//
//

// typedIErrorCLI implements IErrorCLI with immutable, copy-on-write semantics.
type typedIErrorCLI struct {
	function string
	devMsg   string
	userMsg  string
}

// NewErrorCLI creates a new CLI-oriented error with automatic location tracking.
//
// It captures the immediate caller metadata from the runtime stack so the new
// error can report the originating function when rendered.
func NewErrorCLI() IErrorCLI {
	return &typedIErrorCLI{
		function: TraceCallerLocation(1),
	}
}

// NewErrorCLIWithFunc creates a CLI-oriented error with an explicit function location.
//
// Arguments:
//   - functionName: The static text representation of the target scope location (for example, "pkgname::FunctionName").
func NewErrorCLIWithFunc(functionName string) IErrorCLI {
	return &typedIErrorCLI{
		function: functionName,
	}
}

// Error returns the technical diagnostic string expected by the Go error interface.
//
// It formats the tracked function location together with the current developer
// message so the error can be printed or logged directly.
func (e *typedIErrorCLI) Error() string {
	return fmt.Sprintf("[FUNC: %s][MSG: %s]", e.function, e.devMsg)
}

// SetMessage replaces both the developer and user-facing messages.
//
// Arguments:
//   - format: Standard formatting template used to build the new message.
//   - args: Variable payloads to inject into the formatting placeholders.
func (e *typedIErrorCLI) SetMessage(format string, args ...any) IErrorCLI {
	msg := fmt.Sprintf(format, args...)
	return &typedIErrorCLI{
		function: e.function,
		devMsg:   msg,
		userMsg:  msg,
	}
}

// SetDevMessage replaces the developer-facing diagnostic text.
//
// Arguments:
//   - format: Standard formatting template used to build the new developer message.
//   - args: Variable payloads to inject into the formatting placeholders.
func (e *typedIErrorCLI) SetDevMessage(format string, args ...any) IErrorCLI {
	return &typedIErrorCLI{
		function: e.function,
		devMsg:   fmt.Sprintf(format, args...),
		userMsg:  e.userMsg,
	}
}

// SetUserMessage replaces the end-user facing instruction text.
//
// Arguments:
//   - format: Standard formatting template used to build the new user message.
//   - args: Variable payloads to inject into the formatting placeholders.
func (e *typedIErrorCLI) SetUserMessage(format string, args ...any) IErrorCLI {
	return &typedIErrorCLI{
		function: e.function,
		devMsg:   e.devMsg,
		userMsg:  fmt.Sprintf(format, args...),
	}
}

// AppendDevMessage appends a payload to the current developer message.
//
// Arguments:
//   - format: Standard formatting template used to build the appended text.
//   - args: Variable payloads to inject into the formatting placeholders.
func (e *typedIErrorCLI) AppendDevMessage(format string, args ...any) IErrorCLI {
	payload := fmt.Sprintf(format, args...)
	if e.devMsg != "" {
		payload = e.devMsg + payload
	}
	return &typedIErrorCLI{
		function: e.function,
		devMsg:   payload,
		userMsg:  e.userMsg,
	}
}

// AppendUserMessage appends a payload to the current user-facing message.
//
// Arguments:
//   - format: Standard formatting template used to build the appended text.
//   - args: Variable payloads to inject into the formatting placeholders.
func (e *typedIErrorCLI) AppendUserMessage(format string, args ...any) IErrorCLI {
	payload := fmt.Sprintf(format, args...)
	if e.userMsg != "" {
		payload = e.userMsg + payload
	}
	return &typedIErrorCLI{
		function: e.function,
		devMsg:   e.devMsg,
		userMsg:  payload,
	}
}

// AppendLNDevMessage appends a payload and a trailing newline to the developer message.
//
// Arguments:
//   - format: Standard formatting template used to build the appended text.
//   - args: Variable payloads to inject into the formatting placeholders.
func (e *typedIErrorCLI) AppendLNDevMessage(format string, args ...any) IErrorCLI {
	payload := fmt.Sprintf(format, args...) + "\n"
	if e.devMsg != "" {
		payload = e.devMsg + payload
	}
	return &typedIErrorCLI{
		function: e.function,
		devMsg:   payload,
		userMsg:  e.userMsg,
	}
}

// AppendLNUserMessage appends a payload and a trailing newline to the user message.
//
// Arguments:
//   - format: Standard formatting template used to build the appended text.
//   - args: Variable payloads to inject into the formatting placeholders.
func (e *typedIErrorCLI) AppendLNUserMessage(format string, args ...any) IErrorCLI {
	payload := fmt.Sprintf(format, args...) + "\n"
	if e.userMsg != "" {
		payload = e.userMsg + payload
	}
	return &typedIErrorCLI{
		function: e.function,
		devMsg:   e.devMsg,
		userMsg:  payload,
	}
}

// ClearDevMessage removes the developer-facing message content.
//
// It returns a new error instance with the developer message reset to an empty
// string while preserving the other state.
func (e *typedIErrorCLI) ClearDevMessage() IErrorCLI {
	return &typedIErrorCLI{
		function: e.function,
		devMsg:   "",
		userMsg:  e.userMsg,
	}
}

// ClearUserMessage removes the user-facing message content.
//
// It returns a new error instance with the user message reset to an empty string
// while preserving the other state.
func (e *typedIErrorCLI) ClearUserMessage() IErrorCLI {
	return &typedIErrorCLI{
		function: e.function,
		devMsg:   e.devMsg,
		userMsg:  "",
	}
}

// WithDepth recalculates the tracked function location using an offset.
//
// This is useful when the error is created from helper factories or wrappers and
// the original caller location must be shifted upward in the runtime stack.
//
// Arguments:
//   - additionalDepth: Positive stack depth modifier index layer count.
func (e *typedIErrorCLI) WithDepth(additionalDepth int) IErrorCLI {
	// Protect against zero or negative depth alterations
	if additionalDepth <= 0 {
		return e
	}
	return &typedIErrorCLI{
		// +1 accounts for moving away from this wrapper call context scope inside the lib
		function: TraceCallerLocation(1 + additionalDepth),
		devMsg:   e.devMsg,
		userMsg:  e.userMsg,
	}
}

// GetFunction returns the tracked package-qualified function location.
//
// It exposes the runtime scope identifier associated with the error instance.
func (e *typedIErrorCLI) GetFunction() string {
	return e.function
}

// GetDevMessage returns the developer-facing diagnostic text.
//
// It exposes the technical message payload recorded for the error instance.
func (e *typedIErrorCLI) GetDevMessage() string {
	return e.devMsg
}

// GetUserMessage returns the end-user facing instruction text.
//
// It exposes the human-readable guidance recorded for the error instance.
func (e *typedIErrorCLI) GetUserMessage() string {
	return e.userMsg
}

// HasDevMessage reports whether the developer message has been populated.
//
// It returns true when the error carries active technical diagnostics.
func (e *typedIErrorCLI) HasDevMessage() bool {
	return e.devMsg != ""
}

// HasUserMessage reports whether the user message has been populated.
//
// It returns true when the error carries active end-user guidance.
func (e *typedIErrorCLI) HasUserMessage() bool {
	return e.userMsg != ""
}

// HasErrors reports whether the error carries any descriptive payload.
//
// It returns true when either the developer message or the user message is set.
func (e *typedIErrorCLI) HasErrors() bool {
	return e.devMsg != "" || e.userMsg != ""
}
