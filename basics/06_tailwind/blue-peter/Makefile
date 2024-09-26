.PHONY:build
build:
	make templ-build
	make tailwind-build

.PHONY:tailwind-build
tailwind-build:
	tailwindcss -i ./static/input.css -o ./static/style.min.css --minify
	
.PHONY:tailwind-watch
tailwind-watch:
	powershell -Command "Start-Process -NoNewWindow -FilePath 'tailwindcss' -ArgumentList '-i', './static/input.css', '-o', './static/style.css', '--watch'"

.PHONY:templ-build
templ-build:
	templ generate

.PHONY:templ-watch
templ-watch:
	powershell Start-Process -NoNewWindow -FilePath "templ" -ArgumentList "generate", "--watch"

.PHONY:dev
dev:
	make templ-watch
	make tailwind-watch
	air
