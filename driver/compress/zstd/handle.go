package zstd

import (
	"bytes"
	"io"

	"github.com/klauspost/compress/zstd"
)

// ------------------ 單次編碼解碼 字典不共用 ------------------

func CompressString(in string) ([]byte, error) {
	bufIn := bytes.NewBufferString(in)
	bufOut := &bytes.Buffer{}

	err := Compress(bufIn, bufOut)
	return bufOut.Bytes(), err
}

func CompressByte(in []byte) ([]byte, error) {
	bufIn := bytes.NewBuffer(in)
	bufOut := &bytes.Buffer{}

	err := Compress(bufIn, bufOut)
	return bufOut.Bytes(), err
}

func DecompressString(in string) ([]byte, error) {
	bufIn := bytes.NewBufferString(in)
	bufOut := &bytes.Buffer{}

	err := Decompress(bufIn, bufOut)
	return bufOut.Bytes(), err
}

func DecompressByte(in []byte) ([]byte, error) {
	bufIn := bytes.NewBuffer(in)
	bufOut := &bytes.Buffer{}

	err := Decompress(bufIn, bufOut)
	return bufOut.Bytes(), err
}

func Compress(in io.Reader, out io.Writer) error {
	enc, err := zstd.NewWriter(out)
	if err != nil {
		return err
	}
	_, err = io.Copy(enc, in)
	if err != nil {
		enc.Close()
		return err
	}
	return enc.Close()
}

func Decompress(in io.Reader, out io.Writer) error {
	d, err := zstd.NewReader(in)
	if err != nil {
		return err
	}
	defer d.Close()

	// Copy content...
	_, err = io.Copy(out, d)
	return err
}
