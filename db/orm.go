package db

import (
	"fmt"

	"strconv"
	"strings"
	"time"

	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
)

type Client struct {
	e *xorm.Engine
}

func NewClient(e *xorm.Engine) *Client {
	return &Client{e}
}

func (c *Client) UpdateFeedSource(f *FeedSource) error {
	_, err := c.e.Id(f.Id).Update(f)
	return err
}

func (c *Client) DeleteFeedSource(id int64) error {
	var fs FeedSource
	found, err := c.e.Id(id).Get(&fs)
	if err != nil {
		return fmt.Errorf("db: %v", err)
	}
	if !found {
		return fmt.Errorf("db: not found")
	}
	n, err := c.e.Where(fmt.Sprintf("feed_id = %d", id)).Delete(&FeedItem{})
	if err != nil {
		return fmt.Errorf("cannot find all db items: %v", err)
	}
	fmt.Printf("Deleted %d items\n", n)
	_, err = c.e.Delete(&fs)
	if err != nil {
		return fmt.Errorf("cannot delete db source: %v", err)
	}
	return nil
}

func (c *Client) GetActiveFeedSources() ([]*FeedSource, error) {
	var res []*FeedSource
	err := c.e.Where("active = 1").Find(&res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Client) GetActiveFeedSourcesWithStats() ([]*FeedSourceWithUnread, error) {
	return ListAllFeedsWithStats(c.e)
}

func (c *Client) GetNewsItem(id int64) (*FeedItem, error) {
	var i FeedItem
	ok, err := c.e.Id(id).Get(&i)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, fmt.Errorf("get: not found")
	}
	return &i, nil
}

func (c *Client) GetDataPoint(newsID int64) (*DataPoint, error) {
	var res DataPoint
	ok, err := c.e.Where(fmt.Sprintf("news_item_id = %d", newsID)).Get(&res)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, fmt.Errorf("get: not found")
	}
	return &res, nil
}

func (c *Client) SaveDataPoint(point *DataPoint) error {
	_, err := c.e.Id(point.Id).Update(point)
	return err
}

func (c *Client) GetProcessQueue() ([]*ProcessQueue, error) {
	//func (c *Client) GetProcessQueue() ([]*FeedItem, error) {
	var res []*ProcessQueue
	if err := c.e.Find(&res); err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Client) MarkAsProcessed(q *ProcessQueue, keywords []string) error {
	if _, err := c.e.Insert(&DataPoint{NewsItemId: q.NewsItemId, Keywords: keywords}); err != nil {
		return err
	}
	// Note: this is not atomic. TODO for future
	if _, err := c.e.Delete(q); err != nil {
		return err
	}
	return nil
}

// look through datapoints and find news feed that is not a point yet.
func (c *Client) PopulateProcessQueue() error {
	return populateUnprocessedNews(c.e)
}

func (c *Client) CheckWhetherSourceExist(url string) (bool, error) {
	tablePlaceholder := new(FeedSource)
	total, err := c.e.Where(fmt.Sprintf("url_source='%s'", url)).Count(tablePlaceholder)
	if err != nil {
		return false, fmt.Errorf("add: cannot check for existing feed: %v", err)
	}
	return total == 0, nil
}

func (c *Client) GetFeedSource(id int64) (*FeedSource, error) {
	var fs FeedSource
	ok, err := c.e.Id(id).Get(&fs)
	if !ok {
		return nil, fmt.Errorf("db: not found")
	}
	return &fs, err
}

func (c *Client) AddFeed(source *FeedSource) (int64, error) {
	id, err := c.e.Insert(source)
	if err != nil {
		return 0, fmt.Errorf("add: cannot insert to db: %v", err)
	}
	return id, nil
}
func (c *Client) AddNews(feedID int64, source *FeedItem) error {
	_, err := c.e.Insert(source)
	if err != nil {
		return fmt.Errorf("add: cannot insert news into db: %v", err)
	}
	return nil
}

func (c *Client) GetAllFeeds() (res []*FeedSource, err error) {
	err = c.e.Find(&res)
	return
}

func (c *Client) ListFeedItems(id string) error {
	fs, err := c.GetFeedItems(id)
	if err != nil {
		return err
	}
	fmt.Println("id | title | read? | published")
	for _, f := range fs {
		fmt.Printf("%d %s %t %v\n", f.Id, f.Title, f.Read, f.Published)
	}
	return nil
}

func (c *Client) GetFeedItems(id string) (res []*FeedItem, err error) {
	err = c.e.Where("feed_id = " + id).Find(&res)
	return
}

type Filter func(s *xorm.Session) *xorm.Session

func FilterUnread() Filter {
	return func(s *xorm.Session) *xorm.Session {
		return s.Where("read = 0")
	}
}
func FilterLimit(n int) Filter {
	return func(s *xorm.Session) *xorm.Session {
		return s.Limit(n)
	}
}

func FilterFeedIds(feedIDs []int64) Filter {
	return func(s *xorm.Session) *xorm.Session {
		var final []string
		for _, i := range feedIDs {
			final = append(final, strconv.FormatInt(i, 10))
		}
		return s.Where(fmt.Sprintf("feed_id in (%v)", strings.Join(final, ",")))
	}
}

func FilterFeedID(id int64) Filter {
	return func(s *xorm.Session) *xorm.Session {
		return s.Where(fmt.Sprintf("feed_id = %d", id))
	}
}

func FilterGUID(guid string) Filter {
	return func(s *xorm.Session) *xorm.Session {
		return s.Where(fmt.Sprintf("guid = '%s'", guid))
	}
}

func FilterSaved() Filter {
	return func(s *xorm.Session) *xorm.Session {
		// TODO this is tightly coupled to the name of the feed_item table
		return s.Where("feed_item.id in (select news_item_id from saved_item)")
	}
}

func (c *Client) MarkAsReadMulti(newsID []int64) error {
	ids := strings.Replace(strings.Trim(fmt.Sprintf("%v", newsID), "[]"), " ", ",", -1)
	query := fmt.Sprintf(`update %s set read = 1 where id in (%s)`, "feed_item", ids)
	if _, err := c.e.Exec(query); err != nil {
		return fmt.Errorf("db: failed to update news: %v", err)
	}
	return nil
}

func (c *Client) MarkAsRead(newsID int64) error {
	news := new(FeedItem)
	ok, err := c.e.Id(newsID).Get(news)
	if err != nil {
		return err
	}
	if !ok {
		return fmt.Errorf("db: not found")
	}
	news.Read = true
	if _, err := c.e.Id(newsID).Cols("read").Update(news); err != nil {
		return fmt.Errorf("db: failed to update news: %v", err)
	}
	return nil
}

func (c *Client) ToggleSave(newsID int64, feedID int64) (bool, error) {
	var x SavedItem
	// Note: Ordering depends on order of struct
	ok, err := c.e.Id(core.NewPK(feedID, newsID)).Get(&x)
	if err != nil {
		return false, err
	}
	if ok {
		_, err := c.e.Delete(&x)
		if err != nil {
			return false, fmt.Errorf("db: failed to delete saved news: %v", err)
		}
		return false, nil
	}
	if _, err := c.e.InsertOne(&SavedItem{
		NewsItemId: newsID,
		FeedItemId: feedID,
		Date:       time.Now(),
	}); err != nil {
		return false, fmt.Errorf("db: failed to save news: %v", err)
	}
	return true, nil
}

func (c *Client) SearchNews(filters ...Filter) ([]*ExtendedFeedItem, error) {
	var fs []*ExtendedFeedItem
	savedItemStr := `case when saved_item.news_item_id = feed_item.id then 1 else 0 end as is_saved` // used to figure out if something is in the other table or not
	starting := c.e.NewSession()
	starting = starting.SetExpr("foo", "case 1 end")
	starting = starting.Select("feed_item.*, data_point.outcome as classification, " + savedItemStr)
	starting = starting.Join("left", "saved_item", "saved_item.news_item_id = feed_item.id")
	starting = starting.Join("left", "data_point", "data_point.news_item_id = feed_item.id")
	for _, f := range filters {
		starting = f(starting)
	}
	starting = starting.Desc("id")
	err := starting.Find(&fs)
	if err != nil {
		return nil, err
	}
	return fs, nil
}
