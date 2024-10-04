package components

import (
	"fmt"
	"strings"
)

func linkLabel(label string, active string) string {
	if label == active {
		return strings.Replace(`<strong>@label</strong>`, "@label", label, 1)
	} else {
		return label
	}
}

func link(href string, label string, active string) string {
	label = linkLabel(label, active)
	var link = `<li><a href="@href">@label</a></li>`

	link = strings.Replace(link, "@href", href, 1)
	link = strings.Replace(link, "@label", label, 1)

	return link
}

func HeaderMolecule(env string, active string) string {
	h1 := fmt.Sprintf(`<h1>GO Server Docker - %s</h1>`, env)
	homeLink := link("/", "Home", active)
	aboutLink := link("/about", "About", active)
	links := homeLink + aboutLink

	return fmt.Sprintf(`
		%s
		<nav>
			<ul>
				%s
			</ul>
		</nav>`, h1, links)
}
