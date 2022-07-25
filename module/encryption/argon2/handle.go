package argon2

import (
	"crypto/rand"

	"golang.org/x/crypto/argon2"
)

// 推薦的加密選項
// 優點: 安全性高, 沒有不安全的係數
//
// 缺點: 處理時間較長

type Params struct {
	memory      uint32
	iterations  uint32
	parallelism uint8
	saltLength  uint32
	keyLength   uint32
}

// 預設加密參數
var defaultParams = &Params{
	memory:      64 * 1024, // 使用記憶體大小需要以 KiB 為單為增減
	iterations:  5,         // 雜湊次數(過少的話有機率被暴力破解)
	parallelism: 2,         // 執行併發數量(併發處理幣面)
	saltLength:  16,        // 雜訊長度
	keyLength:   32,        // 最後雜湊數長度
}

func Encode(data []byte, params *Params) ([]byte, error) {
	if params == nil {
		params = defaultParams
	}

	// Generate a cryptographically secure random salt.
	salt, err := generateRandomBytes(params.saltLength)
	if err != nil {
		return nil, err
	}

	// Pass the plaintext password, salt and parameters to the argon2.IDKey
	// function. This will generate a hash of the password using the Argon2id
	// variant.
	hash := argon2.IDKey(data, salt, params.iterations, params.memory, params.parallelism, params.keyLength)
	return hash, nil

}

func generateRandomBytes(n uint32) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}
