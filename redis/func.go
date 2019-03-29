package redis

import (
	redigo "github.com/gomodule/redigo/redis"
)

// Set 设置值
func Set(key string, content interface{}, ex ...int64) (string, error) {
	c := pool.Get()
	defer c.Close()
	if len(ex) == 0 || ex[0] == 0 {
		return redigo.String(c.Do("SET", key, content))
	}
	return redigo.String(c.Do("SET", key, content, "EX", ex[0]))
}

// Expire 设置过期时间
func Expire(key string, ex int64) (int64, error) {
	c := pool.Get()
	defer c.Close()
	return redigo.Int64(c.Do("EXPIRE", key, ex))
}

// Get 获取值
func Get(key string) ([]byte, error) {
	c := pool.Get()
	defer c.Close()
	return redigo.Bytes(c.Do("GET", key))
}

// Exists Key存在
func Exists(key string) (bool, error) {
	c := pool.Get()
	defer c.Close()
	return redigo.Bool(c.Do("EXISTS", key))
}

// Delete Key删除
func Delete(key string) (int64, error) {
	c := pool.Get()
	defer c.Close()
	return redigo.Int64(c.Do("DEL", key))
}

// Keys 查询列表
func Keys(key string) ([]string, error) {
	c := pool.Get()
	defer c.Close()
	return redigo.Strings(c.Do("KEYS", key))
}
