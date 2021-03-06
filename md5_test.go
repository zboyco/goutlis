package goutlis_test

import (
	"testing"

	"github.com/zboyco/goutlis"
)

func TestMD5(t *testing.T) {
	t.Log("Test MD5 Start")

	testText := "Hello Golang!"
	resultString := "780c29c240e7c9cf6669eccfa7a321ad"

	{
		md5String := goutlis.MD5(testText)
		if md5String != resultString {
			t.Errorf("Test MD5 End %v", ballotX)
		} else {
			t.Logf("Test MD5 End %v", checkMark)
		}
	}
}
