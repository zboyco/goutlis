package goutlis

import (
	"time"

	redigo "github.com/gomodule/redigo/redis"
)

// RedisConfig Redis配置信息
type RedisConfig struct {
	Address     string
	Password    string
	Db          string
	MaxIdle     int           // 最大空闲连接数
	MaxActive   int           // 最大连接数
	IdleTimeout time.Duration // 空闲连接超时时间
	Wait        bool          // 如果超过最大连接，是否等待，不等待则报错
}

// Redis Redis结构
type Redis struct {
	pool *redigo.Pool
}

// InitRedis Redis初始化
func InitRedis(conf *RedisConfig) *Redis {
	return &Redis{
		pool: &redigo.Pool{
			MaxIdle:     conf.MaxIdle,
			MaxActive:   conf.MaxActive,
			IdleTimeout: conf.IdleTimeout * time.Second,
			Wait:        conf.Wait,
			Dial: func() (redigo.Conn, error) {
				c, err := redigo.Dial("tcp", conf.Address)
				if err != nil {
					return nil, err
				}
				if len(conf.Password) != 0 {
					if _, err := c.Do("AUTH", conf.Password); err != nil {
						c.Close()
						return nil, err
					}
				}
				if len(conf.Db) != 0 {
					if _, err := c.Do("SELECT", conf.Db); err != nil {
						c.Close()
						return nil, err
					}
				}
				r, err := redigo.String(c.Do("PING"))
				if err != nil || r != "PONG" {
					return nil, err
				}
				return c, err
			},
			TestOnBorrow: func(c redigo.Conn, t time.Time) error {
				if time.Since(t) < time.Minute {
					return nil
				}
				_, err := c.Do("PING")
				return err
			},
		},
	}
}

// Set 设置值
func (r *Redis) Set(key string, content interface{}, ex ...int64) (string, error) {
	c := r.pool.Get()
	defer c.Close()
	if len(ex) == 0 || ex[0] == 0 {
		return redigo.String(c.Do("SET", key, content))
	}
	return redigo.String(c.Do("SET", key, content, "EX", ex[0]))
}

// Expire 设置过期时间
func (r *Redis) Expire(key string, ex int64) (int64, error) {
	c := r.pool.Get()
	defer c.Close()
	return redigo.Int64(c.Do("EXPIRE", key, ex))
}

// Get 获取值
func (r *Redis) Get(key string) ([]byte, error) {
	c := r.pool.Get()
	defer c.Close()
	return redigo.Bytes(c.Do("GET", key))
}

// Exists Key存在
func (r *Redis) Exists(key string) (bool, error) {
	c := r.pool.Get()
	defer c.Close()
	return redigo.Bool(c.Do("EXISTS", key))
}

// Delete Key删除
func (r *Redis) Delete(key string) (int64, error) {
	c := r.pool.Get()
	defer c.Close()
	return redigo.Int64(c.Do("DEL", key))
}

// Keys 查询列表
func (r *Redis) Keys(key string) ([]string, error) {
	c := r.pool.Get()
	defer c.Close()
	return redigo.Strings(c.Do("KEYS", key))
}
