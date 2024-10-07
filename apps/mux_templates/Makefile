SHELL := /bin/bash

.PHONY:install
install:
	go install github.com/air-verse/air@latest
	npm i

.PHONY:clean
clean:
	. ./scripts/clean.sh

.PHONY:build-go
# option MODE: build (default) | live
build-go:
	. ./scripts/build_go.sh

.PHONY:build-css
# option MINIFY: true (default) | false
# option MODE: build (default) | live
# If MODE == live, MINIFY = false
build-css:
	. ./scripts/build_css.sh

.PHONY:build
# option MODE: build (default) | live
# option MINIFY: true (default) | false
build: clean build-css build-go

.PHONY:run
run:
	air

.PHONY:docker-build
docker-build:
	docker build --no-cache -t go_starters/mux_templates .

.PHONY:docker-run
docker-run:
	make docker-build
	docker run -e PORT=4000 -p 4000:4000 go_starters/mux_templates

.PHONY:preview
preview: build
	build/main.exe