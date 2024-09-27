# Go Server

Run through the following examples to create a folder `/example` for a simple "Hello World" Webserver in Go.

You can alternatively look at the `blue-peter` folder for a completed example.

## Steps

1. Make our directory `$/basics/03_go_server/example`

   ```bash
   $ cd basics/03-go
   $ mkdir example
   $ cd example
   ```

1. Initialise our module (like npm init...) with [`go mod init`](https://go.dev/ref/mod#go-mod-init)

   ```bash
   go mod init example/03_go_server
   ```

1. Create our main/entry point file `hello.go`

   ```bash
    touch hello.go
   ```

1. Start with our basic Hello World

   ```go
   package main

   import (
      "fmt"
      "log"
      "net/http"
   )

   func main() {
      http.HandleFunc("/", handler)
      log.Fatal(http.ListenAndServe(":8080", nil))
   }

   func handler(w http.ResponseWriter, r *http.Request) {
      fmt.Fprintf(w, "Hello World")
   }
   ```

1. Run our Hello World

   ```bash
   go run .
   ```

1. Visit the the [localhost](http://localhost:8080)
