package utils

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

// md5验证
func Md5(src string) string {
	h := md5.New()
	h.Write([]byte(src)) // 需要加密的字符串为
	fmt.Printf("%s\n", hex.EncodeToString(h.Sum(nil)))
	return hex.EncodeToString(h.Sum(nil))
}

// HmacSha256
func HmacSha256 (src, key string) string {
	m := hmac.New(sha256.New, []byte(key))
	m.Write([]byte(src))
	return hex.EncodeToString(m.Sum(nil))
}

// sha256验证
func SHA256Str(src string) string {
	h := sha256.New()
	h.Write([]byte(src)) // 需要加密的字符串为
	// fmt.Printf("%s\n", hex.EncodeToString(h.Sum(nil))) // 输出加密结果
	return hex.EncodeToString(h.Sum(nil))
}

// base编码
func BASE64EncodeStr(src string) string {
	return string(base64.StdEncoding.EncodeToString([]byte(src)))
}

// base解码
func BASE64DecodeStr(src string) string {
	a, err := base64.StdEncoding.DecodeString(src)
	if err != nil {
		return ""
	}
	return string(a)
}
