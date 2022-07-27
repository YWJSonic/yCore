package example

import (
	"context"
	"log"
	"sync"
	"testing"
	"ycore/driver/compress/zstd"
	"ycore/module/mylog"
	"ycore/util"
)

func Test_To(t *testing.T) {
	var dataStr = `{"1":"%v"}`
	handle := zstd.NewStramHandle(context.TODO())
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
