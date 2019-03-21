package redis_test

import (
	"strconv"
	"testing"

	"github.com/zboyco/goutlis/redis"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

func TestRedis(t *testing.T) {
	t.Log("Test Redis Start")

	conf := &redis.RedisConfig{
		Address:     "192.168.2.99:6379",
		IdleTimeout: 30,
		MaxActive:   3,
		MaxIdle:     3,
	}
	redis.InitRedis(conf)

	{
		t.Log("Test Redis Set Start")
		result, err := redis.Set("Test", "Test", 10)
		if err != nil {
			t.Fatalf("Test Redis Set End. %v %v ", err, ballotX)
		}
		if result != "OK" {
			t.Errorf("Test Redis Set End. %v %v ", result, ballotX)
		} else {
			t.Logf("Test Redis Set End. %v %v ", result, checkMark)
		}
	}

	{
		t.Log("Test Redis Expire Start")
		affect, err := redis.Expire("Test", 10)
		if err != nil {
			t.Fatalf("Test Redis Expire End. %v %v ", err, ballotX)
		}
		if affect == 0 {
			t.Errorf("Test Redis Expire End. %v %v ", affect, ballotX)
		} else {
			t.Logf("Test Redis Expire End. %v %v ", affect, checkMark)
		}
	}

	{
		t.Log("Test Redis Exists Start")
		has, err := redis.Exists("Test")
		if err != nil {
			t.Fatalf("Test Redis Exists End. %v %v ", err, ballotX)
		}
		if !has {
			t.Errorf("Test Redis Exists End. %v %v ", has, ballotX)
		} else {
			t.Logf("Test Redis Exists End. %v %v ", has, checkMark)
		}
	}

	{
		t.Log("Test Redis Delete Start")
		affect, err := redis.Delete("Test")
		if err != nil {
			t.Fatalf("Test Redis Delete End. %v %v ", err, ballotX)
		}
		if affect == 0 {
			t.Errorf("Test Redis Delete End. %v %v ", affect, ballotX)
		} else {
			t.Logf("Test Redis Delete End. %v %v ", affect, checkMark)
		}
	}
}

func BenchmarkRedis(b *testing.B) {

	conf := &redis.RedisConfig{
		Address:     "192.168.2.99:6379",
		IdleTimeout: 30,
		MaxActive:   3,
		MaxIdle:     3,
	}

	redis.InitRedis(conf)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		redis.Set("Test"+strconv.Itoa(i), "Test", 30)
	}
}
