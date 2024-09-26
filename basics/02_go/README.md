# Go

Run through the following examples to create a folder `/example` for a simple "Hello World" in Go.

You can alternatively look at the `blue-peter` folder for a completed example.

## Steps

1. Make our directory `$/basics/02_go/example`

   ```bash
   cd basics/02_go
   mkdir example
   cd example
   ```

1. Initialise our module (like npm init...) with [`go mod init`](https://go.dev/ref/mod#go-mod-init)

   ```bash
   go mod init example/02_go
   ```

1. Create our main/entry point file `hello.go`

   ```bash
    touch hello.go
   ```

1. Create our Hello World

   ```go
   package main

   import "fmt"

   func main() {
       fmt.Println("Hello, World!")
   }
   ```

1. Run our Hello World

   ```bash
   go run .
   ```
