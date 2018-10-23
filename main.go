package main

import (
	"log"

	"github.com/tomocy/crud/controller"
	"github.com/tomocy/crud/model"
	"github.com/tomocy/mvc"
	"github.com/tomocy/mvc/path"
)

func main() {
	app, err := mvc.New("config/app.json")
	if err != nil {
		log.Fatalln(err)
	}

	app.Model.AutoMigrate(model.User{})

	route(app)

	log.Fatalln(app.ListenAndServe())
}

func route(app *mvc.MVC) {
	account := controller.NewAccount()
	app.Router.Register(account, []path.Path{
		{
			Methods: []string{"GET"},
			Path:    "/account/create",
			Controller: path.Controller{
				Method: "New",
			},
		},
		{
			Methods: []string{"POST"},
			Path:    "/account/create",
			Controller: path.Controller{
				Method: "Create",
			},
		},
		{
			Methods: []string{"GET"},
			Path:    "/account/{id}",
			Controller: path.Controller{
				Method: "Show",
			},
		},
		{
			Methods: []string{"OPTIONS"},
			Path:    "/account/{id}",
			Controller: path.Controller{
				Method: "CORSOptions",
			},
		},
		{
			Methods: []string{"DELETE"},
			Path:    "/account/{id}",
			Controller: path.Controller{
				Method: "Delete",
			},
		},
	})
}
