package main

import (
	"bufio"
	"log"
	"net"
	"strings"
)

const (
	NAME = "Test server"
	PORT = "2137"
	ADDR = "127.0.0.1"
	WHITE = "\r\n\t\x20"
)

func main() {
	log.Println("Starting server...")

	listen, err := net.Listen("tcp", strings.Join([]string{ADDR, PORT}, ":"))

	if err != nil {
		log.Printf("Socket listen port %s, failed, %s", PORT, err)
		return
	}

	defer listen.Close()

	log.Printf("Begin listening on port %s...", PORT)

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatalln(err)
			continue
		}
		go clientHandler(conn)
	}
}

func clientHandler(conn net.Conn) {
	defer conn.Close()

	var (
		addr	= conn.LocalAddr().String()
		buff	= make([]byte, 512)
		r		= bufio.NewReader(conn)
		w		= bufio.NewWriter(conn)
		err error
		n int
	)

	log.Printf("Client connected at port: %s\n", addr)

	for err == nil || n <= 0 {
		n, err = r.Read(buff)
		if err != nil {
			break
		}
		log.Printf("Client %s, SAID: %s", addr, string(buff[:n]))
		n, err = w.Write(buff[:n])
		w.Flush()
	}

	log.Printf("Client at ported %s, disconnected", addr)
}