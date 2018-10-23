package controller

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	m "github.com/tomocy/crud/model"
	"github.com/tomocy/mvc/controller"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrEmailEmpty    = errors.New("mvc: email cannot be empty")
	ErrPasswordEmpty = errors.New("mvc: password cannot be empty")
)

type Account struct {
	*controller.Base
}

func NewAccount() controller.Controller {
	return &Account{
		Base: controller.NewBase(),
	}
}

func (cntrl Account) New(w http.ResponseWriter, r *http.Request) {
	if err := cntrl.View.Render(w, "account.new.html", nil); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (cntrl Account) Show(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	s := vars["id"]
	id, err := strconv.Atoi(s)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user := &m.User{
		ID: id,
	}

	cntrl.Model.Find(user)

	cntrl.View.Render(w, "account.show.html", user)
}

func (cntrl Account) Create(w http.ResponseWriter, r *http.Request) {
	if err := validateToCreate(r); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(r.PostFormValue("password")), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user := &m.User{
		Email:    r.PostForm.Get("email"),
		Password: string(hash),
	}

	cntrl.Model.Create(user)

	cntrl.startSession(w, r)

	cntrl.View.Render(w, "account.show.html", user)
}

func validateToCreate(r *http.Request) error {
	if r.PostFormValue("email") == "" {
		return ErrEmailEmpty
	}

	if r.PostFormValue("password") == "" {
		return ErrPasswordEmpty
	}

	return nil
}

func (cntrl Account) startSession(w http.ResponseWriter, r *http.Request) error {
	sess, err := cntrl.SessionStore.New(r, cntrl.Config.Session.Name)
	if err != nil {
		return err
	}

	sess.Values["SESSID"] = uuid.New().String()

	if err := cntrl.SessionStore.Save(r, w, sess); err != nil {
		return err
	}

	return nil
}

func (cntrl Account) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	s := vars["id"]
	id, err := strconv.Atoi(s)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user := &m.User{
		ID: id,
	}

	cntrl.Model.Delete(user)

	cntrl.View.Render(w, "account.new.html", nil)
}
