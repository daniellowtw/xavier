package api

import (
	"fmt"
	"net/http"
	"time"

	"log"

	"encoding/json"

	"github.com/daniellowtw/xavier/client"
	"github.com/daniellowtw/xavier/db"
	"github.com/mmcdole/gofeed"
)

// FeedService is implements the API relating to feeds
type FeedService struct {
	dbClient *db.Client
}

func (s *FeedService) ListAllFeeds() ([]*db.FeedSourceWithUnread, error) {
	return s.dbClient.GetActiveFeedSources()
}

// AddFeed and then update it
func (s *FeedService) AddFeed(url string) error {
	exists, err := s.dbClient.CheckWhetherSourceExist(url)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("api: feed %s already exists", url)
	}
	fp := gofeed.NewParser()
	fp.Client = client.New()
	f, err := fp.ParseURL(url)
	if err != nil {
		return fmt.Errorf("add: cannot parse URL: %v", err)
	}
	favIcon, err := getFavIcon(f.Link)
	if err != nil {
		return err
	}
	item := &db.FeedSource{
		Title:       f.Title,
		UrlSource:   url,
		Description: f.Description,
		LastUpdated: f.UpdatedParsed,
		Active:      true,
		LastChecked: time.Now(),
		FavIcon:     favIcon,
	}
	feedID, err := s.dbClient.AddFeed(item)
	if err != nil {
		return err
	}
	for _, i := range f.Items {
		fixFeedItem(i)
		err := s.dbClient.AddNews(feedID, db.ToFeedItem(item.Id, i))
		if err != nil {
			return err
		}
	}
	log.Printf("Added %d items", len(f.Items))
	return err
}

// check for update and then store news
func (s *FeedService) updateFeedFromURL(f *db.FeedSource) (int, error) {
	fp := gofeed.NewParser()
	fp.Client = client.New()
	gf, err := fp.ParseURL(f.UrlSource)
	if err != nil {
		return 0, err
	}
	lastUpdated := f.LastChecked
	var updatedItemCount int
	for _, i := range gf.Items {
		if i.PublishedParsed != nil {
			if i.PublishedParsed.Before(lastUpdated) {
				continue
			}
		} else if i.UpdatedParsed != nil {
			if i.UpdatedParsed.Before(lastUpdated) {
				continue
			}
		} else {
			fmt.Println("Cannot parsed published time nor updated time. Feed will most likely have duplicated items.")
		}
		fixFeedItem(i)
		err := s.dbClient.AddNews(f.Id, db.ToFeedItem(f.Id, i))
		if err != nil {
			return 0, err
		}
		updatedItemCount++
	}
	fmt.Printf("Updated db %d - %d items added.\n", f.Id, updatedItemCount)
	f.LastChecked = time.Now()
	s.dbClient.UpdateFeedSource(f)
	return updatedItemCount, err
}

func (s *FeedService) UpdateFeed(feedID int64) error {
	f, err := s.dbClient.GetFeedSource(feedID)
	if err != nil {
		return fmt.Errorf("api: cannot get feed source: %v", err)
	}
	_, err = s.updateFeedFromURL(f)
	return err
}

func (s *FeedService) DeleteFeed(id int64) error {
	s.dbClient.DeleteFeedSource(id)
	return nil
}

func (s *FeedService) DebugFeed(id int64) ([]byte, error) {
	f, err := s.dbClient.GetFeedSource(id)
	if err != nil {
		return nil, err
	}

	fp := gofeed.NewParser()
	fp.Client = client.New()
	gf, err := fp.ParseURL(f.UrlSource)
	if err != nil {
		return nil, err
	}
	for _, i := range gf.Items {
		fixFeedItem(i)
	}
	return json.MarshalIndent(gf, "", "\t")
}

func (s *FeedService) UpdateAllFeeds() (int, error) {
	fs, err := s.dbClient.GetActiveFeedSources()
	if err != nil {
		return 0, err
	}
	total := 0
	fmt.Printf("Found %d feeds to update\n", len(fs))
	for _, f := range fs {
		n, err := s.updateFeedFromURL(f.FeedSource)
		if err != nil {
			return 0, err
		}
		total += n
		fmt.Printf("Updated db %s\n", f.Title)
	}
	return total, nil
}

func writeErr(w http.ResponseWriter, statusCode int, err error) {
	if err == nil {
		return
	}
	w.WriteHeader(statusCode)
	w.Write([]byte(err.Error()))
}

// parse the content module http://web.resource.org/rss/1.0/modules/content/
func fixFeedItem(i *gofeed.Item) {
	t1, ok := i.Extensions["content"]
	if !ok {
		return
	}
	t2, ok := t1["encoded"]
	if !ok {
		return
	}
	if len(t2) != 1 {
		return
	}
	i.Content = t2[0].Value
}
