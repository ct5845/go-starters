# installs our GO based dependencies
.PHONY: install-go
install-go:
	go mod init github.com/ct5845/htmx-starter
	go install github.com/air-verse/air@latest
	go install github.com/a-h/templ/cmd/templ@latest
	go get github.com/a-h/templ
	air init
	./tooling/install/air.toml.sh

# installs our JS/Node based dependencies
.PHONY: install-js
install-js:
	npm i
	npx tailwind init

# install deps and makes common set-up changes
.PHONY: install
install:
	make install-go
	make install-js
	./tooling/install/tailwind.config.sh

# clean compiled code
.PHONY: clean
clean:
	rm -rf temp dist

# remove everything related to our installation step
.PHONY: uninstall
uninstall:
	make clean
	rm -f go.mod go.sum .air.toml package-lock.json tailwind.config.js
	rm -rf node_modules

# build the site for production
.PHONY: build
build:
	templ generate
	go build -o ./dist/main ./src/

# run the program in development mode
.PHONY: dev
dev:
	air