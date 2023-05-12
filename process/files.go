package process

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"text/template"

	"github.com/jasontconnell/sitecore/data"
)

func CreateCodeFiles(node data.ItemNode, style, markupPath, backendPath, namespace string, templates []Template) error {
	err := os.MkdirAll(markupPath, os.ModePerm)
	if err != nil && !os.IsExist(err) {
		return err
	}

	err = os.Mkdir(backendPath, os.ModePerm)
	if err != nil && !os.IsExist(err) {
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

		log.Println("writing", filepath.Join(outpath, filename))
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
