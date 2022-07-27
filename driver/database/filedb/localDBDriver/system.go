package localDBDriver

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"ycore/driver/database/filedb/constant"
	"ycore/util"
)

func (self *Driver) NewCollection(collection string) error {
	path := util.Sprintf("%v%v", self.path, collection)
	_, err := os.Open(path)
	if !os.IsNotExist(err) {
		return errors.New("Collection Already Exist")
	}

	err = os.Mkdir(path, 0666)
	if err != nil {
		return err
	}

	head := newFileHead()
	err = self.setHead(collection, head)
	if err != nil {
		return err
	}
	return nil
}

func (self *Driver) Set(collection string, key string, data interface{}) (interface{}, error) {
	path := util.Sprintf("%v/%v/%v", self.path, collection, key)
	result := &insertResult{}
	byteData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	if key == "" {
		head, err := self.getHead(path)
		if err != nil {
			return nil, err
		}

		head.Incr++
		key = strconv.FormatInt(head.Incr, 10)
		path = util.Sprintf("%v%v/%v", self.path, collection, key)
		_ = self.setHead(collection, head)
	}

	err = ioutil.WriteFile(path, byteData, 0666)
	if err != nil {
		return nil, err
	}
	result.Key = key

	return result, nil
}

func (self *Driver) GetLike(collection, key string) (map[string]interface{}, error) {
	datas := make(map[string]interface{})

	path := util.Sprintf("%v/%v", self.path, collection)
	fileInfos, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, errors.New(constant.NoData)
	}

	for _, fileInfo := range fileInfos {
		if ok := strings.Index(fileInfo.Name(), key); ok >= 0 {
			file, err := os.Open(util.Sprintf("%v/%v", path, fileInfo.Name()))
			if err != nil {
				return nil, err
			}
			defer file.Close()
			byteData, err := ioutil.ReadAll(file)
			if err != nil {
				return nil, err
			}

			var data interface{}
			err = json.Unmarshal(byteData, &data)
			if err != nil {
				return nil, err
			}

			datas[fileInfo.Name()] = data
		}
	}
	return datas, nil
}

func (self *Driver) Get(collection, key string) (interface{}, error) {
	var data interface{}

	path := util.Sprintf("%v%v/%v", self.path, collection, key)
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	byteData, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(byteData, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (self *Driver) setHead(collection string, head *fileHead) error {
	data, err := json.Marshal(head)
	if err != nil {
		return err
	}

	path := util.Sprintf("%v%v/%v", self.path, collection, "head")
	file, err := os.OpenFile(path, os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	err = ioutil.WriteFile(path, data, 0666)
	if err != nil {
		return err
	}

	return nil
}

func (self *Driver) getHead(collection string) (*fileHead, error) {

	path := util.Sprintf("%v%v/%v", self.path, collection, "head")
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	headData, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	head := &fileHead{}
	err = json.Unmarshal(headData, head)
	if err != nil {
		return nil, err
	}

	return head, nil
}
