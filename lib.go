package i18n

import (
	"github.com/laipz8200/i18n/pkg/i18n"
)

var std *i18n.I18n

func init() {
	std = NewI18n()
}

func Lang(language string) *i18n.I18n {
	return std.Lang(language)
}

// NewI18n
func NewI18n() *i18n.I18n {
	return i18n.NewI18n()
}

// Sprintf
func Sprintf(format string, a ...any) string {
	return std.Sprintf(format, a...)
}

// Sprintln
func Sprintln(a ...any) string {
	return Sprintf("%v\n", a...)
}

// SetLanguage
func SetLanguage(language string) {
	std.SetLanguage(language)
}
