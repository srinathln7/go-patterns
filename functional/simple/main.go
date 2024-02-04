package main

import "fmt"

type Server struct {
	id    int
	addr  string
	isTLS bool
}

func newServer(id int, addr string, isTLS bool) *Server {
	return &Server{
		id:    id,
		addr:  addr,
		isTLS: isTLS,
	}
}

func main() {
	// Problem with this approach is as the number of server configs grow,
	// we'll have to keep updating this file.
	// Can we do BETTER?
	server := newServer(1, "127.0.0.1", false)
	fmt.Printf("%+v \n", server)
}
