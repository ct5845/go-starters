# Tailwind

In this slightly more invovled example, we're going to bring in Tailwind, which includes some update to the Routing too.

## Steps

1. Get Tailwind CLI

   `windows x64`
   ```bash
   $ cd example
   $ curl -L --ssl-no-revoke https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-windows-x64.exe --output ./tailwindcss.exe
   ```

   `linux x64`
   ```bash
   $ cd example
   $ curl -L --ssl-no-revoke https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-linux-x64 --output ./tailwindcss
   $ chmod +x ./tailwindcss
   ```

1. Create our tailwind config file

   `windows`
   ```bash
   $ ./tailwindcss.exe init
   ```

   `linux`
   ```bash
   $ ./tailwindcss init
   ```

1. Also init air for your local environment

   ```bash
   $ air init
   ```

1. Alter make file to create tailwind build steps

   `windows`
   ```make
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
   ```

   `linux`
   ```make
   .PHONY:build
   build:
      make templ-build
      make tailwind-build

   .PHONY:tailwind-build
   tailwind-build:
      ./tailwindcss -i ./static/input.css -o ./static/style.min.css --minify

   .PHONY:tailwind-watch
   tailwind-watch:
      nohup   ./tailwindcss -i ./static/input.css -o ./static/style.css --watch &

   .PHONY:templ-build
   templ-build:
      templ generate

   .PHONY:templ-watch
   templ-watch:
      templ generate --watch &

   .PHONY:dev
   dev:
      make templ-watch
      make tailwind-watch
      air
   ```

1. Give the build a go, run

   ```bash
   $ make tailwind-build
   ```

   This should create [style.min.css](./example/static/style.min.css) which compiles the base tailwind styles.

   You can play around with adding more/different styles to [input.css](./example/static/input.css)

1. Get Tailwind working for this site

   1. Run the app in dev (watch) mode

      ```bash
      $ make dev
      ```

   1. Change our [hello.templ](./example/libs/ui/hello.templ)

      ```go
      package ui

      templ Hello(name string)  {
         <h1 class="text-2xl underline">Hello {name}!</h1>
      }
      ```

      At this stage you won't see the changes nor will Tailwind be picking up the files.

   1. Update our [tailwind config](./example/tailwind.config.js) to read from .templ files

      ```javascript
      /** @type {import('tailwindcss').Config} */
      module.exports = {
        content: ["**/*.templ"],
        theme: {
          extend: {},
        },
        plugins: [],
      };
      ```

      This will add the two classes into the generated [styles.css](./exmaple/static/styles.css) but we still need to include these in our html response

   1. Add a Page Layout UI component, then use this component to wrap our hello.templ

      ```bash
      $ cd libs/ui
      $ touch page.layout.templ
      ```

      [page.layout](./example/libs/ui/page.layout.templ)

      ```go
      package ui

      templ PageLayout() {
         <html>
            <head>
                  <title>06_Tailwind</title>
                  <meta charset="UTF-8"/>
               <meta name="viewport" content="width=device-width, initial-scale=1.0"/>
                  <link rel="stylesheet" href="/static/style.css"/>
            </head>
            <body>
                  { children... }
            </body>
         </html>
      }
      ```

   1. Then update [hello.templ](./example/libs/ui/hello.templ) to use the layout

      ```go
      package ui

      templ Hello(name string)  {
         @PageLayout() {
            <h1 class="text-2xl underline">Hello again {name}!</h1>
         }
      }
      ```

   1. Update [main.go](./example/main.go) to use the standard router, and handle request to the static files (our css in this case)

      ```go
      package main

      import (
         "06_tailwind/libs/ui"
         "fmt"
         "log"
         "net/http"
         "path/filepath"

         "github.com/a-h/templ"
      )

      func serveStatic(w http.ResponseWriter, r *http.Request) {
         filePath := r.URL.Path[len("/static/"):] // like doing r.URL.path.substring(9);
         fullPath := filepath.Join(".", "static", filePath)
         http.ServeFile(w, r, fullPath)
      }

      func main() {
         router := http.NewServeMux()

         router.HandleFunc("/static/", serveStatic);
         router.Handle("/", templ.Handler(ui.Hello("templ")));

         fmt.Println("Listening on :8080")

         log.Fatal(http.ListenAndServe(":8080", router))
      }
      ```

Your home path should now render your hello/templ route, along with tailwind!
