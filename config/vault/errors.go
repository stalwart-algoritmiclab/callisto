package vault

import "strings"

// error messages
const (
	connectionError = "connection_error"
	readError       = "read_error"
	writeError      = "write_error"
)

type (
	// ClientError - define package error type for vault client initialization.
	ClientError struct {
		What string
	}

	// ReadError - define read data error type.
	ReadError struct {
		What string
	}

	// WriteError - define write data error type.
	WriteError struct {
		What string
	}

	// ValidationError - define vault client validation error.
	ValidationError struct {
		What []string
	}
)

// Error - error interface implementation for ClientError.
func (e ClientError) Error() string {
	if e.What == "" {
		return connectionError
	}
	return connectionError + ": " + e.What
}

// Error - error interface implementation for ValidationError.
func (e ValidationError) Error() string {
	return strings.Join(e.What, ", ")
}

// Error - error interface implementation for ReadError.
func (e ReadError) Error() string {
	if e.What == "" {
		return readError
	}
	return readError + ": " + e.What
}

// Error - error interface implementation for WriteError.
func (e WriteError) Error() string {
	if e.What == "" {
		return writeError
	}
	return writeError + ": " + e.What
}
