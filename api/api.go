package api

import (
	"github.com/daniellowtw/xavier/db"
)

// Service is implements the API
type Service struct {
	*NewsService
	*FeedService
	*learningService
	*filterService
}

func NewService(client *db.Client) *Service {
	return &Service{
		NewsService: &NewsService{
			dbClient: client,
			saveItemClient: &db.SavedItemClient{
				Engine: client.Engine,
			},
		},
		FeedService: &FeedService{
			dbClient: client,
		},
		learningService: &learningService{
			dbClient: client,
			dataPointClient: &db.DataPointClient{
				Engine: client.Engine,
			},
		},
		filterService: &filterService{
			dbClient: client,
			flagItemClient: &db.FlagItemClient{
				Engine: client.Engine,
			},
		},
	}
}
