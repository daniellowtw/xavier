package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/daniellowtw/xavier/db"
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
	group.Methods(http.MethodGet).Path("/news").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		things, err := s.Search(SearchParam{IncludeRead: false})
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
