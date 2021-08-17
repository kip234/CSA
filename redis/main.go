package main

import (
	"github.com/gomodule/redigo/redis"
	"time"
)

//type Redis struct {
//	addr string
//	conn redis.Conn
//}
//
//func (r Redis)Init() (err error) {
//	r.conn,err=redis.Dial("tcp",r.addr)
//	return
//}
//
//func (r Redis)Set(args ...interface{}) (err error) {
//	_,err=r.conn.Do("SET",args)
//}

const (
	Addr        = "localhost:6379"
	IdLeTimeout = 5
	MaxIdle     = 20
	MaxActive   = 8
)

type OptionPool struct {
	addr        string
	idLeTimeout int
	maxIdle     int
	maxActive   int
}

type PoolExt interface {
	apply(*OptionPool)
}

type tempFunc func(pool *OptionPool)

type funcPoolExt struct {
	f tempFunc
}

func (f *funcPoolExt) apply(p *OptionPool) {
	f.f(p)
}
func NewFuncPoolExt(f tempFunc) *funcPoolExt {
	return &funcPoolExt{f: f}
}

type Client struct {
	Option OptionPool
	pool   *redis.Pool
}

var DefaultOption = OptionPool{
	addr:        Addr,
	idLeTimeout: IdLeTimeout,
	maxIdle:     MaxIdle,
	maxActive:   MaxActive,
}

func NewClient(op ...PoolExt) *Client {
	c := &Client{Option: DefaultOption}
	for _, p := range op {
		p.apply(&c.Option)
	}
	c.setRedisPool()
	return c
}

func (pc *Client) setRedisPool() {
	pc.pool = &redis.Pool{
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", pc.Option.addr)
			if conn == nil || err != nil {
				return nil, err
			}
			return conn, nil
		},
		MaxIdle:     pc.Option.maxIdle,                                  // 最大空闲连接数
		MaxActive:   pc.Option.maxActive,                                // 最大活跃连接数
		IdleTimeout: time.Second * time.Duration(pc.Option.idLeTimeout), // 连接等待时间
	}
}

func (pc *Client) Set(args ...interface{}) error {
	c := pc.pool.Get()
	defer c.Close()
	_, err := c.Do("SET", args...)
	if err != nil {
		return err
	}
	return nil
}

func (pc *Client) Get(key string) (interface{}, error) {
	c := pc.pool.Get()
	defer c.Close()
	v, err := c.Do("GET", key)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func main() {
	
}

