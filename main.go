package main

import (
	"flag"
	"log"

	"github.com/jasontconnell/sccreate/conf"
	"github.com/jasontconnell/sccreate/process"
	"github.com/jasontconnell/sitecore/api"
	"github.com/jasontconnell/sitecore/data"
)

func main() {
	cfn := flag.String("c", "config.json", "config filename")
	tmpId := flag.String("t", "", "templateId")
	flag.Parse()

	cfg := conf.LoadConfig(*cfn)

	m, err := process.Load(cfg.ConnectionString, cfg.ProtobufLocation)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(len(m), "items", cfg.CodeStyle)

	tid := api.MustParseUUID(*tmpId)

	node, ok := m[tid]
	if !ok || node.GetTemplateId() != data.TemplateID {
		log.Fatalf("can't find item with id %s or it's not a template", *tmpId)
	}

	err = process.CreateCodeFiles(
		node,
		string(cfg.CodeStyle),
		cfg.MarkupPath,
		cfg.BackendPath,
		cfg.Namespace,
		process.GetTemplatesFromConfig(cfg.Templates),
	)

	if err != nil {
		log.Fatal("can't create code files ", err.Error())
	}
}
