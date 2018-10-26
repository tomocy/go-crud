package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/tomocy/crud/model"
	"github.com/tomocy/mvc/controller"
)

type Comment struct {
	*controller.Base
}

func NewComment() controller.Controller {
	return &Comment{
		Base: controller.NewBase(),
	}
}

func (cntrl Comment) New(w http.ResponseWriter, r *http.Request) {
	s := mux.Vars(r)["id"]
	id, err := strconv.Atoi(s)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	post := &model.Post{
		ID: id,
	}

	cntrl.Model.Find(post)

	if err := cntrl.View.Render(w, "comment.new.html", struct {
		Post *model.Post
	}{
		Post: post,
	}); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
