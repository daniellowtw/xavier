package db

import (
	"fmt"
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
	total, err := c.e.Id(id).Count(&fs)
	if err != nil {
		return fmt.Errorf("db: %v", err)
	}
	if total == 0 {
		return fmt.Errorf("db: not found")
	}
	n, err := c.e.Where(fmt.Sprintf("feed_id = %d", id)).Delete(&FeedItem{})
	if err != nil {
		return fmt.Errorf("cannot find all db items: %v", err)
	}
	fmt.Printf("Deleted %d items\n", n)
	_, err = c.e.Delete(fs)
	if err != nil {
		return fmt.Errorf("cannot delete db source: %v", err)
	}
	return nil
}

func (c *Client) GetActiveFeedSources() ([]*FeedSourceWithUnread, error) {
	return ListAllFeeds(c.e)
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
func (c *Client) ListAllFeeds() error {
	fs, err := c.GetAllFeeds()
	if err != nil {
		return err
	}
	fmt.Println("id | title | unread | total")
	for _, f := range fs {
		fmt.Printf("%d %s %d %d\n", f.Id, f.Title, f.UnreadCount, f.TotalCount)
	}
	return nil
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

type filter func(s *xorm.Session) *xorm.Session

func FilterUnread() filter {
	return func(s *xorm.Session) *xorm.Session {
		return s.Where("read = 0")
	}
}

func FilterFeedID(id int64) filter {
	return func(s *xorm.Session) *xorm.Session {
		return s.Where(fmt.Sprintf("feed_id = %d", id))
	}
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

func (c *Client) SearchNews(filters ...filter) ([]*FeedItem, error) {
	println("searching news")
	var fs []*FeedItem
	starting := c.e.NewSession()
	for _, f := range filters {
		starting = f(starting)
	}
	err := starting.Find(&fs)
	if err != nil {
		return nil, err
	}
	return fs, nil
}
