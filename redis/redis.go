package redis

import (
	redigo "github.com/gomodule/redigo/redis"
	"time"
)

type Cache struct {
	pool *redigo.Pool
}

func New(options Options) *Cache {

	// 初始化
	options.Assign()

	// 建立连接池
	redigoClient := &redigo.Pool{
		MaxIdle:     options.MaxIdle,
		MaxActive:   options.MaxActive,
		IdleTimeout: options.IdleTimeout * time.Second,
		Wait:        true,
		Dial: func() (redigo.Conn, error) {
			// pool.get() call this function if not available connection
			con, err := redigo.Dial("tcp", options.Address,
				redigo.DialPassword(options.Password),
				redigo.DialDatabase(options.Db),
				redigo.DialConnectTimeout(options.ConnectTimeout*time.Second),
				redigo.DialReadTimeout(options.ReadTimeout*time.Second),
				redigo.DialWriteTimeout(options.WriteTimeout*time.Second))
			if err != nil {
				return nil, err
			}
			return con, nil
		},
		TestOnBorrow: func(c redigo.Conn, t time.Time) error {
			// less than IdleTimeout is ok
			if time.Since(t) < options.IdleTimeout*time.Second {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}

	return &Cache{
		pool: redigoClient,
	}
}

func (cache *Cache) Do(commandName string, args ...interface{}) (reply interface{}, err error) {
	conn := cache.pool.Get()
	defer conn.Close()

	return conn.Do(commandName, args...)
}

// Send writes the command to the client's output buffer.
func (cache *Cache) Send(commandName string, args ...interface{}) error {
	conn := cache.pool.Get()
	defer conn.Close()

	return conn.Send(commandName, args...)
}

// Flush flushes the output buffer to the redigo server.
func (cache *Cache) Flush() error {
	conn := cache.pool.Get()
	defer conn.Close()

	return conn.Flush()
}

// Receive receives a single reply from the redigo server
func (cache *Cache) Receive() (reply interface{}, err error) {
	conn := cache.pool.Get()
	defer conn.Close()

	return conn.Receive()
}
