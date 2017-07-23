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
		newsId, err := strconv.ParseInt(mux.Vars(r)["news_id"], 10, 64)
		if err != nil {
			writeErr(w, http.StatusBadRequest, err)
			return
		}
		r.ParseForm()
		action := r.Form.Get("action")
		switch action {
		case "read":
			writeErr(w, http.StatusBadRequest, s.MarkAsRead(newsId))
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
		includeRead := true
		if searchMode == "unread" {
			includeRead = false
		}
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
		things, err := s.Search(SearchParam{IncludeRead: includeRead, Limit: limit, FeedIDs: searchIds})
		if err != nil {
			writeErr(w, http.StatusInternalServerError, err)
			return
		}
		json.NewEncoder(w).Encode(things)
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
		case "1":
			s.HumanClassification(newsId, db.POSITIVE)
		case "-1":
			s.HumanClassification(newsId, db.NEGATIVE)
		default:
			writeErr(w, http.StatusBadRequest, fmt.Errorf("can't prase classification"))
		}
		return
	})
}
