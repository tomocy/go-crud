package controller

import (
	"log"
	"net/http"

	"github.com/tomocy/mvc/controller"
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

	if _, ok := sess.Values["id"]; ok {
		http.Redirect(w, r, "/account/"+sess.Values["userID"].(string), http.StatusSeeOther)
		return
	}

	cntrl.View.Render(w, "session.new.html", nil)
}
