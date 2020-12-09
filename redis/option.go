package redis

import "time"

type Options struct {
	// host:port 形式的IP端口地址，或者 MOD:CMD 形式的 L5 地址
	Address        string
	Db             int
	Password       string
	ConnectTimeout time.Duration
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
	IdleTimeout    time.Duration
	MaxIdle        int
	MaxActive      int
}

var defaultOptions = Options{
	Db:             0,
	Password:       "",
	ConnectTimeout: 30,
	IdleTimeout:    300,
	MaxActive:      20,
	MaxIdle:        10,
	ReadTimeout:    10,
	WriteTimeout:   10,
}

func (options *Options) Assign() {
	if options.Db <= 0 {
		options.Db = defaultOptions.Db
	}
	if options.ConnectTimeout == 0 {
		options.ConnectTimeout = defaultOptions.ConnectTimeout
	}
	if options.IdleTimeout == 0 {
		options.IdleTimeout = defaultOptions.IdleTimeout
	}
	if options.MaxActive == 0 {
		options.MaxActive = defaultOptions.MaxActive
	}
	if options.MaxIdle == 0 {
		options.MaxIdle = defaultOptions.MaxIdle
	}
	if options.ReadTimeout == 0 {
		options.ReadTimeout = defaultOptions.ReadTimeout
	}
	if options.WriteTimeout == 0 {
		options.WriteTimeout = defaultOptions.WriteTimeout
	}
}