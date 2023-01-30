package i18n

var _ Logger = (*mockLogger)(nil)

type mockLogger struct {
	printf int
}

// Printf implements Logger
func (l *mockLogger) Printf(format string, a ...any) {
	l.printf++
}
