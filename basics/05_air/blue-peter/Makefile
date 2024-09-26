.PHONY:build
build:
	templ generate

.PHONY:templ-watch
templ-watch:
	powershell Start-Process -NoNewWindow -FilePath "templ" -ArgumentList "generate", "--watch"

.PHONY:dev
dev:
	templ generate
	make templ-watch
	air
