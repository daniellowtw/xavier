package db

import (
	"fmt"
	"time"

	"xorm.io/core"
	"xorm.io/xorm"
)

type flaggedItem struct {
	FeedItemId int64     `xorm:"pk"`
	NewsItemId int64     `xorm:"pk"`
	Date       time.Time `xorm:"index"`
}

type FlagItemClient struct {
	Engine *xorm.Engine
}

func (c *FlagItemClient) FlagItem(newsID int64, feedID int64) (bool, error) {
	var x flaggedItem
	// Note: Ordering depends on order of struct
	ok, err := c.Engine.ID(core.NewPK(feedID, newsID)).Get(&x)
	if err != nil {
		return false, err
	}
	if ok {
		return true, nil
	}
	if _, err := c.Engine.InsertOne(&flaggedItem{
		NewsItemId: newsID,
		FeedItemId: feedID,
		Date:       time.Now(),
	}); err != nil {
		return false, fmt.Errorf("db: failed to flag news: %v", err)
	}
	return true, nil
}

func (c *FlagItemClient) RemoveFlag(newsID int64, feedID int64) (bool, error) {
	var x flaggedItem
	ok, err := c.Engine.ID(core.NewPK(feedID, newsID)).Get(&x)
	if err != nil {
		return false, err
	}
	if !ok {
		return false, fmt.Errorf("not found")
	}
	if _, err := c.Engine.Delete(&x); err != nil {
		return false, fmt.Errorf("db: failed to delete flagged news: %v", err)
	}
	return true, nil
}

func (c *FlagItemClient) GetAllFlagged(feedID int64) ([]int64, error) {
	var res []*flaggedItem
	err := c.Engine.Where(fmt.Sprintf("feed_item_id = %d", feedID)).Find(&res)
	if err != nil {
		return nil, err
	}
	var ret []int64
	for _, item := range res {
		ret = append(ret, item.NewsItemId)
	}
	return ret, nil
}
