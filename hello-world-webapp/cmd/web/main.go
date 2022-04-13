package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bhavanichandra/golang-learning/hello-world-webapp/pkg/config"
	"github.com/bhavanichandra/golang-learning/hello-world-webapp/pkg/handlers"
	"github.com/bhavanichandra/golang-learning/hello-world-webapp/pkg/render"
)

const PORT = ":8080"

//main is the main application function
func main() {

	var app config.AppConfig
	tc, err := render.CreateTemplateCache()

	render.NewTemplates(&app)

	if err != nil {
		log.Fatal("Cannot create template cache")
	}
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)

	handlers.NewHandlers(repo)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprintf("Listening to port %s", PORT))

	_ = http.ListenAndServe(PORT, nil)
}
