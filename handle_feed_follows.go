package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	util "github.com/pistolpeter/aggreRSS/internal"
	"github.com/pistolpeter/aggreRSS/internal/database"
)

func (cfg *apiConfig) handleFeedFollowsCreate(user database.User) http.HandlerFunc {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		params := parameters{}
		err := decoder.Decode(&params)
		if err != nil {
			util.RespondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
			return
		}
		feedFollow, err := cfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
			ID:        uuid.New(),
			FeedID:    params.FeedID,
			UserID:    user.ID,
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
		})
		if err != nil {
			util.RespondWithError(w, http.StatusInternalServerError, "Couldn't create follow")
			return
		}
		util.RespondWithJSON(w, http.StatusOK, databaseFeedFollowToFeedFollow(feedFollow))
	}
}

func (cfg *apiConfig) handleFeedFollowDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idString := chi.URLParam(r, "feedFollowID")
		id, err := uuid.Parse(idString)
		if err != nil {
			util.RespondWithError(w, http.StatusInternalServerError, "No ID")
			return
		}

		num, err := cfg.DB.DeleteFeedFollow(r.Context(), id)
		if err != nil {
			util.RespondWithError(w, http.StatusInternalServerError, "Couldn't create follow")
			return
		}

		util.RespondWithJSON(w, http.StatusOK, fmt.Sprintf("You've deleted %d row(s)", num))

	}
}

func (cfg *apiConfig) handleFeedFollowGetAll(user database.User) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		feedFollows, err := cfg.DB.GetAllFeedFollows(r.Context(), user.ID)
		if err != nil {
			util.RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}

		var respFeedFollows []FeedFollow
		for _, ff := range feedFollows {
			respFeedFollows = append(respFeedFollows, databaseFeedFollowToFeedFollow(ff))
		}

		util.RespondWithJSON(w, http.StatusOK, respFeedFollows)
	}
}
