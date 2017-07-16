package cmd

import (
	"fmt"
	"net/url"

	"strconv"

	"github.com/daniellowtw/xavier/api"
	"github.com/daniellowtw/xavier/db"
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
			return s.UpdateAllFeeds()
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

	e *xorm.Engine
	s *api.Service
)

func init() {
	ee, err := xorm.NewEngine("sqlite3", "./foo.db")
	if err != nil {
		panic(err)
	}
	if err := ee.CreateTables(&db.FeedSource{}, &db.FeedItem{}); err != nil {
		panic(err)
	}
	e = ee
	s = &api.Service{StoreEngine: e}
}
