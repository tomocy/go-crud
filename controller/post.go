package controller

import (
	"log"
	"net/http"

	"github.com/tomocy/mvc/controller"
)

type Post struct {
	*controller.Base
}

func NewPost() controller.Controller {
	return &Post{
		Base: controller.NewBase(),
	}
}

func (cntrl Post) New(w http.ResponseWriter, r *http.Request) {
	if err := cntrl.View.Render(w, "post.new.html", nil); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
