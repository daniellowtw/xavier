package db

import (
	"github.com/go-xorm/xorm"
)

// TODO comments
type FeedSourceWithUnread struct {
	*FeedSource `xorm:"extends"`
	UnreadCount            int64
}

func (FeedSourceWithUnread) TableName() string {
	return "feed_item"
}

// returns active feeds
func ListAllFeeds(engine *xorm.Engine) ([]*FeedSourceWithUnread, error) {
	var fs []*FeedSourceWithUnread
	err := engine.SQL(`select *, (select count(*) from feed_item as y where y.feed_id = s.id and read = 0) as unread_count from feed_source s where active = 1`).Find(&fs)
	if err != nil {
		return nil, err
	}
	return fs, nil
}

