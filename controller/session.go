package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/sessions"
	"github.com/tomocy/crud/model"
	"github.com/tomocy/mvc/controller"
	"golang.org/x/crypto/bcrypt"
)

type Session struct {
	*controller.Base
}

func NewSession() controller.Controller {
	return &Session{
		Base: controller.NewBase(),
	}
}

func (cntrl Session) New(w http.ResponseWriter, r *http.Request) {
	sess, err := cntrl.SessionStore.Get(r, cntrl.Config.Session.Name)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if _, ok := sess.Values["SESSID"]; ok {
		http.Redirect(w, r, fmt.Sprintf("/account/%d", sess.Values["userID"]), http.StatusSeeOther)
		return
	}

	if err := cntrl.View.Render(w, "session.new.html", nil); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (cntrl Session) Create(w http.ResponseWriter, r *http.Request) {
	user := &model.User{
		Email: r.PostFormValue("email"),
	}
	cntrl.Model.Find(user)

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(r.PostFormValue("password"))); err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	sess, err := cntrl.startSession(w, r)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	sess.Values["userID"] = user.ID

	if err := cntrl.SessionStore.Save(r, w, sess); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/account/%d", user.ID), http.StatusSeeOther)
}

func (cntrl Session) startSession(w http.ResponseWriter, r *http.Request) (*sessions.Session, error) {
	sess, err := cntrl.SessionStore.New(r, cntrl.Config.Session.Name)
	if err != nil {
		return nil, err
	}

	sess.Values["SESSID"] = uuid.New().String()

	if err := cntrl.SessionStore.Save(r, w, sess); err != nil {
		return nil, err
	}

	return sess, nil
}

func (cntrl Session) Delete(w http.ResponseWriter, r *http.Request) {
	sess, err := cntrl.SessionStore.Get(r, cntrl.Config.Session.Name)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	sess.Values = make(map[interface{}]interface{})

	if err := cntrl.SessionStore.Save(r, w, sess); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
