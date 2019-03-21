package encry_test

import (
	"testing"

	"github.com/zboyco/goutlis/encry"
)

func TestHashPassword(t *testing.T) {
	t.Log("Test HashPassword Start")

	testText := "password"
	{
		hashString, err := encry.HashPassword(testText)
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
		ok := encry.VerifyPassword(testText, hashString)
		if !ok {
			t.Errorf("Test Verify Password End %v", ballotX)
		} else {
			t.Logf("Test Verify Password End %v", checkMark)
		}
	}
}
