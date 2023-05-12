# sccreate


**Update 5/12/2023**

Creating new renderings and updating the datasource is a pain. Creating the template is fine and is where the architect can make some interesting choices. The "this thing displays this" isn't interesting.

sccreate was created to make this part of it automatic. And in the future it will be expanded to more things to make the entire process of creating a new website quicker.

This only works on ascx projects at the moment because that is what I needed it for, but I will need it for MVC in the near future.

**Features**

sccreate can take a Sitecore template ID and a second template id (for the parent folder template id which will contain the data items for this template), create the rendering in Sitecore, and create the code files in your project.

`sccreate -c config.json -t "[template id]" -f "[folder template id]"`

**Configuration**

sccreate takes a json file (default: `config.json`) with the following configuration settings

`connectionString` - db connection string to master db

`protobufLocation` - for Sitecore 10, the location of the items stored in protobuf format

`templatePath` - this is reserved for future use

`renderingPath` - places to create new renderings

`datasourcePath` - if specified, will just point the rendering datasource location here

`datasourceQuery` - if specifed, will set the rendering datasource location to a query. this can use template syntax

`markupPath` - the place to save new markup code files. this can be cshtml or ascx, and in mvc is usually different from backend path

`backendPath` - the place to save new backend code files

`namespace` - the namespace to use for new code files

`codeStyle` - ascx or mvc

`templates` - a list of files to process for generating code. each template includes the template path (.txt), the output filename (this can use template syntax), and the type it is, either markup or backend

**Example Configuration**

```{ 
    "connectionString": "Master db connection string" ,
    "protobufLocation": "../WebSite/App_Data/items/master/items.master.dat",

    "templatePath": "/sitecore/templates/User Defined/",
    "renderingPath": "/sitecore/layout/Sublayouts/Education",
    "datasourceQuery": "query:/sitecore/content/Content Folder/*[@@templateid='{{.FolderTemplateId}}']",
    "markupPath": "..\\WebSite\\layouts\\MySite\\Components\\",
    "markupReferencePath": "/layouts/Education/Components/",
    "backendPath": "..\\WebSite\\layouts\\MySite\\Components\\",

    "namespace": "MySite.layouts.Components",
    "codeStyle": "ascx",

    "templates": [
        {
            "templateFilename": "tmpl/ascx/ascx.txt",
            "outputFilename": "{{.CleanName}}.ascx",
            "type": "markup"
        },
        {
            "templateFilename": "tmpl/ascx/ascxcs.txt",
            "outputFilename": "{{.CleanName}}.ascx.cs",
            "type": "backend"
        },
        {
            "templateFilename": "tmpl/ascx/designer.txt",
            "outputFilename": "{{.CleanName}}.ascx.designer.cs",
            "type": "backend"
        }
    ]
}```