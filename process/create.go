package process

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/jasontconnell/sitecore/data"
)

func CreateCodeFiles(node data.ItemNode, style, markupPath, backendPath, namespace string, templates []Template) error {
	err := os.MkdirAll(markupPath, os.ModePerm)
	if err != nil {
		return err
	}

	err = os.Mkdir(backendPath, os.ModePerm)
	if err != nil {
		return err
	}

	var allerr error
	tmodel := TemplateModel{Name: node.GetName(), CleanName: getCleanName(node.GetName()), Namespace: namespace}
	for _, tmp := range templates {
		filename, err := processInlineTemplate(tmp.OutputFilename, tmodel)
		if err != nil {
			allerr = errors.Join(allerr, fmt.Errorf("couldn't execute template %s. %w", tmp.OutputFilename, err))
			continue
		}

		outpath := markupPath
		if tmp.Type == "backend" {
			outpath = backendPath
		}

		err = processFileTemplate(tmp.TemplateFilename, filepath.Join(outpath, filename), tmodel)
	}

	return allerr
}

func processFileTemplate(filename string, outputPath string, model TemplateModel) error {
	tmpl, err := template.New(filename).ParseFiles(filename)
	if err != nil {
		return fmt.Errorf("error parsing template %s: %w", filename, err)
	}
	buffer := new(bytes.Buffer)
	_, templateName := filepath.Split(filename)
	err = tmpl.ExecuteTemplate(buffer, templateName, model)

	if err != nil {
		return fmt.Errorf("error executing template %s: %w", filename, err)
	}

	return os.WriteFile(outputPath, buffer.Bytes(), os.ModePerm)
}

func processInlineTemplate(code string, model TemplateModel) (string, error) {
	var value string

	ftmpl := template.New("Template")
	ftmp, err := ftmpl.Parse(code)
	if err != nil {
		return "", fmt.Errorf("parsing string %s: %w", code, err)
	}

	fbuf := new(bytes.Buffer)
	err = ftmp.Execute(fbuf, model)
	if err != nil {
		return "", fmt.Errorf("executing template inline %s %s: %w", code, model.Name, err)
	}

	value = string(fbuf.Bytes())
	return value, nil
}

func getCleanName(name string) string {
	prefix := ""
	if strings.IndexAny(name, "0123456789") == 0 {
		prefix = "_"
	}
	return prefix + strings.Replace(strings.Replace(strings.Title(name), "-", "", -1), " ", "", -1)
}
