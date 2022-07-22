package arango

import (
	"context"
	"fmt"

	"github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
)

func New(addr, username, password, database string) (driver.Database, error) {
	conn, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{addr},
	})
	if err != nil {
		fmt.Printf("[Arango][New] Http new connection error, err: %v\n", err)
		return nil, err
	}

	c, err := driver.NewClient(driver.ClientConfig{
		Connection:     conn,
		Authentication: driver.BasicAuthentication(username, password),
	})
	if err != nil {
		fmt.Printf("[Arango][New] Driver new client error, err: %v\n", err)
		return nil, err
	}

	db, err := c.Database(context.TODO(), database)
	if err != nil {
		fmt.Printf("[Arango][New] Client database error, database: %v, err: %v\n", database, err)
		return nil, err
	}

	fmt.Printf("[Arango][New] Connect success, address: %v, database: %v\n", addr, database)
	return db, nil
}
