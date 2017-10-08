package db

import (
	"fmt"

	"github.com/go-xorm/xorm"
)

type RuleType int
type RuleAction int

const (
	Regex RuleType = iota
)

const (
	MarkAsRead RuleAction = iota
	Flag RuleAction = iota
	Delete
	Save
)

type FeedRule struct {
	FeedId       int64 `xorm:"index"`
	FilterType   RuleType
	FilterAction RuleAction
	FilterString string
}

type FeedRuleClient struct {
	Engine *xorm.Engine
}

func (c *FeedRuleClient) AddRegexRule(feedID int64, regex string, action RuleAction) error {
	f := &FeedRule{
		FilterType:   Regex,
		FilterString: regex,
		FeedId:       feedID,
		FilterAction: action,
	}
	if _, err := c.Engine.InsertOne(f); err != nil {
		return fmt.Errorf("cannot add filter: %v", err)
	}
	return nil
}

func (c *FeedRuleClient) GetRules() ([]*FeedRule, error) {
	var rules []*FeedRule
	if err := c.Engine.Find(&rules); err != nil {
		return nil, fmt.Errorf("cannot get filters: %v", err)
	}
	return rules, nil
}
