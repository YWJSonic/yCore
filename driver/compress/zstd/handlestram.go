package zstd

import (
	"context"

	"github.com/klauspost/compress/zstd"
)

// ------------------ 串流解碼 關閉前使用同一份字典 ------------------

type Handle struct {
	encoder *zstd.Encoder
	decoder *zstd.Decoder
	ctx     context.Context
}

func NewStramHandle(ctx context.Context) *Handle {

	var encoder, _ = zstd.NewWriter(nil, zstd.WithWindowSize(WithWindowSize))
	var decoder, _ = zstd.NewReader(nil, zstd.WithDecoderMaxWindow(WithDecoderMaxWindowSize), zstd.WithDecoderMaxMemory(WithDecoderMaxMemorySize))

	hanle := &Handle{
		encoder: encoder,
		decoder: decoder,
		ctx:     ctx,
	}
	go hanle.listen()
	return hanle
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
