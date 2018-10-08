package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/semcelik/go-mongo-example/app"
)

func main() {
	fmt.Println("App starting...")
	app := &app.App{}
	app.Initialize()
	fmt.Println("App started at 8000")
	log.Fatal(http.ListenAndServe(":8000", app.Router))
}
