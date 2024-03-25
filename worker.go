package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

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
						fmt.Println("Error Fetching feed", err)
						return
					}

					log.Println(data.Channel.Title)
				}(feed.Url)
				db.MarkFeedFetched(context.TODO(), feed.ID)
			}
			wg.Wait()
		}
	}
}
