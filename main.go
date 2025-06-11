package main

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/mkamelll/goth-template/components"
)

func main() {
	port := ":3000"
	http.Handle("/", templ.Handler(components.Index("hello, world!")))

	fmt.Println("You're app is running on: http://localhost" + port)
	http.ListenAndServe(port, nil)
}
