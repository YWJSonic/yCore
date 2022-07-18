package arango

import (
	"context"
	"fmt"

	"github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
)

func New(addr, username, password, database string) (*Manager, error) {

	obj := &Manager{}

	conn, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{addr},
	})
	if err != nil {
		fmt.Printf("[Arango][New] Http new connection error, err: %v", err)
		return nil, err
	}

	c, err := driver.NewClient(driver.ClientConfig{
		Connection:     conn,
		Authentication: driver.BasicAuthentication(username, password),
	})
	if err != nil {
		fmt.Printf("[Arango][New] Driver new client error, err: %v", err)
		return nil, err
	}

	db, err := c.Database(context.TODO(), database)
	if err != nil {
		fmt.Printf("[Arango][New] Client database error, database: %v, err: %v", database, err)
		return nil, err
	}

	obj.Client = db
	fmt.Printf("[Arango][New] Connect success, address: %v, database: %v", addr, database)

	return obj, nil
}

type Manager struct {
	Client driver.Database
}
