package localDBDriver

import (
	"fmt"
	"testing"
)

func TestLocal(t *testing.T) {
	db := NewDriver(struct{ Path string }{Path: "./"})
	db.NewCollection("testCollection")

	head := newFileHead()
	head.Incr = 999
	result, err := db.Set("testCollection", "", head)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	fmt.Println(result)

	result, err = db.Get("testCollection", result.(*insertResult).Key)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	fmt.Println(result)

	data := []map[string]interface{}{}
	data = append(data, map[string]interface{}{"1edqww": 123})
	result, err = db.Set("testCollection", "", data)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	fmt.Println(result)

	result, err = db.Get("testCollection", result.(*insertResult).Key)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	fmt.Println(result)
}
