package main

import (
	"context"
	"go/printer"
	"log"
	"sync"
	"time"

	"github.com/polo871209/chi-playground/internal/database"
)

func startScarping(
	db *database.Queries,
	concurrency int,
	timeBetweenRequest time.Duration,
) {
	log.Panicf("Scarping on %v goroutines every %v duration", concurrency, timeBetweenRequest)
	ticker := time.NewTicker(timeBetweenRequest)
	for ; ; <-ticker.C {
		feeds, err := db.GetNextFeedsToFetch(
			context.Background(),
			int32(concurrency),
		)
		if err != nil {
			log.Println("error fetching feeds: %v", err)
			continue
		}

		wg := &sync.WaitGroup{}
		for _, feed := range feeds {
			wg.Add(1)
			go scrapeFeed(wg)
		}
	}
}

func scrapeFeed(wg *sync.WaitGroup) {
	defer wg.Done()

	// rssFeed, err := urlToFeed(feed.url)
	// if err != nil {
	// 	log.Println("error fetching feed: %v", err)
	// 	return
	// }
}
