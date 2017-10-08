package db

import (
	"fmt"

	"github.com/go-xorm/xorm"
)

type UserClassification int

const (
	NO_DECISION UserClassification = iota
	POSITIVE
	NEGATIVE
)

// DataPoint for doing our learning
type DataPoint struct {
	// auto public key
	Id         int64
	NewsItemId int64 `xorm:"index"`
	Keywords   []string
	Outcome    UserClassification
}

type ProcessQueue struct {
	// auto public key
	Id         int64
	FeedItemId int64 `xorm:"index"`
	NewsItemId int64 `xorm:"index"`
}

type DataPointClient struct {
	Engine *xorm.Engine
}

func (c *DataPointClient) GetProcessQueue() ([]*ProcessQueue, error) {
	//func (c *Client) GetProcessQueue() ([]*FeedItem, error) {
	var res []*ProcessQueue
	if err := c.Engine.Find(&res); err != nil {
		return nil, err
	}
	return res, nil
}

func (c *DataPointClient) MarkAsProcessed(q *ProcessQueue, keywords []string) error {
	if _, err := c.Engine.Insert(&DataPoint{NewsItemId: q.NewsItemId, Keywords: keywords}); err != nil {
		return err
	}
	// Note: this is not atomic. TODO for future
	if _, err := c.Engine.Delete(q); err != nil {
		return err
	}
	return nil
}

// look through datapoints and find news feed that is not a point yet.
func (c *DataPointClient) PopulateProcessQueue() error {
	return populateUnprocessedNews(c.Engine)
}

func (c *DataPointClient) GetDataPoint(newsID int64) (*DataPoint, error) {
	var res DataPoint
	ok, err := c.Engine.Where(fmt.Sprintf("news_item_id = %d", newsID)).Get(&res)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, fmt.Errorf("get: not found")
	}
	return &res, nil
}

func (c *DataPointClient) SaveDataPoint(point *DataPoint) error {
	_, err := c.Engine.Id(point.Id).Update(point)
	return err
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
