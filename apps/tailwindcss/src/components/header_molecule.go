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
	var link = `<li><a class="block border-4 border-double rounded-md p-4" href="@href">@label</a></li>`

	link = strings.Replace(link, "@href", href, 1)
	link = strings.Replace(link, "@label", label, 1)

	return link
}

func HeaderMolecule(active string) string {
	h1 := `<h1 class="text-4xl mb-4">GO Website Start</h1>`

	homeLink := link("/", "Home", active)
	aboutLink := link("/about", "About", active)
	links := homeLink + aboutLink

	return fmt.Sprintf(`
		%s
		<nav class="border border-2 p-4 rounded-2xl">
			<ul class="flex gap-8">
				%s
			</ul>
		</nav>`, h1, links)
}
