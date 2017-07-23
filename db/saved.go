package db

import "time"

type SavedItem struct {
	FeedItemId int64     `xorm:"pk"`
	NewsItemId int64     `xorm:"pk"`
	Date       time.Time `xorm:"index"`
}
