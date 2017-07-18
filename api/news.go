package api

import (
	"github.com/daniellowtw/xavier/db"
)

// NewsService is implements the API
type NewsService struct {
	dbClient    *db.Client
}

type SearchParam struct {
	// false for all
	IncludeRead bool
}

func (s *NewsService) Search(param SearchParam) ([]*db.FeedItem, error) {
	if param.IncludeRead {
		// search all if we want to find read as well
		return s.dbClient.SearchNews()
	}
	return s.dbClient.SearchNews(db.FilterUnread())
}

func (s *NewsService) ListAllNewsForFeed(feedID int64) ([]*db.FeedItem, error) {
	return s.dbClient.SearchNews(db.FilterFeedID(feedID))
}

func (s *NewsService) MarkAsRead(newsID int64) error {
	return s.dbClient.MarkAsRead(newsID)
}
