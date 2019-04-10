package goutlis_test

import (
	"testing"

	"github.com/zboyco/goutlis"
)

func TestHashPassword(t *testing.T) {
	t.Log("Test HashPassword Start")

	testText := "password"
	{
		hashString, err := goutlis.HashPassword(testText)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("Test HashPassword End , MD5 : %v %v", hashString, checkMark)
	}
}

func TestVerifyPassword(t *testing.T) {
	t.Log("Test Verify Password Start")

	testText := "password"

	hashString := "$2a$10$oIWfMhLEDjmn7QcMHApOGOahyg4gIJ1zmsp.WNHdxypL092IrBMyq"
	{
		ok := goutlis.VerifyPassword(testText, hashString)
		if !ok {
			t.Errorf("Test Verify Password End %v", ballotX)
		} else {
			t.Logf("Test Verify Password End %v", checkMark)
		}
	}
}
