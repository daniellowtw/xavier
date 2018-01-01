package api

import (
	"github.com/daniellowtw/xavier/db"
	client2 "github.com/daniellowtw/xavier/client"
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
			httpClient: client2.NewDefaultClient(),
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
