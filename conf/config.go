package conf

import "github.com/jasontconnell/conf"

type Style string

const (
	WebForms Style = "ascx"
	MVC            = "mvc"
)

type Config struct {
	ConnectionString string `json:"connectionString"`
	ProtobufLocation string `json:"protobufLocation"`

	TemplatePath   string `json:"templatePath"`
	RenderingPath  string `json:"renderingPath"`
	DatasourcePath string `json:"datasourcePath"`
	MarkupPath     string `json:"markupPath"`
	BackendPath    string `json:"backendPath"`
	Namespace      string `json:"namespace"`
	CodeStyle      Style  `json:"codeStyle"`

	Templates []TemplateConfig `json:"templates"`
}

type TemplateConfig struct {
	TemplateFilename string `json:"templateFilename"`
	OutputFilename   string `json:"outputFilename"`
	Type             string `json:"type"`
}

func LoadConfig(file string) Config {
	config := Config{}
	conf.LoadConfig(file, &config)

	return config
}
