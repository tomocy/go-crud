package middleware

import (
	"fmt"
	"log"
	"net/http"

	"github.com/tomocy/mvc/context"
)

func Authenticate(ctx *context.Context, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sess, err := ctx.SessionStore.Get(r, ctx.Config.Session.Name)
		if err != nil {
			log.Println(err)
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		if _, ok := sess.Values["userID"]; !ok {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		h.ServeHTTP(w, r)
	})
}

func Welcome(ctx *context.Context, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sess, _ := ctx.SessionStore.Get(r, ctx.Config.Session.Name)

		if userID, ok := sess.Values["userID"]; ok {
			http.Redirect(w, r, fmt.Sprintf("/account/%d", userID), http.StatusSeeOther)
			return
		}

		h.ServeHTTP(w, r)
	})
}
