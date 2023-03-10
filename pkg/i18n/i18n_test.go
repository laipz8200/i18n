package i18n

import (
	"log"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_i18n_Sprintf(t *testing.T) {
	type fields struct {
		language string
		textMap  map[string]map[string]string
		dir      string
		logger   *mockLogger
	}
	type args struct {
		format string
		a      []any
		lang   string
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		want       string
		wantPrintf int
	}{
		{
			name: "case 1",
			fields: fields{
				language: "",
				dir:      "fixtures/i18n",
				logger:   &mockLogger{},
			},
			args:       args{format: "%s", a: []any{"hello, world!"}, lang: "zh-cns"},
			want:       "你好，世界！",
			wantPrintf: 0,
		},
		{
			name: "format number",
			fields: fields{
				language: "",
				dir:      "fixtures/i18n",
				logger:   &mockLogger{},
			},
			args:       args{format: "found %d errors", a: []any{20}, lang: "zh-cns"},
			want:       "找到 20 个错误",
			wantPrintf: 0,
		},
		{
			name: "format string",
			fields: fields{
				language: "",
				dir:      "fixtures/i18n",
				logger:   &mockLogger{},
			},
			args:       args{format: "hello, %s", a: []any{"xiaoming"}, lang: "zh-cns"},
			want:       "你好，小明",
			wantPrintf: 0,
		},
		{
			name: "invalid dir",
			fields: fields{
				language: "",
				dir:      "wrong dir path",
				logger:   &mockLogger{},
			},
			args:       args{},
			want:       "",
			wantPrintf: 1,
		},
		{
			name: "invalid file",
			fields: fields{
				language: "",
				dir:      "fixtures/wrong",
				logger:   &mockLogger{},
			},
			args:       args{},
			want:       "",
			wantPrintf: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &I18n{
				language: tt.fields.language,
				textMap:  tt.fields.textMap,
				dir:      tt.fields.dir,
				logger:   tt.fields.logger,
			}
			got := i.Lang(tt.args.lang).Sprintf(tt.args.format, tt.args.a...)

			asserts := assert.New(t)
			asserts.Equal(tt.want, got)
			asserts.Equal(tt.wantPrintf, tt.fields.logger.printf)
		})
	}
}

func Test_i18n_copy(t *testing.T) {
	type fields struct {
		language string
		textMap  map[string]map[string]string
		dir      string
		logger   Logger
	}
	tests := []struct {
		name   string
		fields fields
		want   *I18n
	}{
		{
			name: "test copy",
			fields: fields{
				language: "default language",
				textMap: map[string]map[string]string{
					"language 1": {
						"text 1": "translate 1",
						"text 2": "translate 2",
					},
				},
				dir:    "some dir",
				logger: &log.Logger{},
			},
			want: &I18n{
				language: "default language",
				textMap: map[string]map[string]string{
					"language 1": {
						"text 1": "translate 1",
						"text 2": "translate 2",
					},
				},
				dir:    "some dir",
				logger: &log.Logger{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &I18n{
				language: tt.fields.language,
				textMap:  tt.fields.textMap,
				dir:      tt.fields.dir,
				logger:   tt.fields.logger,
			}
			if got := i.copy(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("i18n.copy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_i18n_Sprintln(t *testing.T) {
	type fields struct {
		language string
		textMap  map[string]map[string]string
		dir      string
		logger   Logger
	}
	type args struct {
		a    []any
		lang string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "case 1",
			fields: fields{
				dir: "fixtures/i18n",
			},
			args: args{
				a:    []any{"hello, world!"},
				lang: "zh-cns",
			},
			want: "你好，世界！\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &I18n{
				language: tt.fields.language,
				textMap:  tt.fields.textMap,
				dir:      tt.fields.dir,
				logger:   tt.fields.logger,
			}
			if got := i.Lang(tt.args.lang).Sprintln(tt.args.a...); got != tt.want {
				t.Errorf("i18n.Sprintln() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestI18n_SetLanguage(t *testing.T) {
	type fields struct {
		language string
		textMap  map[string]map[string]string
		dir      string
		logger   Logger
	}
	type args struct {
		language string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "case 1",
			fields: fields{},
			args: args{
				language: "zh-cn",
			},
		},
		{
			name:   "case 2",
			fields: fields{},
			args: args{
				language: "en-us",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &I18n{
				language: tt.fields.language,
				textMap:  tt.fields.textMap,
				dir:      tt.fields.dir,
				logger:   tt.fields.logger,
			}
			i.SetLanguage(tt.args.language)
			assert.Equal(t, tt.args.language, i.language)
		})
	}
}

func TestNewI18n(t *testing.T) {
	tests := []struct {
		name string
		want *I18n
	}{
		{
			name: "case 1",
			want: &I18n{
				language: "",
				textMap:  nil,
				dir:      "i18n",
				logger:   log.Default(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewI18n(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewI18n() = %v, want %v", got, tt.want)
			}
		})
	}
}
