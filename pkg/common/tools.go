package common

import (
	"crypto/md5"
	"encoding/hex"
)

func GetMD5WithSalt(s string) string {
	salt := "password salt with sunny weather"
	h := md5.New()
	h.Write([]byte(s + salt))
	return hex.EncodeToString(h.Sum(nil))
}
