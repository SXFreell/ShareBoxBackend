package utils

import (
	"crypto/rand"
)

// GenerateCode 生成随机Code
func GenerateCode() string {
	// 获取26个字母随机数
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	number := "0123456789"
	var bytes = make([]byte, 6)
	rand.Read(bytes)
	code := ""
	for i, b := range bytes {
		if i == 0 {
			code += string(alphabet[b%26])
		} else {
			code += string(number[b%10])
		}
	}
	return code
}
