package generate

import (
	"bytes"
	"html/template"
)

func getTemplateString(t *template.Template, name string, payload any) (string, error) {
	buf := &bytes.Buffer{}
	err := t.ExecuteTemplate(buf, name, payload)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
