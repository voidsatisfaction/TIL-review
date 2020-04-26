package templateengine

import (
	"bytes"
	"fmt"
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
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	htmlString := fmt.Sprintf("%s%s", mime, buf.String())
	return htmlString, nil
}
