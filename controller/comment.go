package controller

import (
	"fmt"
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

func (cntrl Comment) Create(w http.ResponseWriter, r *http.Request) {
	s := mux.Vars(r)["id"]
	postID, err := strconv.Atoi(s)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	sess, _ := cntrl.SessionStore.Get(r, cntrl.Config.Session.Name)
	userID := sess.Values["userID"].(int)

	comment := &model.Comment{
		UserID:  userID,
		PostID:  postID,
		Content: r.PostFormValue("content"),
	}

	cntrl.Model.Create(comment)

	http.Redirect(w, r, fmt.Sprintf("/post/%d/comment", postID), http.StatusFound)
}

func (cntrl Comment) Show(w http.ResponseWriter, r *http.Request) {
	s := mux.Vars(r)["id"]
	postID, err := strconv.Atoi(s)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	post := &model.Post{
		ID: postID,
	}

	cntrl.Model.Find(post)

	sess, _ := cntrl.SessionStore.Get(r, cntrl.Config.Session.Name)
	userID := sess.Values["userID"].(int)

	user := &model.User{
		ID: userID,
	}

	comments := []model.Comment{}

	cntrl.Model.Find(&comments, "post_id=?", postID)

	cntrl.View.Render(w, "post.show.html", struct {
		User     *model.User
		Post     *model.Post
		Comments []model.Comment
	}{
		User:     user,
		Post:     post,
		Comments: comments,
	})
}
