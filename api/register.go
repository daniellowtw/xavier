// This file exposes the http api endpoints and marshals the input and calls the respective service
package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"strings"

	"github.com/daniellowtw/xavier/db"
	"github.com/gorilla/mux"
)

func Register(s *Service, group *mux.Router) {
	group.Methods(http.MethodGet).Path("/feeds").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		things, err := s.ListAllFeeds()
		if err != nil {
			writeErr(w, http.StatusInternalServerError, err)
			return
		}
		json.NewEncoder(w).Encode(things)
	})
	group.Methods(http.MethodPost).Path("/feeds").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		total, err := s.UpdateAllFeeds()
		if err != nil {
			writeErr(w, http.StatusInternalServerError, err)
			return
		}
		w.Write([]byte(fmt.Sprintf("updated %d items", total)))
	})
	group.Methods(http.MethodPut).Path("/feeds").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		url := r.Form.Get("url")
		if err := s.AddFeed(url); err != nil {
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
		newsID, err := strconv.ParseInt(mux.Vars(r)["news_id"], 10, 64)
		if err != nil {
			writeErr(w, http.StatusBadRequest, err)
			return
		}
		feedID, err := strconv.ParseInt(mux.Vars(r)["news_id"], 10, 64)
		if err != nil {
			writeErr(w, http.StatusBadRequest, err)
			return
		}
		r.ParseForm()
		action := r.Form.Get("action")
		switch action {
		case "read":
			writeErr(w, http.StatusBadRequest, s.MarkAsRead(newsID))
			return
		case "toggle-save":
			isSaved, err := s.ToggleNews(newsID, feedID)
			if err != nil {
				writeErr(w, http.StatusInternalServerError, err)
				return
			}
			w.Write([]byte(fmt.Sprintf("%v", isSaved)))
			return
		default:
		}
	})
	group.Methods(http.MethodPost).Path("/news").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		limit := 0
		inputLimit := r.Form.Get("limit")
		if inputLimit != "" {
			newLimit, err := strconv.ParseInt(inputLimit, 10, 32)
			if err != nil {
				writeErr(w, http.StatusBadRequest, err)
				return
			}
			limit = int(newLimit)
		}
		searchMode := r.Form.Get("search")
		var searchIds []int64
		feedIDs := r.Form.Get("ids")
		if feedIDs != "" {
			for _, i := range strings.Split(feedIDs, ",") {
				ii, err := strconv.ParseInt(i, 10, 64)
				if err != nil {
					continue
				}
				searchIds = append(searchIds, ii)
			}
		}
		things, err := s.Search(SearchParam{SearchMode: searchMode, Limit: limit, FeedIDs: searchIds})
		if err != nil {
			writeErr(w, http.StatusInternalServerError, err)
			return
		}
		json.NewEncoder(w).Encode(things)
	})
	group.Methods(http.MethodPost).Path("/read").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		newsIDString := r.Form.Get("news_id")
		if err := s.NewsService.MarkAsReadMulti(parseMultiInt(newsIDString)); err != nil {
			writeErr(w, http.StatusBadRequest, fmt.Errorf("cannot mark all as read"))
		}
	})
	group.Methods(http.MethodPost).Path("/learn").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		newsIDString := r.Form.Get("news_id")
		classification := r.Form.Get("classification")
		if newsIDString == "" || classification == "" {
			writeErr(w, http.StatusBadRequest, fmt.Errorf("missing news id or classification"))
			return
		}
		newsId, err := strconv.ParseInt(newsIDString, 10, 64)
		if err != nil {
			writeErr(w, http.StatusBadRequest, err)
			return
		}
		switch classification {
		// TODO These cases are too brittle
		case "1":
			writeErr(w, http.StatusBadRequest, s.HumanClassification(newsId, db.POSITIVE))
		case "2":
			writeErr(w, http.StatusBadRequest, s.HumanClassification(newsId, db.NEGATIVE))
		default:
			writeErr(w, http.StatusBadRequest, fmt.Errorf("can't parse classification"))
		}
		return
	})
}

// returns valid ints
func parseMultiInt(s string) []int64 {
	var res []int64
	for _, i := range strings.Split(s, ",") {
		j, err := strconv.ParseInt(i, 10, 64)
		if err != nil {
			continue
		}
		res = append(res, j)
	}
	return res
}
