package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/pistolpeter/aggreRSS/internal/database"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	APIKey    string    `json:"api_key"`
}

type Feed struct {
	ID            uuid.UUID  `json:"id"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	Name          string     `json:"name"`
	URL           string     `json:"url"`
	UserID        uuid.UUID  `json:"user_id"`
	LastFetchedAT *time.Time `json:"last_fetched_at"`
}

type FeedFollow struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	FeedID    uuid.UUID `json:"feed_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func databaseUserToUser(user database.User) User {
	return User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Name:      user.Name,
		APIKey:    user.Apikey,
	}
}

func databaseFeedsToFeeds(feed database.Feed) Feed {
	return Feed{
		ID:            feed.ID,
		CreatedAt:     feed.CreatedAt,
		UpdatedAt:     feed.UpdatedAt,
		Name:          feed.Name,
		URL:           feed.Url,
		UserID:        feed.UserID,
		LastFetchedAT: &feed.LastFetchedAt.Time,
	}
}

func databaseFeedFollowToFeedFollow(feedFollow *database.FeedFollow) FeedFollow {
	return FeedFollow{
		ID:        feedFollow.ID,
		UserID:    feedFollow.UserID,
		FeedID:    feedFollow.FeedID,
		CreatedAt: feedFollow.CreatedAt,
		UpdatedAt: feedFollow.UpdatedAt,
	}
}
