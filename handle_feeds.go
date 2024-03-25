package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	util "github.com/pistolpeter/aggreRSS/internal"
	"github.com/pistolpeter/aggreRSS/internal/database"
)

func (cfg *apiConfig) handleFeedsCreate(user database.User) http.HandlerFunc {
	type parameters struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}
	type response struct {
		Feed       Feed       `json:"feed"`
		FeedFollow FeedFollow `json:"feed_follow"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		params := parameters{}
		err := decoder.Decode(&params)
		if err != nil {
			util.RespondWithError(w, http.StatusInternalServerError, "Couldn't Request body")
			return
		}

		feed, err := cfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
			ID:        uuid.New(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			Name:      params.Name,
			Url:       params.Url,
			UserID:    user.ID,
		})
		if err != nil {
			util.RespondWithError(w, http.StatusInternalServerError, "Couldn't create user")
			return
		}
		feedFollow, err := cfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
			ID:        uuid.New(),
			FeedID:    feed.ID,
			UserID:    user.ID,
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
		})

		resp := response{
			Feed:       databaseFeedsToFeeds(feed),
			FeedFollow: databaseFeedFollowToFeedFollow(&feedFollow),
		}

		util.RespondWithJSON(w, http.StatusOK, resp)
	}
}

func (cfg *apiConfig) handleFeedsGetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		feeds, err := cfg.DB.FeedsGetAll(r.Context())
		if err != nil {
			util.RespondWithError(w, http.StatusInternalServerError, "Could not retrieve feeds")
			return
		}
		var respFeeds []Feed
		for _, feed := range feeds {
			f := databaseFeedsToFeeds(feed)
			respFeeds = append(respFeeds, f)
		}
		util.RespondWithJSON(w, 200, respFeeds)
	}
}
