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

func NewService(e *db.Client) *Service {
	return &Service{
		NewsService: &NewsService{
			dbClient: e,
		},
		FeedService: &FeedService{
			dbClient: e,
		},
		learningService: &learningService{
			dbClient: e,
		},
		filterService: &filterService{
			dbClient: e,
		},
	}
}
