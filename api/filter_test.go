package api

import (
	"testing"

	"github.com/daniellowtw/xavier/db"
	testing2 "github.com/daniellowtw/xavier/db/testing"
	"github.com/stretchr/testify/require"
)

func TestFilterService_AddRegexFilter(t *testing.T) {
	client, err := testing2.NewFakeDB()
	defer testing2.CleanUp()
	require.NoError(t, err)
	testLine := "abc"
	client.AddFeed(&db.FeedSource{
		Id: 1,
	})
	client.AddNews(1, &db.FeedItem{
		Id:     1000,
		FeedId: 1,
		Title:  testLine,
	})
	client.AddNews(1, &db.FeedItem{
		Id:          1001,
		FeedId:      1,
		Description: testLine,
	})
	client.AddNews(1, &db.FeedItem{
		Id:      1002,
		FeedId:  1,
		Content: testLine,
	})
	client.AddNews(1, &db.FeedItem{
		Id:      1003,
		FeedId:  1,
		Content: "non match",
	})

	s := filterService{
		dbClient: client,
		flagItemClient: &db.FlagItemClient{
			Engine: client.Engine,
		},
	}
	err = s.AddRegexFilter(1, testLine, db.Flag)
	require.NoError(t, err)
	err = s.RunFilter()
	require.NoError(t, err)
	ids, err := s.flagItemClient.GetAllFlagged(1)
	require.NoError(t, err)
	require.Equal(t, 3, len(ids), "regex should pass for title, content and description")
}
