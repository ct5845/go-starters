package ui

import (
	"bytes"
	"html/template"
)

var headerTemplate *template.Template

type LinkLabel struct {
	Label  string
	Active bool
}

type Link struct {
	Href      string
	LinkLabel LinkLabel
}

type Header struct {
	Links []Link
}

func init() {
	// Create the base template and parse both "linkLabel" and "link" into the same template object
	// First, create a new template and define the "link" template
	headerTemplate = template.Must(template.New("header").Parse(`
<h1 class="text-4xl mb-4">GO Website Start</h1>
<nav class="border border-2 p-4 rounded-2xl">
	<ul class="flex gap-8">
		{{ range .Links }}
			{{ template "link" . }}
		{{end}}
	</ul>
</nav>`))

	template.Must(headerTemplate.New("link").Parse(`
<li>
	<a class="block border-4 border-double rounded-md p-4" href="{{ .Href }}">
  	{{ template "linkLabel" .LinkLabel }}
	</a>
</li>`))

	// Now define the "linkLabel" template and associate it with the same template object
	template.Must(headerTemplate.New("linkLabel").Parse(`
{{if .Active}}
	<strong class="underline">{{ .Label }}</strong>
{{else}}
	{{ .Label }}
{{end}}`))
}

func HeaderTemplate() *template.Template {
	return headerTemplate
}

func HeaderMolecule(active string) template.HTML {
	header := Header{
		Links: []Link{
			{
				Href: "/",
				LinkLabel: LinkLabel{
					Label:  "Home",
					Active: active == "Home",
				},
			},
			{
				Href: "/about",
				LinkLabel: LinkLabel{
					Label:  "About",
					Active: active == "About",
				},
			},
		},
	}

	var output bytes.Buffer

	err := headerTemplate.Execute(&output, header)

	if err != nil {
		panic(err)
	}

	return template.HTML(output.String())
}
