package example

import (
	"compress/flate"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/YWJSonic/ycore/driver/compress/gzip"
)

// BestSpeed
// 原始數據大小: 1564206
// 壓縮後的數據大小: 68183
// cpu: 13th Gen Intel(R) Core(TM) i9-13900K
// Benchmark_Gzip-32    	     646	   1829857 ns/op	 1471518 B/op	      31 allocs/op

// BestCompression
// 原始數據大小: 1564206
// 壓縮後的數據大小: 29569
// cpu: 13th Gen Intel(R) Core(TM) i9-13900K
// Benchmark_Gzip-32    	     160	   7279419 ns/op	  889044 B/op	      27 allocs/op

// DefaultCompression
// 原始數據大小: 1564206
// 壓縮後的數據大小: 32608
// cpu: 13th Gen Intel(R) Core(TM) i9-13900K
// Benchmark_Gzip-32    	     258	   4467703 ns/op	  885329 B/op	      27 allocs/op

// HuffmanOnly
// 原始數據大小: 1564206
// 壓縮後的數據大小: 971602

func Test_gzip(t *testing.T) {

	encode, err := gzip.CompressByte([]byte(text), 2)
	if err != nil {
		log.Fatalln(err)
	}

	// 壓縮後的數據現在位於 'compressedData' 緩衝區中
	fmt.Println("原始數據大小:", len([]byte(text)))
	fmt.Println("壓縮後的數據大小:", len(encode))

	// 將壓縮後的數據保存到 Gzip 文件
	err = os.WriteFile("compressed_data.gz", encode, os.ModePerm)
	if err != nil {
		log.Fatalf("保存 Gzip 文件時出錯：%v", err)
	}

	// 要解壓縮數據，您可以使用 Gzip 讀取器
	// 創建一個 Gzip 讀取器
	_, err = gzip.DecompressByte(encode)
	if err != nil {
		log.Fatalf("解壓縮失敗: %v", err)
	}
}

func Benchmark_Gzip(b *testing.B) {
	mock := []byte(text)
	// NoCompression      = 0
	// BestSpeed          = 1
	// BestCompression    = 9
	// DefaultCompression = -1

	for i := 0; i < b.N; i++ {
		_, err := gzip.CompressByte(mock, flate.DefaultCompression)
		if err != nil {
			log.Fatalln(err)
		}
	}
}
