package cmd

import (
	"fmt"
	"net/url"

	"strconv"

	"github.com/daniellowtw/xavier/api"
	"github.com/daniellowtw/xavier/db"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
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
			return s.AddFeed(u)
		},
	}

	UpdateAllCmd = &cobra.Command{
		Use: "update-all",
		RunE: func(cmd *cobra.Command, args []string) error {
			_, err := s.UpdateAllFeeds()
			return err
		},
	}

	UpdateFeedCmd = &cobra.Command{
		Use: "update <id>",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return fmt.Errorf("Need one argument: <urL>")
			}
			u := args[0]
			i, err := strconv.ParseInt(u, 10, 64)
			if err != nil {
				return err
			}
			return s.UpdateFeed(i)
		},
	}

	ListAllCmd = &cobra.Command{
		Use: "list",
		RunE: func(cmd *cobra.Command, args []string) error {
			feeds, err := s.ListAllFeeds()
			if err != nil {
				return err
			}
			for _, f := range feeds {
				fmt.Println(f.Id, f.Title, f.Description, f.UnreadCount, f.TotalCount)
			}
			return nil
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
			return s.DeleteFeed(int64(i))
		},
	}

	LearnCmd = &cobra.Command{
		Use: "learn",
		RunE: func(*cobra.Command, []string) error {
			return s.LearnFromNewNews()
		},
	}

	AddNewToQueueCmd = &cobra.Command{
		Use: "add-to-queue",
		RunE: func(*cobra.Command, []string) error {
			return db.NewClient(e).PopulateProcessQueue()
		},
	}

	e *xorm.Engine
	s *api.Service
)

func init() {
	ee, err := xorm.NewEngine("sqlite3", "./foo.db")
	if err != nil {
		panic(err)
	}
	if err := ee.CreateTables(&db.FeedSource{}, &db.FeedItem{}, &db.DataPoint{}, &db.ProcessQueue{}); err != nil {
		panic(err)
	}
	e = ee
	s = api.NewService(db.NewClient(e))
	e.ShowSQL(true)
	e.Logger().SetLevel(core.LOG_ERR)
}
