package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/daniellowtw/xavier/client"
	"github.com/daniellowtw/xavier/db"
	"github.com/davecgh/go-spew/spew"
	"github.com/go-xorm/xorm"
	"github.com/gorilla/mux"
	"github.com/mmcdole/gofeed"
)

// Service is implements the API
type Service struct {
	StoreEngine *xorm.Engine
	dbClient    db.Client
}

func (s *Service) AddFeed(url string) error {
	fp := gofeed.NewParser()
	fp.Client = client.New()
	f, err := fp.ParseURL(url)
	if err != nil {
		return fmt.Errorf("add: cannot parse URL: %v", err)
	}
	var existing []*db.FeedSource
	if err := s.StoreEngine.Where(fmt.Sprintf("url_source='%s'", url)).Find(&existing); err != nil {
		return fmt.Errorf("add: cannot check for existing feed: %v", err)
	}
	if len(existing) > 0 {
		return fmt.Errorf("add: URL already exist: %s", url)
	}
	item := &db.FeedSource{
		Title:       f.Title,
		UrlSource:   url,
		Description: f.Description,
		LastUpdated: f.UpdatedParsed,
		Active:      true,
		LastChecked: time.Now(),
	}
	_, err = s.StoreEngine.Insert(item)
	if err != nil {
		return err
	}
	for _, i := range f.Items {
		_, err := s.StoreEngine.Insert(db.ToFeedItem(item.Id, i))
		if err != nil {
			return err
		}
	}
	item.TotalCount += len(f.Items)
	item.UnreadCount += len(f.Items)
	if _, err := s.StoreEngine.Id(item.Id).Update(item); err != nil {
		return err
	}
	fmt.Printf("Added %d items for db %s\n", len(f.Items), f.Title)
	return err
}

func (s *Service) updateFeedFromURL(f *db.FeedSource) error {
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
		_, err := s.StoreEngine.Insert(db.ToFeedItem(f.Id, i))
		if err != nil {
			return err
		}
	}
	f.LastChecked = time.Now()
	f.UnreadCount += updatedItemCount
	f.TotalCount += updatedItemCount
	_, err = s.StoreEngine.Id(f.Id).Update(f)
	fmt.Printf("Updated db %d - %d items added.\n", f.Id, updatedItemCount)
	return err
}

func (s *Service) UpdateFeed(feedID int64) error {
	var fs []*db.FeedSource
	err := s.StoreEngine.Id(feedID).Find(&fs)
	if err != nil {
		return fmt.Errorf("cannot find db: %v", err)
	}
	if len(fs) == 0 {
		return fmt.Errorf("no db found")
	}
	f := fs[0]
	return s.updateFeedFromURL(f)
}

func (s *Service) DeleteFeed(id int64) error {
	var fs []*db.FeedSource
	err := s.StoreEngine.Id(id).Find(&fs)
	if err != nil {
		return fmt.Errorf("no such db: %v", err)
	}
	n, err := s.StoreEngine.Where(fmt.Sprintf("feed_id = %d", id)).Delete(&db.FeedItem{})
	if err != nil {
		return fmt.Errorf("cannot find all db items: %v", err)
	}
	fmt.Printf("Deleted %d items\n", n)
	_, err = s.StoreEngine.Delete(fs[0])
	if err != nil {
		return fmt.Errorf("cannot delete db source: %v", err)
	}
	fmt.Printf("Deleted db %s\n", fs[0].Title)
	return nil
}

func (s *Service) UpdateAllFeeds() error {
	var fs []*db.FeedSource
	err := s.StoreEngine.Where("active = 1").Find(&fs)
	if err != nil {
		return err
	}
	fmt.Printf("Found %d feeds to update\n", len(fs))
	for _, f := range fs {
		err := s.updateFeedFromURL(f)
		if err != nil {
			return err
		}
		fmt.Printf("Updated db %s\n", f.Title)
	}
	return nil
}

// TODO cleanup
type temp struct {
	db.FeedSource `xorm:"extends"`
	Un int64
}
func (temp) TableName() string {
	return "feed_item"
}

func (s *Service) ListAllFeeds() ([]*temp, error) {
	var fs []*temp
	err := s.StoreEngine.SQL(`select *, (select count(*) from feed_item as y where y.feed_id = s.id and read = 0) as un from feed_source s`).Find(&fs)
	//err := s.StoreEngine.Where("active = 1").Find(&fs)
	if err != nil {
		return nil, err
	}
	return fs, nil
}

func (s *Service) ListAllNews() ([]*db.FeedItem, error) {
	var fs []*db.FeedItem
	err := s.StoreEngine.Find(&fs)
	if err != nil {
		return nil, err
	}
	return fs, nil
}

type filter func(s *xorm.Session) *xorm.Session

func filterUnread() filter {
	return func(s *xorm.Session) *xorm.Session {
		return s.Where("read = 0")
	}
}

func (s *Service) Search(filters ...filter) ([]*db.FeedItem, error) {
	println("searching news")
	var fs []*db.FeedItem
	starting := s.StoreEngine.NewSession()
	for _, f := range filters {
		starting = f(starting)
	}
	err := starting.Find(&fs)
	if err != nil {
		return nil, err
	}
	return fs, nil
}

func (s *Service) ListAllNewsForFeed(feedID int64) ([]*db.FeedItem, error) {
	var fs []*db.FeedItem
	err := s.StoreEngine.Where(fmt.Sprintf("feed_id = %d", feedID)).Find(&fs)
	if err != nil {
		return nil, err
	}
	return fs, nil
}

func (s *Service) MarkAsRead(feedID, newsID int64) error {
	fs := db.FeedItem{
		Read: true,
	}
	if _, err := s.StoreEngine.Id(newsID).Cols("read").Update(&fs); err != nil {
		return err
	}
	return nil
}

func (s *Service) Register(group *mux.Router) {
	group.Methods(http.MethodGet).Path("/feeds").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		things, err := s.ListAllFeeds()
		if err != nil {
			writeErr(w, http.StatusInternalServerError, err)
			return
		}
		json.NewEncoder(w).Encode(things)
	})
	group.Methods(http.MethodPost).Path("/feeds").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := s.UpdateAllFeeds(); err != nil {
			writeErr(w, http.StatusInternalServerError, err)
			return
		}
	})
	group.Methods(http.MethodDelete).Path("/feeds/{id}").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		feedID := mux.Vars(r)["id"]
		n, err := strconv.ParseInt(feedID, 10, 64)
		if err != nil {
			writeErr(w, http.StatusBadRequest, fmt.Errorf("cannot parse feed ID"))
			return
		}
		if err := s.DeleteFeed(n); err != nil {
			writeErr(w, http.StatusInternalServerError, err)
			return
		}
	})
	group.Methods(http.MethodPost).Path("/feeds/{id}").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		feedID := mux.Vars(r)["id"]
		n, err := strconv.ParseInt(feedID, 10, 64)
		if err != nil {
			writeErr(w, http.StatusBadRequest, fmt.Errorf("cannot parse feed ID"))
			return
		}
		if err := s.UpdateFeed(n); err != nil {
			writeErr(w, http.StatusInternalServerError, err)
			return
		}
	})
	group.Methods(http.MethodGet).Path("/feeds/{id}/news").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		feedID := mux.Vars(r)["id"]
		n, err := strconv.ParseInt(feedID, 10, 64)
		if err != nil {
			writeErr(w, http.StatusBadRequest, fmt.Errorf("cannot parse feed ID"))
			return
		}
		things, err := s.ListAllNewsForFeed(n)
		if err != nil {
			writeErr(w, http.StatusInternalServerError, err)
			return
		}
		json.NewEncoder(w).Encode(things)
	})
	group.Methods(http.MethodPost).Path("/feeds/{feed_id}/news/{news_id}").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		newsId, err := strconv.ParseInt(mux.Vars(r)["news_id"], 10, 64)
		feedID, err := strconv.ParseInt(mux.Vars(r)["feed_id"], 10, 64)
		if err != nil {
			writeErr(w, http.StatusBadRequest, err)
			return
		}
		r.ParseForm()
		action := r.Form.Get("action")
		switch action {
		case "read":
			writeErr(w, http.StatusBadRequest, s.MarkAsRead(feedID, newsId))
			return
		default:
			spew.Dump(action, r.Form)
		}
	})
	group.Methods(http.MethodGet).Path("/news").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		things, err := s.Search(filterUnread())
		println("filtere")
		if err != nil {
			writeErr(w, http.StatusInternalServerError, err)
			return
		}
		json.NewEncoder(w).Encode(things)
	})
}

func writeErr(w http.ResponseWriter, statusCode int, err error) {
	if err == nil {
		return
	}
	w.WriteHeader(statusCode)
	w.Write([]byte(err.Error()))
}
