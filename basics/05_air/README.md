# air

Air is a `hot reloader` so you can see your changes immediately. Follow steps below to get it working in the `example` project, or you can alternatively look at the `blue-peter` folder for a completed example.

## Steps

1. Download and install air

   ```bash
   $ go install github.com/air-verse/air@latest
   ```

1. Create air configuration file

   ```bash
   $ cd example
   $ air init
   ```

1. Run the web app by calling `air`

   ```bash
   $ air
   ```

   Visit the server http://localhost:8080

1. Try a hot-reload by editing a few files

   [hello.templ](./example/libs/ui/hello.templ) line 4

   ```go
   <h1>Hello again {name}!</h1>
   ```

   [main.go](./example/main.go) line 12

   ```go
   component := ui.Hello("Underworld!")
   ```

   At this point refreshing you should see the change from `main.go` but not `hello.templ` yet.

1. Update our [Makefile](./example/Makefile) to watch and compile templates & then run air for hot-reloading;

   ```make
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
   ```

   Visit the server http://localhost:8080, and make changes in the [hello.templ](./example/libs/ui/hello.templ) and [main.go](./example/main.go) and they should both be reloaded!
