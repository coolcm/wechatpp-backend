package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

// 计算输入字符串的哈希值
func CalHash(s string) (hash string){
	h := sha256.New()
	h.Write([]byte(s))
	hash = hex.EncodeToString(h.Sum(nil)) //将二进制转化为十六进制表示
	return
}
