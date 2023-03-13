package stringx

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"

	"github.com/spaolacci/murmur3"
)

func Hash(data []byte) uint64 {
	return murmur3.Sum64(data)
}

// MD5加密
func MD5(s string) string {
	hash := md5.New()
	hash.Write([]byte(s))
	return hex.EncodeToString(hash.Sum(nil))
}

// sha1加密
func SHA1(s string) string {
	hash := sha1.New()
	hash.Write([]byte(s))
	return hex.EncodeToString(hash.Sum(nil))
}

// 密码加盐
func PasswordSalt(password, salt string) string {
	hash := hmac.New(sha1.New, []byte(salt))
	hash.Write([]byte(password))
	return hex.EncodeToString(hash.Sum(nil))
}
