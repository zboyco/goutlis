package goutlis

import (
	"crypto/md5"
	"encoding/hex"
)

// MD5 MD5
func MD5(text string) string {
	h := md5.New()
	h.Write([]byte(text))
	return hex.EncodeToString(h.Sum(nil))
}
