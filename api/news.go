package api

import (
	"github.com/daniellowtw/xavier/db"
)

// NewsService is implements the API
type NewsService struct {
	dbClient *db.Client
}

type SearchParam struct {
	// false for all
	SearchMode string
	Limit      int
	FeedIDs    []int64
}

func (s *NewsService) Search(param SearchParam) ([]*db.ExtendedFeedItem, error) {
	var filters []db.Filter
	switch param.SearchMode {
	case "all":
	case "unread":
		filters = append(filters, db.FilterUnread())
	case "saved":
		filters = append(filters, db.FilterSaved())
	default:
	}
	if param.Limit == 0 {
		param.Limit = 100
	}
	filters = append(filters, db.FilterLimit(param.Limit))
	if len(param.FeedIDs) != 0 {
		filters = append(filters, db.FilterFeedIds(param.FeedIDs))
	}
	return s.dbClient.SearchNews(filters...)
}

func (s *NewsService) ListAllNewsForFeed(feedID int64) ([]*db.ExtendedFeedItem, error) {
	return s.dbClient.SearchNews(db.FilterFeedID(feedID))
}

func (s *NewsService) MarkAsRead(newsID int64) error {
	return s.dbClient.MarkAsRead(newsID)
}

// TODO: combine with single
func (s *NewsService) MarkAsReadMulti(newsID []int64) error {
	return s.dbClient.MarkAsReadMulti(newsID)
}

func (s *NewsService) ToggleNews(newsID int64, feedID int64) (bool, error) {
	return s.dbClient.ToggleSave(newsID, feedID)
}
