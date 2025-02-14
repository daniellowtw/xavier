package db

import (
	"xorm.io/xorm"
)

// TODO comments
type FeedSourceWithUnread struct {
	*FeedSource `xorm:"extends"`
	UnreadCount int64
	TotalCount  int64
}

func (FeedSourceWithUnread) TableName() string {
	return "feed_item"
}

// ListAllFeedsWithStatus returns the feeds along with its stats. Note: This can be quite slow.
func ListAllFeedsWithStats(engine *xorm.Engine) ([]*FeedSourceWithUnread, error) {
	var fs []*FeedSourceWithUnread
	err := engine.SQL(` select *, (select count(*) from feed_item as y where y.feed_id = s.id and read = 0) as unread_count, (select count(*) from feed_item as y where y.feed_id = s.id) as total_count from feed_source s where active = 1`).Find(&fs)
	if err != nil {
		return nil, err
	}
	return fs, nil
}
