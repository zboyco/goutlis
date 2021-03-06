package goutlis_test

import (
	"encoding/base64"
	"testing"

	"github.com/zboyco/goutlis"
)

func TestAesEncrypt(t *testing.T) {
	t.Log("Test Aes Encrypt Start")

	testText := "Hello Golang!"
	key := "1234567812345678"
	resultString := "a7XvRl1kgWSHX6FgXw7UhA=="

	t.Logf("Text : \"%v\" , key : \"%v\"", testText, key)
	{
		bytes, err := goutlis.AesEncrypt([]byte(testText), key)
		if err != nil {
			t.Fatal(err)
		}
		bytesString := base64.StdEncoding.EncodeToString(bytes)
		if bytesString != resultString {
			t.Errorf("Test Aes Encrypt End %v", ballotX)
		} else {
			t.Logf("Test Aes Encrypt End %v", checkMark)
		}
	}
}
