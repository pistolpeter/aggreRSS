package main

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/pistolpeter/aggreRSS/internal/database"
)

func worker(interval time.Duration, fetchCount int, db *database.Queries) {
	ticker := time.NewTicker(interval)

	for {
		select {
		case <-ticker.C:
			feeds, err := db.GetNextFeedsToFetch(context.TODO(), 10)
			if err != nil {
				fmt.Println("Error Fetching feed", err)

				return
			}
			var wg sync.WaitGroup

			for _, feed := range feeds {
				wg.Add(1)

				go func(url string) {
					defer wg.Done()

					data, err := FetchFeedData(url)
					if err != nil {
						fmt.Println("Error Fetching Posts: ", err)
						return
					}

					for _, item := range data.Channel.Items {
						var publishedAt sql.NullTime
						if itemTime, err := time.Parse(time.RFC1123Z, item.PubDate); err == nil {
							publishedAt = sql.NullTime{
								Time:  itemTime,
								Valid: true,
							}
						}
						_, err := db.PostCreate(context.TODO(), database.PostCreateParams{
							ID:        uuid.New(),
							CreatedAt: time.Now(),
							UpdatedAt: time.Now(),
							Title:     item.Title,
							Url:       item.Link,
							Description: sql.NullString{
								String: item.Description,
								Valid:  true,
							},
							PublishedAt: publishedAt,
							FeedID:      feed.ID,
						})
						if err != nil {
							if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
								continue
							}
							fmt.Println("error on post creatation: ", err)
							continue
						}
					}
				}(feed.Url)
				db.MarkFeedFetched(context.TODO(), feed.ID)
			}
			wg.Wait()
		}
	}
}
