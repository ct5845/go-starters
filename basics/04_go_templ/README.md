# Go with Templ

Run through the following examples to use a structure similiar to our Web UI projects and using HTMX to serve a page.

You can alternatively look at the `blue-peter` folder for a completed example.

## Steps

1. Download and install templ

   ```bash
   $ go install github.com/a-h/templ/cmd/templ@latest
   $ cd example
   $ go get github.com/a-h/templ
   ```

1. Create our first view a basic hello templ

   ```bash
   $ mkdir libs
   $ cd libs
   $ mddir ui
   $ cd ui
   $ touch hello.templ
   ```

   Then edit [hello.templ](./example/libs/ui/hello.templ)

   ```go
   package main

   templ hello(name string)  {
      <h1>Hello {name}!</h1>
   }
   ```

1. Add template compile to [Makefile](./example/Makefile)

   ```make
   templates:
      templ generate
   ```

   then run

   ```bash
   $ make
   ```

1. Create program entry point program

   ```bash
   $ touch main.go
   ```

   then edit this file to render ouytput from our hello component

   ```go
   package main

   import (
      "04_go_templ/libs/ui"
      "fmt"
      "log"
      "net/http"

      "github.com/a-h/templ"
   )

   func main() {
      component := ui.Hello("World!")

      http.Handle("/", templ.Handler(component))

      fmt.Println("Listening on :8080")

      log.Fatal(http.ListenAndServe(":8080", nil))
   }
   ```

1. Run our server

   ```bash
   $ go run .
   ```