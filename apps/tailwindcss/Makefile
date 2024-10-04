.PHONY:clean
clean:
	rm -rf .build build

.PHONY:install
install:
	go install github.com/air-verse/air@latest

.PHONY:build
build:
	go build -o ./build/ ./app/
	cp .env* ./build/

.PHONY:docker-build
docker-build:
	docker build $(FLAGS) -t go_starters/server_docker .

.PHONY:docker-run
docker-run:
	make docker-build
	docker run -e PORT=8080 -p 8080:8080 go_starters/server_docker

.PHONY:run
run:
	air

.PHONY:preview
preview:
	make build
	APP_ENV=production build/app.exe