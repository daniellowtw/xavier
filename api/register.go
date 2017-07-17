package api

import (
	"github.com/gorilla/mux"
	"strconv"
	"fmt"
	"encoding/json"
	"github.com/davecgh/go-spew/spew"
	"net/http"
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
			spew.Dump(action, r.Form)
		}
	})
	group.Methods(http.MethodGet).Path("/news").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		things, err := s.Search(SearchParam{IncludeRead: false})
		println("filtere")
		if err != nil {
			writeErr(w, http.StatusInternalServerError, err)
			return
		}
		json.NewEncoder(w).Encode(things)
	})
}

