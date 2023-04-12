package example

import (
	"context"
	"log"
	"sync"
	"testing"

	"github.com/YWJSonic/ycore/driver/compress/zstd"
	"github.com/YWJSonic/ycore/module/mylog"
	"github.com/YWJSonic/ycore/util"
)

func Test_To(t *testing.T) {
	var dataStr = `{"1":"%v"}`
	handle, _ := zstd.NewStramHandle(context.TODO())
	wg := sync.WaitGroup{}
	wg.Add(20)
	for i := 0; i < 20; i++ {

		go func(i int) {

			out := handle.Compress([]byte(util.Sprintf(dataStr, i)))

			get, err := handle.Decompress(out)
			if err != nil {
				log.Fatalln(err)
			}

			mylog.Info(string(get))
			wg.Done()
		}(i)
	}
	wg.Wait()
}
