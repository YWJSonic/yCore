package mycompress

import (
	"context"

	"github.com/YWJSonic/ycore/driver/compress/zstd"
)

type Handle struct {
	*zstd.Handle
	close context.CancelFunc
}

func New() (*Handle, error) {

	ctx := context.TODO()
	drierHandle, err := zstd.NewStramHandle(ctx)
	if err != nil {
		return nil, err
	}
	_, cancle := context.WithCancel(ctx)

	handle := &Handle{}
	handle.Handle = drierHandle
	handle.close = cancle

	return handle, nil
}

func (self *Handle) Stop() {
	self.close()
}
