package process

import (
	"fmt"
	"path"
	"time"

	"github.com/google/uuid"
	"github.com/jasontconnell/sccreate/conf"
	"github.com/jasontconnell/sitecore/api"
	"github.com/jasontconnell/sitecore/data"
)

func CreateRendering(connstr string, m data.ItemMap, node data.ItemNode, folderTemplate data.ItemNode, renderingPath, datasourcePath, datasourceQuery, markupReferencePath string, style conf.Style) error {
	id := uuid.New()
	cname := getCleanName(node.GetName()) + "Component"

	parent := m.FindItemByPath(renderingPath)
	if parent == nil {
		return fmt.Errorf("couldn't find rendering path %s", renderingPath)
	}

	for _, c := range parent.GetChildren() {
		if c.GetName() == cname {
			return fmt.Errorf("component already exists with name %s", cname)
		}
	}

	templateId := data.ControllerRenderingId
	if style == conf.WebForms {
		templateId = data.SublayoutRenderingId
	}

	rendering := data.NewItemNode(id, cname, templateId, parent.GetId(), data.EmptyID)

	if style == conf.WebForms {
		rpath := path.Join(markupReferencePath, getCleanName(node.GetName())+".ascx")
		pathfv := data.NewFieldValue(data.SublayoutRenderingPathFieldId, rendering.GetId(), "", rpath, data.English, 1, data.SharedFields)
		rendering.AddFieldValue(pathfv)
	} else if style == conf.MVC {
		return fmt.Errorf("not yet supported")
	}

	dstmpfv := data.NewFieldValue(data.RenderingDatasourceTemplateFieldId, rendering.GetId(), "", node.GetPath(), data.English, 1, data.SharedFields)
	rendering.AddFieldValue(dstmpfv)

	dslocation := datasourcePath
	if datasourceQuery != "" {
		qmodel := QueryModel{FolderTemplateId: sitecoreStyleGuid(folderTemplate.GetId()), TemplateId: sitecoreStyleGuid(node.GetId())}
		tmp, dslerr := processInlineTemplate(datasourceQuery, qmodel)
		if dslerr != nil {
			return fmt.Errorf("error processing datasource query template %s. %w", datasourceQuery, dslerr)
		}
		dslocation = tmp
	}

	dslocfv := data.NewFieldValue(data.RenderingDatasourceLocationFieldId, rendering.GetId(), "", dslocation, data.English, 1, data.SharedFields)
	rendering.AddFieldValue(dslocfv)

	dispv := data.NewFieldValue(data.DisplayNameFieldId, rendering.GetId(), "", "", data.English, 1, data.UnversionedFields)
	rendering.AddFieldValue(dispv)

	cuser := data.NewFieldValue(data.CreatedByFieldId, rendering.GetId(), "", "sitecore\\admin", data.English, 1, data.VersionedFields)
	rendering.AddFieldValue(cuser)

	uuser := data.NewFieldValue(data.UpdatedByFieldId, rendering.GetId(), "", "sitecore\\admin", data.English, 1, data.VersionedFields)
	rendering.AddFieldValue(uuser)

	cdate := data.NewFieldValue(data.CreateDateFieldId, rendering.GetId(), "", time.Now().Format("20060201T150405"), data.English, 1, data.VersionedFields)
	rendering.AddFieldValue(cdate)

	upditems, updflds := api.BuildUpdateItems(m, nil, []data.ItemNode{rendering})

	_, err := api.Update(connstr, upditems, updflds)

	for _, i := range upditems {
		fmt.Println(i)
	}

	for _, f := range updflds {
		fmt.Println(f)
	}

	return err
}
