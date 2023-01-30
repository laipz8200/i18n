package i18n

import (
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

type i18n struct {
	language string
	textMap  map[string]map[string]string
	dir      string
	logger   Logger
}

func (i *i18n) loadFile() {
	i.textMap = map[string]map[string]string{}

	files, err := os.ReadDir(i.dir)
	if err != nil {
		i.logger.Printf("[Warning] Cannot read dir %q: %v", i.dir, err)
		return
	}

	for _, file := range files {
		if !file.IsDir() {
			bytes, err := os.ReadFile(i.dir + "/" + file.Name())
			if err != nil {
				i.logger.Printf("[Warning] Cannot read file %q: %v", file.Name(), err)
				return
			}

			data := make(map[string]string)
			err = yaml.Unmarshal(bytes, &data)
			if err != nil {
				i.logger.Printf("[Warning] Cannot unmarshal yaml file %q: %v", file.Name(), err)
			}

			sp := strings.Split(file.Name(), ".")
			name := strings.Join(sp[:len(sp)-1], ".")
			i.textMap[name] = data
		}
	}
}

func (i *i18n) copy() *i18n {
	return &i18n{
		language: i.language,
		textMap:  i.textMap,
		dir:      i.dir,
		logger:   i.logger,
	}
}

func (i *i18n) Lang(language string) *i18n {
	newObj := i.copy()
	newObj.language = language
	return newObj
}

func (i *i18n) Sprintf(format string, a ...any) string {
	if i.textMap == nil {
		i.loadFile()
	}

	objs := make([]any, len(a))
	copy(objs, a)
	if transSet, ok := i.textMap[i.language]; ok {
		if target, ok := transSet[format]; ok {
			format = target
		}

		for i, obj := range objs {
			if target, ok := transSet[fmt.Sprintf("%v", obj)]; ok {
				objs[i] = target
			}
		}
	}
	return fmt.Sprintf(format, objs...)
}

func (i *i18n) Sprintln(a ...any) string {
	return i.Sprintf("%v\n", a...)
}
