package i18n

import (
	"log"
)

const DEFAULT_DIR = "i18n"

var std *i18n

func init() {
	std = NewI18n()
}

func Lang(language string) *i18n {
	return std.Lang(language)
}

// NewI18n
func NewI18n() *i18n {
	return &i18n{
		dir:    DEFAULT_DIR,
		logger: log.Default(),
	}
}

// Sprintf
func Sprintf(format string, a ...any) string {
	return std.Sprintf(format, a...)
}

// Sprintln
func Sprintln(a ...any) string {
	return Sprintf("%v\n", a...)
}
