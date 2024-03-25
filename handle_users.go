package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	util "github.com/pistolpeter/aggreRSS/internal"
	"github.com/pistolpeter/aggreRSS/internal/database"
)

func handleUsersCreate(db *database.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		type parameters struct {
			Name string
		}
		decoder := json.NewDecoder(r.Body)
		params := parameters{}
		err := decoder.Decode(&params)
		if err != nil {
			util.RespondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
			return
		}

		user, err := db.CreateUser(r.Context(), database.CreateUserParams{
			ID:        uuid.New(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			Name:      params.Name,
		})
		if err != nil {
			util.RespondWithError(w, http.StatusInternalServerError, "Couldn't create user")
			return
		}

		util.RespondWithJSON(w, http.StatusOK, databaseUserToUser(user))
	}
}

func handleUsersGet(user database.User) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		util.RespondWithJSON(w, http.StatusOK, databaseUserToUser(user))
	}
}
