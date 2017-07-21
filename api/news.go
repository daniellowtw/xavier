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
	Limit int
	FeedIDs []int64
}

func (s *NewsService) Search(param SearchParam) ([]*db.FeedItem, error) {
	var filters []db.Filter
	if !param.IncludeRead {
		filters = append(filters, db.FilterUnread())
	}
	if param.Limit != 0 {
		filters = append(filters, db.FilterLimit(param.Limit))
	}
	if len(param.FeedIDs) != 0 {
		filters = append(filters, db.FilterFeedIds(param.FeedIDs))
	}
	return s.dbClient.SearchNews(filters...)
}

func (s *NewsService) ListAllNewsForFeed(feedID int64) ([]*db.FeedItem, error) {
	return s.dbClient.SearchNews(db.FilterFeedID(feedID))
}

func (s *NewsService) MarkAsRead(newsID int64) error {
	return s.dbClient.MarkAsRead(newsID)
}
