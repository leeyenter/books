package db_test

import (
	"github.com/leeyenter/books/backend/db"
	"testing"
)

func TestCreateDB(t *testing.T) {
	dbObj := db.Get()
	dbObj.Close()
}
