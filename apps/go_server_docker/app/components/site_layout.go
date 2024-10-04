package components

import "strings"

func SiteLayout(header string, content string) string {
	layout := `
		<html>
			<head>
				<title>GO Server</title>
			</head>
			<body>
				<div>
					@header
					@content
				</div>
			</body>
		</html
	`

	layout = strings.Replace(layout, "@header", header, 1)
	layout = strings.Replace(layout, "@content", content, 1)

	return layout
}
