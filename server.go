package main

import (
	"io"
	"log"
	"net"
	"strconv"
	"sync"
)

func server(port int, conn net.Conn, wg *sync.WaitGroup) {
	addr := ":" + strconv.Itoa(port)

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Println(err)
	}
	defer listener.Close()

	for {
		cli, err := listener.Accept()
		if err != nil {
			wg.Done()
			break
		}
		go io.Copy(conn, cli)
		go io.Copy(cli, conn)
	}

}
func main() {

	addr := ":9999"
	listener, err := net.Listen("tcp", addr)

	var wg sync.WaitGroup
	wg.Add(1)

	if err != nil {
		log.Println(err)
	}
	defer listener.Close()

	conn, err := listener.Accept()
	if err != nil {
		log.Println(err)
	}
	go server(81, conn, &wg)
	go server(82, conn, &wg)

	wg.Wait()

}
