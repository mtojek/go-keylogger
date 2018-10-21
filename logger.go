package keylogger

// Logger exposes an interface to collect errors.
type Logger interface {
	Error(string)
}
