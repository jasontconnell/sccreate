package process

import "github.com/jasontconnell/sccreate/conf"

type Template struct {
	TemplateFilename string
	OutputFilename   string
	Type             string
}

func GetTemplatesFromConfig(tmpls []conf.TemplateConfig) []Template {
	list := []Template{}
	for _, tmpl := range tmpls {
		list = append(list, Template{
			TemplateFilename: tmpl.TemplateFilename,
			OutputFilename:   tmpl.OutputFilename,
			Type:             tmpl.Type,
		})
	}
	return list
}

type TemplateModel struct {
	Name      string
	CleanName string
	Namespace string
}

type QueryModel struct {
	FolderTemplateId string
	TemplateId       string
}
