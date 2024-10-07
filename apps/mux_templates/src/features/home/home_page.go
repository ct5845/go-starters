package home

import (
	"bytes"
	"html/template"
	"mux_templates/src/features/ui"
	"path/filepath"
	"runtime"
)

var homeTemplate *template.Template

func init() {
	// Get the directory where the current file (home.go) is located
	_, b, _, _ := runtime.Caller(0)
	basePath := filepath.Dir(b)

	// Build the path to the template relative to this file
	templatePath := filepath.Join(basePath, "home_page.html")

	homeTemplate = template.Must(
		template.ParseFiles(templatePath),
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
