package main

import "fmt"

type Options struct {
	OptionString1 string
	OptionString2 string
	OptionString3 string
	OptionInt1    int
	OptionInt2    int
	OptionInt3    int
}

type Option func(opts *Options)

func InitOption(opts ...Option) {
	option := &Options{}
	for _, opt := range opts {
		opt(option)
	}
	fmt.Printf("init options %#v\n", option)
}

func WithOptionString1(str string) Option {
	return func(ops *Options) {
		ops.OptionString1 = str
	}
}

func WithOptionString2(str string) Option {
	return func(ops *Options) {
		ops.OptionString1 = str
	}
}

func main() {
	InitOption(WithOptionString1("jerry"), WithOptionString2("tom"))
}
