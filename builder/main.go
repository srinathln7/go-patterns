package main

import "fmt"

type Config struct {
	addr    string
	maxConn int
	isTLS   bool
}

type server struct {
	config *Config
}

func (c Config) String() string {
	return fmt.Sprintf("{addr: %s, maxConn: %d, isTLS: %t}", c.addr, c.maxConn, c.isTLS)
}

func (c *Config) withAddr(addr string) *Config {
	c.addr = addr
	return c
}

func (c *Config) withMaxConn(num int) *Config {
	c.maxConn = num
	return c
}

func (c *Config) withTLS() *Config {
	c.isTLS = true
	return c
}

func defaultConfig() *Config {
	return &Config{
		addr:    ":8080",
		maxConn: 1,
		isTLS:   false,
	}
}

func newServer() *server {
	return &server{
		config: defaultConfig(),
	}
}

func newServerWithConfig(config *Config) *server {
	return &server{
		config: config,
	}
}

func main() {
	// Compare and contrast this with functional pattern
	config := defaultConfig().
		withAddr(":80").
		withTLS().
		withMaxConn(100)

	defaultserver := newServer()
	server := newServerWithConfig(config)

	// OUTPUT
	// 	default server :{addr: :8080, maxConn: 1, isTLS: false}
	// configured server :&main.Config{addr:":80", maxConn:100, isTLS:true}
	fmt.Printf("default server :%+v\n", defaultserver.config)
	fmt.Printf("configured server :%#v\n", server.config)
}
