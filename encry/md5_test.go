package encry

import "testing"

func TestMD5(t *testing.T) {
	t.Log("Test MD5 Start")

	testText := "Hello Golang!"
	resultString := "780c29c240e7c9cf6669eccfa7a321ad"

	{
		md5String := MD5(testText)
		if md5String != resultString {
			t.Errorf("Test MD5 End %v", ballotX)
		} else {
			t.Logf("Test MD5 End %v", checkMark)
		}
	}
}
