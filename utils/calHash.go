package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

func CalHash(s string) (hash string){
	h := sha256.New()
	h.Write([]byte(s))
	hash = hex.EncodeToString(h.Sum(nil)) //计算hash值
	return
}
