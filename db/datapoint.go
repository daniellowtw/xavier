package db

type UserClassification int

const (
	NO_DECISION UserClassification = iota
	POSITIVE
	NEGATIVE
)

// DataPoint for doing our learning
type DataPoint struct {
	// auto public key
	Id         int64
	NewsItemId int64 `xorm:"index"`
	Keywords   []string
	Outcome    UserClassification
}

type ProcessQueue struct {
	// auto public key
	Id         int64
	FeedItemId int64 `xorm:"index"`
	NewsItemId int64 `xorm:"index"`
}
