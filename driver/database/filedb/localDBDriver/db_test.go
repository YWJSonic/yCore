package localDBDriver

import (
	"testing"
	"ycore/module/mylog"
)

func TestLocal(t *testing.T) {
	db := NewDriver(struct{ Path string }{Path: "./"})
	_ = db.NewCollection("testCollection")

	head := newFileHead()
	head.Incr = 999
	result, err := db.Set("testCollection", "", head)
	if err != nil {
		mylog.Error(err)
		t.FailNow()
	}
	mylog.Info(result)

	result, err = db.Get("testCollection", result.(*insertResult).Key)
	if err != nil {
		mylog.Error(err)
		t.FailNow()
	}
	mylog.Info(result)

	data := []map[string]interface{}{}
	data = append(data, map[string]interface{}{"1edqww": 123})
	result, err = db.Set("testCollection", "", data)
	if err != nil {
		mylog.Error(err)
		t.FailNow()
	}
	mylog.Info(result)

	result, err = db.Get("testCollection", result.(*insertResult).Key)
	if err != nil {
		mylog.Error(err)
		t.FailNow()
	}
	mylog.Info(result)
}
