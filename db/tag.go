package db

import (
	"github.com/go-xorm/xorm"
	"fmt"
)

type TaggedItem struct {
	FeedId int64    `xorm:"pk"`
	NewsId int64    `xorm:"pk"`
	Tags   []string `xorm:"index"`
}

type TagItemClient struct {
	Engine *xorm.Engine
}

func (c *TagItemClient) SetTags(feedID int64, newsId int64, tags []string) error {
	i := &TaggedItem{
		NewsId: newsId,
		FeedId: feedID,
		Tags: tags,
	}
	// try to insert
	if _, err := c.Engine.InsertOne(i); err != nil {
		return fmt.Errorf("cannot add filter: %v", err)
	}
	return nil
}

// TODO: Getter for tags
