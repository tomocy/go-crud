package main

import (
	"log"

	"github.com/tomocy/crud/controller"
	"github.com/tomocy/crud/middleware"
	"github.com/tomocy/crud/model"
	"github.com/tomocy/mvc"
	"github.com/tomocy/mvc/path"
)

func main() {
	app, err := mvc.New("config/app.json")
	if err != nil {
		log.Fatalln(err)
	}

	app.Model.AutoMigrate(
		model.User{}, model.Post{}, model.Comment{},
	)

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
	}, middleware.Welcome)
	app.Router.Register(account, []path.Path{
		{
			Methods: []string{"GET"},
			Path:    "/account/{id}",
			Controller: path.Controller{
				Method: "Show",
			},
		},
		{
			Methods: []string{"DELETE"},
			Path:    "/account/{id}",
			Controller: path.Controller{
				Method: "Delete",
			},
		},
	}, middleware.Authenticate)

	sess := controller.NewSession()
	app.Router.Register(sess, []path.Path{
		{
			Methods: []string{"GET"},
			Path:    "/login",
			Controller: path.Controller{
				Method: "New",
			},
		},
		{
			Methods: []string{"POST"},
			Path:    "/login",
			Controller: path.Controller{
				Method: "Create",
			},
		},
	}, middleware.Welcome)
	app.Router.Register(sess, []path.Path{
		{
			Methods: []string{"DELETE"},
			Path:    "/logout",
			Controller: path.Controller{
				Method: "Delete",
			},
		},
	}, middleware.Authenticate)

	post := controller.NewPost()
	app.Router.Register(post, []path.Path{
		{
			Methods: []string{"GET"},
			Path:    "/post/create",
			Controller: path.Controller{
				Method: "New",
			},
		},
		{
			Methods: []string{"POST"},
			Path:    "/post/create",
			Controller: path.Controller{
				Method: "Create",
			},
		},
		{
			Methods: []string{"DELETE"},
			Path:    "/post/{id}",
			Controller: path.Controller{
				Method: "Delete",
			},
		},
	}, middleware.Authenticate)
	app.Router.Register(post, []path.Path{
		{
			Methods: []string{"GET"},
			Path:    "/post/{id}",
			Controller: path.Controller{
				Method: "Show",
			},
		},
	})

	comment := controller.NewComment()
	app.Router.Register(comment, []path.Path{
		{
			Methods: []string{"GET"},
			Path:    "/post/{id}/comment/create",
			Controller: path.Controller{
				Method: "New",
			},
		},
		{
			Methods: []string{"POST"},
			Path:    "/post/{id}/comment/create",
			Controller: path.Controller{
				Method: "Create",
			},
		},
	})
}
