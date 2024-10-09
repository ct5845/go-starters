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
	docker build -t go_starters/insightful .

.PHONY:docker-run
docker-run:
	make docker-build
	cat ./config/.env ./config/.env.local > ./config/.env.docker.local
	docker run --env-file ./config/.env.local \
		-e PORT=8888 \
		-p 8888:8888 \
		-v ${HOME}/.aws:/root/.aws:ro \
		go_starters/insightful
	rm ./config/.env.docker.local

.PHONY:preview
preview: build
	build/main.exe