package gzip

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
)

func CompressByte(in []byte, level int) ([]byte, error) {
	// 創建一個緩衝區來存儲壓縮後的數據
	var compressedData bytes.Buffer

	// 創建一個 Gzip 寫入器
	gzipWriter, _ := gzip.NewWriterLevel(&compressedData, level)
	// NoCompression      = flate.NoCompression
	// BestSpeed          = flate.BestSpeed
	// BestCompression    = flate.BestCompression
	// DefaultCompression = flate.DefaultCompression
	// HuffmanOnly        = flate.HuffmanOnly

	// 將文本數據寫入 Gzip 寫入器
	_, err := gzipWriter.Write(in)
	if err != nil {
		fmt.Println("寫入數據到 Gzip 寫入器時出錯：", err)
		return nil, err
	}

	// 關閉 Gzip 寫入器（重要！）
	err = gzipWriter.Close()
	if err != nil {
		fmt.Println("關閉 Gzip 寫入器時出錯：", err)
		return nil, err
	}

	return compressedData.Bytes(), nil
}

func DecompressByte(in []byte) ([]byte, error) {
	// 創建一個緩衝區來存儲壓縮後的數據
	decode := bytes.NewBuffer(in)

	// 要解壓縮數據，您可以使用 Gzip 讀取器
	// 創建一個 Gzip 讀取器
	gzipReader, err := gzip.NewReader(decode)
	if err != nil {
		return nil, err
	}
	defer gzipReader.Close()

	// 讀取解壓縮後的數據
	decompressedData, err := io.ReadAll(gzipReader)
	if err != nil {
		return nil, err
	}

	return decompressedData, nil
}
