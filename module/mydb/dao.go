package mydb

import (
	"context"

	"github.com/arangodb/go-driver"
)

type Manager struct {
	Client driver.Database
}

// Quary Create
func (self *Manager) Create(ctx context.Context, collection string, doc interface{}) error {
	_, err := self.Client.Query(
		ctx,
		"Insert @doc Into @@collection",
		map[string]interface{}{
			"doc":         doc,
			"@collection": collection,
		})
	return err
}

// Quary Read
func (self *Manager) Read(ctx context.Context, collection string, key string, outDoc interface{}) error {
	cursor, err := self.Client.Query(
		ctx,
		`FOR doc IN @@collection
		FILTER doc._key == @key
		RETURN doc`,
		map[string]interface{}{
			"key":         key,
			"@collection": collection,
		})
	if err != nil {
		return err
	}
	defer cursor.Close()

	if outDoc != nil && cursor.HasMore() {
		_, err := cursor.ReadDocument(ctx, outDoc)
		return err
	} else {
		return nil
	}
}

// Quary Update
func (self *Manager) Update(ctx context.Context, collection string, key string, doc interface{}) error {
	_, err := self.Client.Query(
		ctx,
		`UPDATE @key WITH @doc IN @@collection`,
		map[string]interface{}{
			"key":         key,
			"doc":         doc,
			"@collection": collection,
		})
	return err
}

// Quary Update
func (self *Manager) Delete(ctx context.Context, collection string, key string) error {
	_, err := self.Client.Query(
		ctx,
		`REMOVE @key IN @@collection`,
		map[string]interface{}{
			"key":         key,
			"@collection": collection,
		})
	return err
}
