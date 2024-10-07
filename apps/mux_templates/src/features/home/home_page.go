package home

import (
	"bytes"
	_ "embed"
	"html/template"
	"mux_templates/src/features/ui"
)

//go:embed home_page.html
var homePageTemplate string

var homeTemplate *template.Template

func init() {
	homeTemplate = template.Must(
		template.New("home_page").Parse(homePageTemplate),
	)
}

func Page() string {
	header := ui.HeaderMolecule("Home")

	var output bytes.Buffer

	err := homeTemplate.Execute(&output, nil)

	if err != nil {
		panic(err)
	}

	return ui.SiteLayout(header, template.HTML(output.String()))
}
