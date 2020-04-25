package templateEngine

import (
	"bytes"
	"fmt"
	"html/template"
)

type ITemplate interface {
	ParseFiles(...string) (*template.Template, error)
}

type TemplateEngine struct {
	template ITemplate
}

func New() *TemplateEngine {
	t := template.New("short_term.html")
	return &TemplateEngine{
		template: t,
	}
}

func (te *TemplateEngine) CreateHTML(templateFilePath string, data interface{}) (string, error) {
	t, err := te.template.ParseFiles(templateFilePath)
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	fmt.Printf("%+v %+v\n", t, data)
	if err = t.Execute(buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}
