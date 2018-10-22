package main

import (
	"log"
	"net/http"

	"github.com/tomocy/crud/controller"
	"github.com/tomocy/crud/model"
	"github.com/tomocy/mvc"
	"github.com/tomocy/mvc/path"
)

func main() {
	app, err := mvc.New()
	if err != nil {
		log.Fatalln(err)
	}

	app.Model.AutoMigrate(model.User{})

	route(app)

	addr := ":5050"
	log.Println("listning on " + addr + "...")
	log.Fatalln(http.ListenAndServe(addr, app))
}

func route(app *mvc.MVC) {
	account := controller.NewAccount(app.Model)
	app.Router.Register(path.Path{
		Methods: []string{"GET"},
		Path:    "/account/create",
		Controller: path.Controller{
			Method: "New",
		},
	}, account)
	app.Router.Register(path.Path{
		Methods: []string{"POST"},
		Path:    "/account/create",
		Controller: path.Controller{
			Method: "Create",
		},
	}, account)
	app.Router.Register(path.Path{
		Methods: []string{"GET"},
		Path:    "/account/{id}",
		Controller: path.Controller{
			Method: "Show",
		},
	}, account)
}
