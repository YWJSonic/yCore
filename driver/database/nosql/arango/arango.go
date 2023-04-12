package arango

import (
	"context"

	"github.com/YWJSonic/ycore/module/mylog"

	"github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
)

func New(addr, username, password, database string) (driver.Database, error) {
	conn, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{addr},
	})
	if err != nil {
		mylog.Errorf("[Arango][New] Http new connection error, err: %v", err)
		return nil, err
	}

	c, err := driver.NewClient(driver.ClientConfig{
		Connection:     conn,
		Authentication: driver.BasicAuthentication(username, password),
	})
	if err != nil {
		mylog.Errorf("[Arango][New] Driver new client error, err: %v", err)
		return nil, err
	}

	db, err := c.Database(context.TODO(), database)
	if err != nil {
		mylog.Errorf("[Arango][New] Client database error, database: %v, err: %v", database, err)
		return nil, err
	}

	mylog.Infof("[Arango][New] Connect success, address: %v, database: %v", addr, database)
	return db, nil
}
