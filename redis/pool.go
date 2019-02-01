package redis

import (
	"time"

	redigo "github.com/gomodule/redigo/redis"
)

// RedisConfig Redis配置信息
type RedisConfig struct {
	Address     string
	Password    string
	Db          string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}

var pool *redigo.Pool

// InitRedis Redis初始化
func InitRedis(conf *RedisConfig) {
	pool = &redigo.Pool{
		MaxIdle:     conf.MaxIdle,
		MaxActive:   conf.MaxActive,
		IdleTimeout: (conf.IdleTimeout * time.Second),
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
	}
}
