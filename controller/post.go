package controller

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

func (cntrl Post) Show(w http.ResponseWriter, r *http.Request) {
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

	user := &model.User{
		ID: post.UserID,
	}

	cntrl.Model.Find(user)

	cntrl.View.Render(w, "post.show.html", struct {
		User *model.User
		Post *model.Post
	}{
		User: user,
		Post: post,
	})
}

func (cntrl Post) Delete(w http.ResponseWriter, r *http.Request) {
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

	cntrl.Model.Delete(post)
}
