package xerrors

// MetaMessage stores fallback message text, field-rule metadata, and extra tag names.
//
// It provides the package with reusable, structured formatting details that can
// be attached to error codes for predictable rendering.
type MetaMessage struct {
	message   string
	fieldRule string
	extraTags []string
}

// NewMetaMessage creates and returns a populated MetaMessage instance.
//
// It exposes a simple constructor for registering fallback metadata that can be
// reused by the package when an error code requires a standard message layout.
func NewMetaMessage(
	message string,
	fieldRule string,
	extraTags []string,
) MetaMessage {
	return MetaMessage{
		message:   message,
		fieldRule: fieldRule,
		extraTags: extraTags,
	}
}
