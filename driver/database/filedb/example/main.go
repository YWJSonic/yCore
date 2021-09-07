package example

import (
	"fmt"
	"yangServer/service/database/filedb/localDBDriver"
)

type LocalDB struct {
	driver *localDBDriver.Driver
}

func Connect(setting struct{ Path string }) (db *LocalDB, err error) {
	return &LocalDB{
		localDBDriver.NewDriver(setting),
	}, nil
}

func formatKey(datas ...interface{}) string {
	var key string
	for _, data := range datas {
		key += fmt.Sprintf("%v_", data)
	}
	if len(key) > 0 {
		key = key[:len(key)-1]
	}
	return key
}
