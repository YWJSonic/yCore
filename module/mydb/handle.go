package mydb

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/YWJSonic/ycore/util"
	"github.com/arangodb/go-driver"
)

// Quary String
func (self *Manager) Quary(ctx context.Context, query string, bindVars map[string]interface{}, outDoc interface{}) error {
	cursor, err := self.Client.Query(ctx, query, bindVars)
	if err != nil {
		return err
	}
	defer cursor.Close()

	if !cursor.HasMore() {
		return nil
	}

	list := []interface{}{}
	for {
		doc := make(map[string]interface{})
		_, err = cursor.ReadDocument(ctx, &doc)
		if driver.IsNoMoreDocuments(err) {
			break
		}
		list = append(list, doc)
	}

	// 避免driver撈出科學記號數字變成浮點數，過一層轉換將金額欄位維持整數格式
	bytes, err := util.Marshal(list)
	if err != nil {
		return err
	}
	err = util.Unmarshal(bytes, outDoc)
	if err != nil {
		return err
	}

	return err
}

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

// Quary Reads
func (self *Manager) Reads(ctx context.Context, collection string, keys []string, outDoc interface{}) error {
	coll, err := self.Client.Collection(ctx, collection)
	if err != nil {
		return err
	}
	_, sliceErr, err := coll.ReadDocuments(ctx, keys, outDoc)
	fmt.Println(sliceErr, err)
	return err
}

// Quary Update
func (self *Manager) Update(ctx context.Context, collection string, key string, doc interface{}) error {
	coll, err := self.Client.Collection(ctx, collection)
	if err != nil {
		return err
	}
	_, err = coll.UpdateDocument(ctx, key, doc)
	return err
}

// Quary Delete
func (self *Manager) Delete(ctx context.Context, collection string, key string) error {
	coll, err := self.Client.Collection(ctx, collection)
	if err != nil {
		return err
	}
	_, err = coll.RemoveDocument(ctx, key)
	return err
}

// Transaction
func (self *Manager) Transaction(ctx context.Context, action string, opt *driver.TransactionOptions, outDoc interface{}) error {
	res, err := self.Client.Transaction(ctx, action, opt)
	switch reflect.TypeOf(outDoc).Kind() {
	case reflect.Ptr, reflect.Struct:
		jsByte, _ := json.Marshal(res)
		err := json.Unmarshal(jsByte, outDoc)
		if err != nil {
			return err
		}

	}
	return err
}
