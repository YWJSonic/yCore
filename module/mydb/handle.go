package mydb

import (
	"context"
)

// Quary Create
func (self *Manager) Create(ctx context.Context, collection string, doc interface{}) error {
	coll, err := self.Client.Collection(ctx, collection)
	if err != nil {
		return err
	}

	_, err = coll.CreateDocument(ctx, doc)
	return err
}

func (self *Manager) CreateAndResKey(ctx context.Context, collection string, doc interface{}) (string, error) {
	coll, err := self.Client.Collection(ctx, collection)
	if err != nil {
		return "", err
	}

	meta, err := coll.CreateDocument(ctx, doc)
	return meta.Key, err
}

// Quary Read
func (self *Manager) Read(ctx context.Context, collection string, key string, outDoc interface{}) error {
	coll, err := self.Client.Collection(ctx, collection)
	if err != nil {
		return err
	}
	_, err = coll.ReadDocument(ctx, key, outDoc)
	return err
}

// Quary Update
func (self *Manager) Update(ctx context.Context, collection string, key string, doc interface{}) error {
	cursor, err := self.Client.Query(
		ctx,
		`UPDATE @key WITH @doc IN @@collection`,
		map[string]interface{}{
			"key":         key,
			"doc":         doc,
			"@collection": collection,
		})
	cursor.Close()
	return err
}

// Quary Update
func (self *Manager) Delete(ctx context.Context, collection string, key string) error {
	cursor, err := self.Client.Query(
		ctx,
		`REMOVE @key IN @@collection`,
		map[string]interface{}{
			"key":         key,
			"@collection": collection,
		})
	cursor.Close()
	return err
}
