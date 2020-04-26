package templateengine

import (
	"bytes"
	"html/template"
)

type TemplateEngine struct {
}

func New() *TemplateEngine {
	return &TemplateEngine{}
}

/*
	CreateHTML is ...
*/
func (te *TemplateEngine) CreateHTML(templateFilePathList []string, data interface{}) (string, error) {
	t, err := template.ParseFiles(templateFilePathList...)
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return "", err
	}
	htmlString := buf.String()
	return htmlString, nil
}
