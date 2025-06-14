package main

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	views "github.com/mkamelll/goth-template/resources/views"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	port := ":3000"
	http.Handle("/", templ.Handler(views.Index("hello, world!")))

	fmt.Println("You're app is running on: http://localhost" + port)
	http.ListenAndServe(port, nil)
}
