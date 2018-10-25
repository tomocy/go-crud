package controller

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/tomocy/crud/model"
	"github.com/tomocy/mvc/controller"
)

var (
	ErrContentEmpty = errors.New("mvc: content cannot be empty")
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

func (cntrl Post) Create(w http.ResponseWriter, r *http.Request) {
	sess, _ := cntrl.SessionStore.Get(r, cntrl.Config.Session.Name)
	userID := sess.Values["userID"].(int)

	if err := cntrl.validateToCreate(r); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	post := &model.Post{
		UserID:  userID,
		Content: r.PostFormValue("content"),
	}

	cntrl.Model.Create(post)

	http.Redirect(w, r, fmt.Sprintf("/post/%d", post.ID), http.StatusFound)
}

func (cntrl Post) validateToCreate(r *http.Request) error {
	if r.PostFormValue("content") == "" {
		return ErrContentEmpty
	}

	return nil
}
