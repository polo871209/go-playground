package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/polo871209/chi-playground/internal/database"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	APIKey    string    `json:"api_key"`
}

func databaseUserToUser(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name:      dbUser.Name,
		APIKey:    dbUser.ApiKey,
	}
}

type Feed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	UserID    uuid.UUID `json:"user_id"`
}

func databaseFeedToFeed(dbFeed database.Feed) Feed {
	return Feed{
		ID:        dbFeed.ID,
		CreatedAt: dbFeed.CreatedAt,
		UpdatedAt: dbFeed.UpdatedAt,
		Name:      dbFeed.Name,
		Url:       dbFeed.Url,
		UserID:    dbFeed.UserID,
	}
}

func databaseFeedsToFeeds(dbFeeds []database.Feed) []Feed {
	feeds := []Feed{}
	for _, dbFeed := range dbFeeds {
		feeds = append(feeds, databaseFeedToFeed(dbFeed))
	}
	return feeds
}

type FeedFellow struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    uuid.UUID `json:"user_id"`
	FeedID    uuid.UUID `json:"feed_id"`
}

func databaseFeedFollowToFeedFollow(dbFeedFellow database.FeedFollow) FeedFellow {
	return FeedFellow{
		ID:        dbFeedFellow.ID,
		CreatedAt: dbFeedFellow.CreatedAt,
		UpdatedAt: dbFeedFellow.UpdatedAt,
		UserID:    dbFeedFellow.UserID,
		FeedID:    dbFeedFellow.FeedID,
	}
}

func databaseFeedFollowsToFeedFollows(dbFeedsFellows []database.FeedFollow) []FeedFellow {
	feedFellows := []FeedFellow{}
	for _, dbFeedFellow := range dbFeedsFellows {
		feedFellows = append(feedFellows, databaseFeedFollowToFeedFollow(dbFeedFellow))
	}
	return feedFellows
}
