package components

import (
	"strings"
)

func stylesHref() string {
	// TODO probably get this with a cache busting hash
	return "static/styles.css"
}

func SiteLayout(header string, content string) string {
	styles := stylesHref()

	layout := `
		<html>
			<head>
				<title>GO Server</title>
				<link rel="stylesheet" href="@stylecss"/>
			</head>
			<body>
				<div class="h-screen w-screen flex flex-col gap-4 items-center justify-center">
					@header
					@content
				</div>
			</body>
		</html>
	`

	layout = strings.Replace(layout, "@stylecss", styles, 1)
	layout = strings.Replace(layout, "@header", header, 1)
	layout = strings.Replace(layout, "@content", content, 1)

	return layout
}
