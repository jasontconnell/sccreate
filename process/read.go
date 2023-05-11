package process

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/jasontconnell/sitecore/api"
	"github.com/jasontconnell/sitecore/data"
)

func Load(connstr, protobufLocation string) (data.ItemMap, error) {
	items, err := api.LoadItems(connstr)
	if err != nil {
		return nil, fmt.Errorf("can't load items from %s. %w", connstr, err)
	}

	var pitems []data.ItemNode
	var perr error
	if protobufLocation != "" {
		wd, _ := os.Getwd()
		path := protobufLocation
		if !filepath.IsAbs(path) {
			path = filepath.Join(wd, path)
		}
		pitems, perr = api.ReadProtobuf(path)
		if perr != nil {
			return nil, fmt.Errorf("can't read items from protobuf %s. %w", protobufLocation, perr)
		}
	}

	_, m := api.LoadItemMap(append(pitems, items...))

	return m, nil
}
