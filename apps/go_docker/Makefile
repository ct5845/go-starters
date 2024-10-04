.PHONY:clean
clean:
	rm -rf .build build

.PHONY:build
build:
	go build -o ./build/ ./app/

.PHONY:docker-build
docker-build:
	docker build $(FLAGS) -t go_starters/docker .

.PHONY:run
run:
	air

.PHONY:preview
preview:
	make build
	build/app.exe