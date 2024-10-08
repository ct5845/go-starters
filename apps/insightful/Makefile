SHELL := /bin/bash

.PHONY:clean
clean:
	. ./scripts/clean.sh

.PHONY:build-go
# option MODE: build (default) | live
build-go:
	. ./scripts/build_go.sh

.PHONY:build
# option MODE: build (default) | live
# option MINIFY: true (default) | false
build: clean build-go

.PHONY:run
run:
	air

.PHONY:docker-build
docker-build:
	docker build --no-cache -t go_starters/openapi .

.PHONY:docker-run
docker-run:
	make docker-build
	docker run -e PORT=8888 -p 8888:8888 go_starters/openapi

.PHONY:preview
preview: build
	build/main.exe