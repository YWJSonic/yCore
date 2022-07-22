package mydb

import "ycore/driver/database/nosql/arango"

func NewArangoDB(addr, username, password, database string) (*Manager, error) {
	obj := &Manager{}
	db, err := arango.New(addr, username, password, database)
	if err != nil {
		return nil, err
	}
	obj.Client = db
	return obj, nil
}
