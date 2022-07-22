package mydb

import (
	"context"

	"github.com/arangodb/go-driver"
)

type Manager struct {
	Client driver.Database
}

// Quary Insert
func (self *Manager) Insert(ctx context.Context, collection string, doc interface{}) error {
	_, err := self.Client.Query(
		ctx,
		"Insert @doc Into @@collection",
		map[string]interface{}{
			"doc":         doc,
			"@collection": collection,
		})
	return err
}
