package about

import (
	"html/template"
	"mux_templates/src/features/ui"
)

var tpl = `<h2 class="text-2xl">About</h2>`

func Page() string {
	header := ui.HeaderMolecule("About")
	content := template.HTML(tpl)
	return ui.SiteLayout(header, content)
}
