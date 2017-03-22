package feed

import (
	"time"

	"encoding/json"

	_ "github.com/mattn/go-sqlite3"
	"github.com/mmcdole/gofeed"
)

type FeedSource struct {
	Id          int64
	Active      bool
	Title       string
	Description string
	UrlSource   string `xorm:"unique"`
	UnreadCount int
	TotalCount  int
	Status      string // TODO
	LastUpdated *time.Time
	LastChecked time.Time
	Created     time.Time `xorm:"created"`
}

type FeedItem struct {
	Id     int64
	FeedId int64 `xorm:"index"`
	Read   bool  `xorm:"index"`

	Title      string `xorm:"index"`
	Published  *time.Time
	LinkHref   string `xorm:"unique"`
	Content    string
	AuthorName string
	Category   []string `xorm:"index"`
	Guid       string
	Enclosure  string
	Custom     map[string]string
}

func ToFeedItem(feedID int64, i *gofeed.Item) *FeedItem {
	encl, _ := json.Marshal(i.Enclosures)
	author := ""
	if i.Author != nil {
		author = i.Author.Name
	}
	return &FeedItem{
		FeedId:     feedID,
		Title:      i.Title,
		AuthorName: author,
		Category:   i.Categories,
		Content:    i.Content,
		LinkHref:   i.Link,
		Published:  i.PublishedParsed,
		Guid:       i.GUID,
		Enclosure:  string(encl),
		Custom:     i.Custom,
	}
}
