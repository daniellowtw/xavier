package api

import (
	"fmt"
	"regexp"

	"github.com/daniellowtw/xavier/db"
)

type filterService struct {
	dbClient *db.Client

	flagItemClient  *db.FlagItemClient
	savedItemClient *db.SavedItemClient
}

func (s *filterService) AddRegexFilter(feedID int64, regexString string, action db.RuleAction) error {
	if _, err := regexp.Compile(regexString); err != nil {
		return fmt.Errorf("invalid regex: %v", err)
	}
	return s.dbClient.AddRegexRule(feedID, regexString, action)
}

func (s *filterService) RunFilter() error {
	rules, err := s.dbClient.GetRules()
	if err != nil {
		return err
	}
	for _, r := range rules {
		switch r.FilterType {
		case db.Regex:
			s.runRegexRule(r)
		}
	}
	return nil
}

func (s *filterService) runRegexRule(r *db.FeedRule) error {
	re, err := regexp.Compile(r.FilterString)
	if err != nil {
		return err
	}
	news, err := s.dbClient.SearchNews(db.FilterFeedID(r.FeedId), db.FilterUnread())
	if err != nil {
		return err
	}
	for _, n := range news {
		match := false
		match = match || re.MatchString(n.Title)
		match = match || re.MatchString(n.Description)
		match = match || re.MatchString(n.Content)
		if !match {
			continue
		}
		s.applyAction(n, r.FilterAction)
	}
	return nil
}

func (s *filterService) applyAction(news *db.ExtendedFeedItem, action db.RuleAction) {
	switch action {
	case db.Flag:
		if _, err := s.flagItemClient.FlagItem(news.Id, news.FeedId); err != nil {
			fmt.Println(err)
		}
	case db.Delete:
		//TODO: Implement me
	case db.MarkAsRead:
		s.dbClient.MarkAsRead(news.Id)
	case db.Save:
		s.savedItemClient.Save(news.Id, news.FeedId)
	}
}
