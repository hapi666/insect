package model

import (
	"errors"
	"time"

	"github.com/gomodule/redigo/redis"
)

type DB interface {
	PushQueue(url []string) error
	PopQueue() (string, error)
	RangeQueue() ([]string, error)
	SisMember(URLmd5 string) (bool, error)
	SadD(crawledURL []string) error
}

type Pool struct {
	redis.Pool
}

type Conn struct {
	redis.Conn
}

func NewRedisPool(server, password string) (*Pool, error) {
	if server == "" {
		server = ":6379"
	}
	p := redis.Pool{
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			if password != "" {
				if _, err := conn.Do("AUTH", password); err != nil {
					conn.Close()
					return nil, err
				}
			}
			return conn, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
	}
	return &Pool{p}, nil
}

func (p *Pool) GetRedisConn() (DB, error) {
	conn := p.Get()
	if conn.Err() != nil {
		return nil, conn.Err()
	}
	return &Conn{conn}, nil
}

func (c *Conn) PushQueue(urls []string) error {
	defer c.Close()
	_, err := c.Do("rpush", "queue", urls)
	if err != nil {
		return err
	}
	return nil
}

func (c *Conn) PopQueue() (string, error) {
	defer c.Close()
	result, err := c.Do("lpop", "queue")
	if err == redis.ErrNil {
		return "", redis.ErrNil
	} else if err != nil {
		return "", err
	}
	if url, ok := result.(string); ok {
		return url, nil
	}
	return "", errors.New("类型断言失败")
}

func (c *Conn) RangeQueue() ([]string, error) {
	defer c.Close()
	result, err := c.Do("lrange", "queue", "0", "-1")
	if err == redis.ErrNil {
		return nil, redis.ErrNil
	} else if err != nil {
		return nil, err
	}
	if urls, ok := result.([]string); ok {
		return urls, nil
	}
	return nil, errors.New("类型断言失败")
}

func (c *Conn) SisMember(URLmd5 string) (bool, error) {
	defer c.Close()
	isExist, err := c.Do("sismember", "urlmd5", URLmd5)
	if err != redis.ErrNil {
		return false, err
	}
	if i, ok := isExist.(int); ok {
		if i == 0 {
			return false, nil
		} else {
			return true, err
		}
	}
	return false, errors.New("类型断言错误")
}

func (c *Conn) SadD(crawledURL []string) error {
	defer c.Close()
	_, err := c.Do("sadd", "urlmd5", crawledURL)
	if err != redis.ErrNil {
		return err
	}
	return redis.ErrNil
}
