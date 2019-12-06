package encrypt

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
)

func Hmac(key []byte, data []byte) string {
	obj := hmac.New(md5.New, key)
	obj.Write(data)
	return hex.EncodeToString(obj.Sum([]byte(nil)))
}

//result length 32
func Md5(data []byte) string {
	obj := md5.New()
	obj.Write(data)
	return hex.EncodeToString(obj.Sum(nil))
}

//result length 40
func Sha1(data []byte) string {
	obj := sha1.New()
	obj.Write(data)
	return hex.EncodeToString(obj.Sum(nil))
}
