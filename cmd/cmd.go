package cmd

import (
	"fmt"
	"net/url"
	"time"

	"strconv"

	"github.com/daniellowtw/xavier/client"
	"github.com/daniellowtw/xavier/feed"
	"github.com/go-xorm/xorm"
	"github.com/mmcdole/gofeed"
	"github.com/spf13/cobra"
)

var (
	AddCmd = &cobra.Command{
		Use: "add <url>",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return fmt.Errorf("Need one argument: <urL>")
			}
			u := args[0]
			_, err := url.Parse(u)
			if err != nil {
				return err
			}
			return addFeed(e, u)
		},
	}

	UpdateAllCmd = &cobra.Command{
		Use: "update-all",
		RunE: func(cmd *cobra.Command, args []string) error {
			return updateAllFeeds(e)
		},
	}
	ListAllCmd = &cobra.Command{
		Use: "list",
		RunE: func(cmd *cobra.Command, args []string) error {
			return listAllFeeds(e)
		},
	}

	DeleteCmd = &cobra.Command{
		Use: "delete <id>",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return fmt.Errorf("Need one argument: <urL>")
			}
			u := args[0]
			i, err := strconv.Atoi(u)
			if err != nil {
				return err
			}
			return deleteFeed(e, i)
		},
	}
	e *xorm.Engine
)

func init() {
	ee, err := xorm.NewEngine("sqlite3", "./foo.db")
	if err != nil {
		panic(err)
	}
	if err := ee.CreateTables(&feed.FeedSource{}, &feed.FeedItem{}); err != nil {
		panic(err)
	}
	e = ee
}

func deleteFeed(e *xorm.Engine, id int) error {
	var fs []*feed.FeedSource
	err := e.Id(id).Find(&fs)
	if err != nil {
		return fmt.Errorf("no such feed: %v\n", err)
	}
	n, err := e.Where(fmt.Sprintf("feed_id = %d", id)).Delete(&feed.FeedItem{})
	if err != nil {
		return fmt.Errorf("cannot find all feed items: %v\n", err)
	}
	fmt.Printf("Deleted %d items\n", n)
	_, err = e.Delete(fs[0])
	if err != nil {
		return fmt.Errorf("cannot detel feed source: %v\n", err)
	}
	fmt.Printf("Deleted feed %s\n", fs[0].Title)
	return nil
}
func listAllFeeds(e *xorm.Engine) error {
	var fs []*feed.FeedSource
	err := e.Find(&fs)
	if err != nil {
		return err
	}
	fmt.Println("id | title | unread | total")
	for _, f := range fs {
		fmt.Printf("%d %s %d %d\n", f.Id, f.Title, f.UnreadCount, f.TotalCount)
	}
	return nil
}

func updateAllFeeds(e *xorm.Engine) error {
	var fs []*feed.FeedSource
	err := e.Where("active = 1").Find(&fs)
	if err != nil {
		return err
	}
	fmt.Printf("Found %d feeds to update\n", len(fs))
	for _, f := range fs {
		err := updateFeedFromURL(f, e)
		if err != nil {
			return err
		}
		fmt.Printf("Updated feed %s\n", f.Title)
	}
	return nil
}

func updateFeed(e *xorm.Engine, feedId int64) error {
	var fs []*feed.FeedSource
	err := e.Id(feedId).Find(&fs)
	if err != nil {
		return fmt.Errorf("cannot find feed: %v\n", err)
	}
	if len(fs) == 0 {
		return fmt.Errorf("no feed found\n")
	}
	f := fs[0]
	return updateFeedFromURL(f, e)
}

func updateFeedFromURL(f *feed.FeedSource, e *xorm.Engine) error {
	fp := gofeed.NewParser()
	fp.Client = client.New()
	gf, err := fp.ParseURL(f.UrlSource)
	if err != nil {
		return err
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
		updatedItemCount++
		_, err := e.Insert(feed.ToFeedItem(f.Id, i))
		if err != nil {
			return err
		}
	}
	f.LastChecked = time.Now()
	f.UnreadCount += updatedItemCount
	f.TotalCount += updatedItemCount
	_, err = e.Id(f.Id).Update(f)
	fmt.Printf("Updated feed %d - %d items added.\n", f.Id, updatedItemCount)
	return err
}

func addFeed(e *xorm.Engine, url string) error {
	c := client.New()
	fp := gofeed.NewParser()
	fp.Client = c
	f, err := fp.ParseURL(url)
	if err != nil {
		return fmt.Errorf("add: cannot parse URL: %v", err)
	}
	item := &feed.FeedSource{
		Title:       f.Title,
		UrlSource:   url,
		Description: f.Description,
		LastUpdated: f.UpdatedParsed,
		Active:      true,
		LastChecked: time.Now(),
	}
	_, err = e.Insert(item)
	if err != nil {
		return err
	}
	for _, i := range f.Items {
		_, err := e.Insert(feed.ToFeedItem(item.Id, i))
		if err != nil {
			return err
		}
	}
	item.TotalCount += len(f.Items)
	item.UnreadCount += len(f.Items)
	if _, err := e.Id(item.Id).Update(item); err != nil {
		return err
	}
	fmt.Printf("Added %d items for feed %s\n", len(f.Items), f.Title)
	return err
}
