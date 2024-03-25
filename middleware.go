package main

import (
	"net/http"

	util "github.com/pistolpeter/aggreRSS/internal"
	"github.com/pistolpeter/aggreRSS/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, *database.Queries)

type authedUser func(database.User) http.HandlerFunc

func middlewareAuth(handler authedUser, db *database.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("Authorization")

		user, err := db.GetUser(r.Context(), apiKey)
		if err != nil {
			util.RespondWithError(w, http.StatusUnauthorized, "Couldn't get user")
			return
		}
		handler(user)(w, r)
	}
}
