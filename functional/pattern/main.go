package main

import "fmt"

type Config func(*Opts)

type Opts struct {
	id      int
	addr    string
	maxConn uint
	isTLS   bool
}

type Server struct {
	opts Opts
}

func defaultOpts() Opts {
	return Opts{
		maxConn: 1,
		isTLS:   false,
	}
}

func withTLS(opts *Opts) {
	opts.isTLS = true
}

func maxConnections(num uint) Config {
	return func(opts *Opts) {
		opts.maxConn = num
	}
}

func newServer(configs ...Config) *Server {
	opts := defaultOpts()
	for _, config := range configs {
		config(&opts)
	}

	return &Server{
		opts: opts,
	}
}

func main() {
	server := newServer()
	fmt.Printf("%+v \n", server)

	server1 := newServer(withTLS)
	fmt.Printf("%+v \n", server1)

	server2 := newServer(withTLS, maxConnections(100))
	fmt.Printf("%+v \n", server2)
}
