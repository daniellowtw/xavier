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
