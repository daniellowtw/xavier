package api

import (
	"strings"

	"github.com/daniellowtw/xavier/db"
)

type learningService struct {
	dbClient *db.Client
	dataPointClient *db.DataPointClient
}

func (s *learningService) LearnFromNewNews() error {
	queue, err := s.dataPointClient.GetProcessQueue()
	if err != nil {
		return err
	}
	for _, q := range queue {
		n, err := s.dbClient.GetNewsItem(q.NewsItemId)
		if err != nil {
			return err
		}
		var keywords []string
		keywords = append(keywords, extractKeywords(n.Title)...)
		// description can be chatty, if it is, ignore to reduce noise
		if len(n.Description) <= 500 {
			keywords = append(keywords, extractKeywords(n.Description)...)
		}
		keywords = append(keywords, extractKeywords(strings.Join(n.Category, " "))...)

		seen := map[string]struct{}{}
		var filtered []string
		for _, s := range keywords {
			s = strings.ToLower(s)
			if _, ok := seen[s]; ok {
				continue
			}
			seen[s] = struct{}{}
			filtered = append(filtered, s)
		}
		if err := s.dataPointClient.MarkAsProcessed(q, filtered); err != nil {
			return err
		}
	}
	return nil
}

func (s *learningService) HumanClassification(newsID int64, outcome db.UserClassification) error {
	d, err := s.dataPointClient.GetDataPoint(newsID)
	if err != nil {
		return err
	}
	d.Outcome = outcome
	return s.dataPointClient.SaveDataPoint(d)
}
