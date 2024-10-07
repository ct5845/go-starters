package main

import (
	"context"
	"fmt"
	"net/http"
	"open_api/src/utils"
	"os"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humago"
)

type FooInput struct {
	Body struct {
		Foo string `json:"foo" maxLength:"3" doc:"Foo?"`
	}
}

type FooBody struct {
	Foo string `json:"foo" doc:"Foo?"`
}

type FooOutput struct {
	Body FooBody
}

var port string

func init() {
	utils.LoadEnv()

	port = os.Getenv("PORT")
	if port == "" {
		port = "8888"
	}
}

func main() {
	utils.LoadEnv()

	router := http.NewServeMux()
	api := humago.New(router, huma.DefaultConfig("My API", "1.0.0"))

	huma.Register(api, huma.Operation{
		OperationID:   "foo",
		Method:        http.MethodPost,
		Path:          "/foo",
		Summary:       "Foo?",
		DefaultStatus: http.StatusOK,
		Errors:        []int{http.StatusUnprocessableEntity},
		Tags:          []string{"Test"},
	},
		func(ctx context.Context, input *FooInput) (*FooOutput, error) {
			if input.Body.Foo == "bar" {
				return &FooOutput{Body: FooBody{Foo: "!!!! FooBar !!!!"}}, nil
			} else {
				return &FooOutput{Body: FooBody{Foo: input.Body.Foo + " ???!!!"}}, nil
			}
		})

	fmt.Printf("Server running; \033[34;1;4mhttp://localhost:%s/docs\033[0m\n", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), router)
}
