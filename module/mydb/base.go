package mydb

import (
	"errors"

	"github.com/YWJSonic/ycore/driver/database/nosql/arango"
	"github.com/arangodb/go-driver"
)

var NoDataError = errors.New("data not find")

type Manager struct {
	Client driver.Database
}

func NewArangoDB(addr, username, password, database string) (*Manager, error) {
	obj := &Manager{}
	db, err := arango.New(addr, username, password, database)
	if err != nil {
		return nil, err
	}
	obj.Client = db
	return obj, nil
}
