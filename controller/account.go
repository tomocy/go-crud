package controller

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	m "github.com/tomocy/crud/model"
	"github.com/tomocy/mvc/controller"
	"github.com/tomocy/mvc/model"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrEmailEmpty    = errors.New("mvc: email cannot be empty")
	ErrPasswordEmpty = errors.New("mvc: password cannot be empty")
)

type Account struct {
	*controller.Base
}

func NewAccount(model *model.Model) controller.Controller {
	return &Account{
		Base: controller.NewBase(model),
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

	http.Redirect(w, r, "/", http.StatusFound)
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
