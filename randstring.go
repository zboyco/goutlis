package goutlis

import (
	"math/rand"
	"time"
)

const randStringBase = "0123456789abcdefghijklmnopqrstuvwxyz"

// RandString 生成随机字符串
func RandString(lenght int64) string {
	rand.Seed(time.Now().UnixNano())
	result := make([]byte, lenght)
	baseLenght := len(randStringBase)
	for i := range result {
		result[i] = randStringBase[rand.Int()%baseLenght]
	}
	return string(result)
}
