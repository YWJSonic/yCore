package zstd

import (
	"context"

	"github.com/klauspost/compress/zstd"
)

// ------------------ 串流解碼 與單次的區別？------------------

type Handle struct {
	encoder *zstd.Encoder
	decoder *zstd.Decoder
	ctx     context.Context
}

func NewStramHandle(ctx context.Context) (*Handle, error) {

	encoder, err := zstd.NewWriter(nil, zstd.WithWindowSize(WithWindowSize))
	if err != nil {
		return nil, err
	}

	decoder, err1 := zstd.NewReader(nil, zstd.WithDecoderMaxWindow(WithDecoderMaxWindowSize), zstd.WithDecoderMaxMemory(WithDecoderMaxMemorySize))
	if err1 != nil {
		return nil, err1
	}

	hanle := &Handle{
		encoder: encoder,
		decoder: decoder,
		ctx:     ctx,
	}
	go hanle.listen()
	return hanle, nil
}

func (self *Handle) Compress(src []byte) []byte {
	return self.encoder.EncodeAll(src, nil)
}

func (self *Handle) Decompress(src []byte) ([]byte, error) {
	return self.decoder.DecodeAll(src, nil)
}

func (self *Handle) listen() {
	<-self.ctx.Done()
	self.decoder.Close()
	self.encoder.Close()
}
