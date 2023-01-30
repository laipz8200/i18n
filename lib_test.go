package i18n

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLang(t *testing.T) {
	type args struct {
		language string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case 1",
			args: args{
				language: "zh-cn",
			},
			want: "zh-cn",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Lang(tt.args.language)
			want := NewI18n()
			want.language = tt.want
			assert.Equal(t, want, got)
		})
	}
}

func TestSprintln(t *testing.T) {
	type args struct {
		a []any
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case 1",
			args: args{
				a: []any{"hello, world"},
			},
			want: "hello, world\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Sprintln(tt.args.a...)
			assert.Equal(t, tt.want, got)
		})
	}
}
