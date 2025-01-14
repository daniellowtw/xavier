package db

import (
	"fmt"
	"time"

	"xorm.io/core"
	"xorm.io/xorm"
)

type SavedItem struct {
	FeedItemId int64     `xorm:"pk"`
	NewsItemId int64     `xorm:"pk"`
	Date       time.Time `xorm:"index"`
}

type SavedItemClient struct {
	Engine *xorm.Engine
}

func (c *SavedItemClient) ToggleSave(newsID int64, feedID int64) (bool, error) {
	var x SavedItem
	// Note: Ordering depends on order of struct
	ok, err := c.Engine.ID(core.NewPK(feedID, newsID)).Get(&x)
	if err != nil {
		return false, err
	}
	if ok {
		_, err := c.Engine.Delete(&x)
		if err != nil {
			return false, fmt.Errorf("db: failed to delete saved news: %v", err)
		}
		return false, nil
	}
	if _, err := c.Engine.InsertOne(&SavedItem{
		NewsItemId: newsID,
		FeedItemId: feedID,
		Date:       time.Now(),
	}); err != nil {
		return false, fmt.Errorf("db: failed to save news: %v", err)
	}
	return true, nil
}

// Save does nothing if the item is already saved.
func (c *SavedItemClient) Save(newsID int64, feedID int64) error {
	var x SavedItem
	// Note: Ordering depends on order of struct
	ok, err := c.Engine.ID(core.NewPK(feedID, newsID)).Get(&x)
	if err != nil {
		return err
	}
	if ok {
		return nil
	}
	if _, err := c.Engine.InsertOne(&SavedItem{
		NewsItemId: newsID,
		FeedItemId: feedID,
		Date:       time.Now(),
	}); err != nil {
		return fmt.Errorf("db: failed to save news: %v", err)
	}
	return nil
}

func (c *SavedItemClient) RemoveFromSaved(newsID int64, feedID int64) error {
	var x SavedItem
	// Note: Ordering depends on order of struct
	ok, err := c.Engine.ID(core.NewPK(feedID, newsID)).Get(&x)
	if err != nil {
		return err
	}
	if !ok {
		return fmt.Errorf("not found")
	}
	if _, err := c.Engine.Delete(&x); err != nil {
		return fmt.Errorf("db: failed to delete saved news: %v", err)
	}
	return nil
}
