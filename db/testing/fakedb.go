package testing

import (
	"os"

	"github.com/daniellowtw/xavier/db"
	"xorm.io/core"
)

var testFileName = "./for-testing.db"

func NewFakeDB() (*db.Client, error) {
	ll := core.LogLevel(core.LOG_OFF)
	return db.NewSqlite3Client(testFileName, false, ll)
}

func CleanUp() error {
	return os.Remove(testFileName)
}
