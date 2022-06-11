package main

import (
	"io"
	"log"
	"net"
)

func main() {

	conn, err := net.Dial("tcp", ":9999")

	if err != nil {
		log.Println(err)
	}

	agent, err := net.Dial("tcp", ":80")

	if err != nil {
		log.Println(err)
	}

	// x
	go io.Copy(agent, conn)
	io.Copy(conn, agent)
}
