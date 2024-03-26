package main

import (
	"net/http"
	"strconv"

	util "github.com/pistolpeter/aggreRSS/internal"
	"github.com/pistolpeter/aggreRSS/internal/database"
)

func (cfg *apiConfig) handlerPostsGet(user database.User) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		limitStr := r.URL.Query().Get("limit")
		limit := 10
		if specifiedLimit, err := strconv.Atoi(limitStr); err == nil {
			limit = specifiedLimit
		}

		posts, err := cfg.DB.PostsGetByUser(r.Context(), database.PostsGetByUserParams{
			UserID: user.ID,
			Limit:  int32(limit),
		})
		if err != nil {
			util.RespondWithError(w, http.StatusInternalServerError, "Couldn't get posts for user")
			return
		}

		util.RespondWithJSON(w, http.StatusOK, databasePostsToPosts(posts))
	}
}
