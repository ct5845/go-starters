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

.PHONY:run
run:
	air

.PHONY:preview
preview:
	make build
	APP_ENV=production build/app.exe