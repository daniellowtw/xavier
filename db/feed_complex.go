package db

import (
	"github.com/go-xorm/xorm"
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

// returns active feeds
func ListAllFeeds(engine *xorm.Engine) ([]*FeedSourceWithUnread, error) {
	var fs []*FeedSourceWithUnread
	err := engine.SQL(` select *, (select count(*) from feed_item as y where y.feed_id = s.id and read = 0) as unread_count, (select count(*) from feed_item as y where y.feed_id = s.id) as total_count from feed_source s where active = 1`).Find(&fs)
	if err != nil {
		return nil, err
	}
	return fs, nil
}

func populateUnprocessedNews(engine *xorm.Engine) error {
	for {
		var res []*FeedItem
		err := engine.SQL(`select id, feed_id from feed_item as y where y.Id not in (select news_item_id from data_point) and y.Id not in (select news_item_id from process_queue) LIMIT 200`).Find(&res)
		if err != nil {
			return err
		}
		if len(res) == 0 {
			break
		}
		var b []*ProcessQueue
		for _, r := range res {
			b = append(b, &ProcessQueue{FeedItemId: r.FeedId, NewsItemId: r.Id})
		}
		if _, err := engine.Insert(b); err != nil {
			return err
		}
	}
	return nil
}
