package main

import (
	"flag"
	"log"

	"github.com/google/uuid"
	"github.com/jasontconnell/sccreate/conf"
	"github.com/jasontconnell/sccreate/process"
	"github.com/jasontconnell/sitecore/api"
	"github.com/jasontconnell/sitecore/data"
)

func main() {
	cfn := flag.String("c", "config.json", "config filename")
	tmpId := flag.String("t", "", "template id")
	folderTmpId := flag.String("f", "", "folder template id")
	flag.Parse()

	cfg := conf.LoadConfig(*cfn)

	m, err := process.Load(cfg.ConnectionString, cfg.ProtobufLocation)
	if err != nil {
		log.Fatal(err)
	}

	tid := api.MustParseUUID(*tmpId)
	node, ok := m[tid]
	if !ok || node.GetTemplateId() != data.TemplateID {
		log.Fatalf("can't find item with id %s or it's not a template", *tmpId)
	}

	fid, _ := uuid.Parse(*folderTmpId)
	dsfolder, ok := m[fid]

	log.Println("creating rendering for", node.GetName())
	err = process.CreateRendering(
		cfg.ConnectionString,
		m,
		node,
		dsfolder,
		cfg.RenderingPath,
		cfg.DatasourcePath,
		cfg.DatasourceQuery,
		cfg.MarkupReferencePath,
		cfg.CodeStyle,
	)

	if err != nil {
		log.Println("can't create rendering ", err.Error())
	}

	log.Println("creating code files (", cfg.CodeStyle, ") for", node.GetName())
	err = process.CreateCodeFiles(
		node,
		string(cfg.CodeStyle),
		cfg.MarkupPath,
		cfg.BackendPath,
		cfg.Namespace,
		process.GetTemplatesFromConfig(cfg.Templates),
	)

	if err != nil {
		log.Println("can't create code files ", err.Error())
	}
}
