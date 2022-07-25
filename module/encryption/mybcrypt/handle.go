package mybcrypt

import (
	"golang.org/x/crypto/bcrypt"
)

// 普遍使用的加密方式
// 優點: 速度快
//
// 缺點: 存在不安全參數(不知道哪不安全了)

func Encode(data []byte, cost int) ([]byte, error) {
	return bcrypt.GenerateFromPassword(data, cost)
}

func Compare(target, hash []byte) bool {
	if err := bcrypt.CompareHashAndPassword(hash, target); err != nil {
		return false
	}
	return true
}
