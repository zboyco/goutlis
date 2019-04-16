package goutlis_test

import (
	"testing"
	"time"

	"github.com/zboyco/goutlis"
)

func Test_RandString(t *testing.T) {
	t.Log(time.Now())
	for i := 0; i < 100; i++ {
		randString := goutlis.RandString(8)
		t.Log(randString)
	}
}
