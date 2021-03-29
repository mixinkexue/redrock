package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"time"
)

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
func (pc *Client) TTL(key string) (interface{},error){
	c:=pc.pool.Get()
	defer c.Close()
	v, err := c.Do("TTL", key)
	if err != nil {
		return nil, err
	}
	return v, nil
}
func (pc *Client) HSET(args ...interface{})(interface{},error){
	c:=pc.pool.Get()
	defer c.Close()
	v,err:=c.Do("HSET",args...)
	if err != nil {
		return nil, err
	}
	return v, nil
}
func (pc *Client) HGET(args ...interface{})(interface{},error){
	c := pc.pool.Get()
	defer c.Close()
	v, err := c.Do("HGET",args...)
	if err != nil {
		return nil, err
	}
	return v, nil
}
func (pc *Client) HLEN(hash string)(interface{},error){
	c := pc.pool.Get()
	defer c.Close()
	v, err := c.Do("HLEN", hash)
	if err != nil {
		return nil, err
	}
	return v, nil
}
func (pc *Client) LPUSH(args ...interface{})(interface{},error){
	c := pc.pool.Get()
	defer c.Close()
	v, err := c.Do("LPUSH",args...)
	if err != nil {
		return nil, err
	}
	return v, nil
}
func (pc *Client) LRANGE(args ...interface{})(interface{},error){
	c := pc.pool.Get()
	defer c.Close()
	v, err := c.Do("LRANGE",args...)
	if err != nil {
		return nil, err
	}
	return v, nil
}
func (pc *Client) LPOP(key string)(interface{},error){
	c := pc.pool.Get()
	defer c.Close()
	v, err := c.Do("LPOP",key)
	if err != nil {
		return nil, err
	}
	return v, nil
}
func (pc *Client) LSET(args ...interface{})(interface{},error){
	c := pc.pool.Get()
	defer c.Close()
	v, err := c.Do("LSET",args...)
	if err != nil {
		return nil, err
	}
	return v, nil
}
func (pc *Client) SADD(args ...interface{})(interface{},error){
	c := pc.pool.Get()
	defer c.Close()
	v, err := c.Do("SADD",args...)
	if err != nil {
		return nil, err
	}
	return v, nil
}
func (pc *Client) SMEMBERS(key string)(interface{},error){
	c := pc.pool.Get()
	defer c.Close()
	v, err := c.Do("SMEMBERS",key)
	if err != nil {
		return nil, err
	}
	return v, nil
}
func (pc *Client) SPOP(key string)(interface{},error){
	c := pc.pool.Get()
	defer c.Close()
	v, err := c.Do("SPOP",key)
	if err != nil {
		return nil, err
	}
	return v, nil
}
func (pc *Client) SMOVE(args ...interface{})(interface{},error){
	c := pc.pool.Get()
	defer c.Close()
	v, err := c.Do("SMOVE",args...)
	if err != nil {
		return nil, err
	}
	return v, nil
}
func (pc *Client) ZADD(args ...interface{})(interface{},error){
	c := pc.pool.Get()
	defer c.Close()
	v, err := c.Do("ZADD",args...)
	if err != nil {
		return nil, err
	}
	return v, nil
}
func (pc *Client) ZRANGE(args ...interface{})(interface{},error){
	c := pc.pool.Get()
	defer c.Close()
	v, err := c.Do("ZRANGE",args...)
	if err != nil {
		return nil, err
	}
	return v, nil
}
func (pc *Client) ZREM(args ...interface{})(interface{},error){
	c := pc.pool.Get()
	defer c.Close()
	v, err := c.Do("ZREM",args...)
	if err != nil {
		return nil, err
	}
	return v, nil
}
func main()  {
	c:=NewClient()
	var rp interface{}
	var err error
	c.ZADD("key2","1","m","2","n")
	rp,err=c.ZREM("key2","m")
	if err != nil {
		panic(err)
	}
	rp,err=c.ZRANGE("key2","0","-1")
	if err != nil {
		panic(err)
	}

	fmt.Println("key", rp)
}