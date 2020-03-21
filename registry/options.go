package registry

import "time"

/*
	选项设计模式，用来提供多组件方法
*/

type Options struct {
	Addrs        []string      //地址
	Timeout      time.Duration //超时时间
	HeartBeat    int64         //心跳时间
	RegistryPath string        //注册地址
}

type Option func(opts *Options)

func WithAddrs(addrs []string) Option {
	return func(opts *Options) {
		opts.Addrs = addrs
	}
}

func WithTimeout(timeout time.Duration) Option {
	return func(opts *Options) {
		opts.Timeout = timeout
	}
}

func WithHeartBeat(heartBeat int64) Option {
	return func(opts *Options) {
		opts.HeartBeat = heartBeat
	}
}

func WithRegister(registryPath string) Option {
	return func(opts *Options) {
		opts.RegistryPath = registryPath
	}
}
