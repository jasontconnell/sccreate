package process

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"github.com/google/uuid"
)

func getCleanName(name string) string {
	prefix := ""
	if strings.IndexAny(name, "0123456789") == 0 {
		prefix = "_"
	}
	return prefix + strings.Replace(strings.Replace(strings.Title(name), "-", "", -1), " ", "", -1)
}

func processInlineTemplate(code string, model interface{}) (string, error) {
	var value string

	ftmpl := template.New("Template")
	ftmp, err := ftmpl.Parse(code)
	if err != nil {
		return "", fmt.Errorf("parsing string %s: %w", code, err)
	}

	fbuf := new(bytes.Buffer)
	err = ftmp.Execute(fbuf, model)
	if err != nil {
		return "", fmt.Errorf("executing template inline %s: %w", code, err)
	}

	value = string(fbuf.Bytes())
	return value, nil
}

func sitecoreStyleGuid(uid uuid.UUID) string {
	return fmt.Sprintf("{%s}", strings.ToLower(uid.String()))
}
