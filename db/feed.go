package db

import (
	"time"

	"encoding/json"

	_ "github.com/mattn/go-sqlite3"
	"github.com/mmcdole/gofeed"
)

type FeedSource struct {
	// auto public key
	Id          int64
	Active      bool
	Title       string
	Description string
	UrlSource   string `xorm:"unique"`
	Status      string // TODO
	LastUpdated *time.Time
	LastChecked time.Time
	Created     time.Time `xorm:"created"`
	FavIcon     string
}

type FeedItem struct {
	// auto public key
	Id     int64 `xorm:"index"`
	FeedId int64 `xorm:"index"`
	Read   bool  `xorm:"index"`

	Title       string `xorm:"index"`
	Published   *time.Time
	LinkHref    string `xorm:"unique"`
	Description string
	Content     string
	AuthorName  string
	Category    []string `xorm:"index"`
	Guid        string
	Enclosure   string
	Custom      map[string]string
}

func ToFeedItem(feedID int64, i *gofeed.Item) *FeedItem {
	encl, _ := json.Marshal(i.Enclosures)
	author := ""
	if i.Author != nil {
		author = i.Author.Name
	}
	return &FeedItem{
		FeedId:      feedID,
		Title:       i.Title,
		AuthorName:  author,
		Category:    i.Categories,
		Description: i.Description,
		Content:     i.Content,
		LinkHref:    i.Link,
		Published:   i.PublishedParsed,
		Guid:        i.GUID,
		Enclosure:   string(encl),
		Custom:      i.Custom,
	}
}
