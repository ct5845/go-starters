package ui

import (
	"html/template"
	"strings"
)

var siteLayoutTmpl *template.Template

func init() {
	layout := `
		<html>
			<head>
				<title>GO Server</title>
				<link rel="stylesheet" href="{{ .Styles }}"/>
			</head>
			<body>
				<div class="h-screen w-screen flex flex-col gap-4 items-center justify-center">
					{{ .Header }}
					{{ .Content }}
				</div>
			</body>
		</html>
	`

	siteLayoutTmpl = template.Must(template.New("siteLayout").Parse(layout))
}

func stylesHref() string {
	// TODO probably get this with a cache busting hash
	return "static/styles.css"
}

func SiteLayout(header template.HTML, content template.HTML) string {
	model := struct {
		Styles  string
		Header  template.HTML
		Content template.HTML
	}{
		Styles:  stylesHref(),
		Header:  header,
		Content: content,
	}

	var output strings.Builder

	err := siteLayoutTmpl.Execute(&output, model)
	if err != nil {
		panic(err)
	}

	return output.String()
}
