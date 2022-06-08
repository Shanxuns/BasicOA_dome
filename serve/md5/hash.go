package md5

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"time"
)

// GetMD5Encode 加密
func GetMD5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

// Hash 加盐加密
func Hash(str string) string {
	return GetMD5Encode(GetMD5Encode(str) + GetMD5Encode(hash))
}

func SetHash(_hash string) {
	hash = _hash
}

// GenerateSubId 随机字符串
func GenerateSubId() string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.New(rand.NewSource(time.Now().UnixNano())).Intn(len(letterRunes))]
	}
	return GetMD5Encode(string(b))
}

// File 文件MD5
func File(file []byte) string {
	return hex.EncodeToString(file)
}
