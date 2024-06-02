package main

import (
	"bufio"
	"log"
	"net"
	"strconv"
	"time"
)

func RunTCPServer(addr string, closing chan int) {
	log.Printf("Starting server with address %s...\n", addr)

	listen, err := net.Listen("tcp", addr)

	if err != nil {
		log.Printf("Server with addr %s failed\n", addr)
		log.Println(err)
		return
	}

	defer func() {
		log.Printf("Closing TCP connection on addr %s...", addr)
		listen.Close()
	}()

	log.Printf("Begin listening (%s:TCP) ...\n", addr)
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatalln(err)
		} else {
			go clientHandler(conn)
		}
	}

}

func clientHandler(conn net.Conn) {
	var (
		buff		= make([]byte, 512)
		reader		= bufio.NewReader(conn)
		writer		= bufio.NewWriter(conn)
		err 		error
		readBytes	int
	)

	defer func() {
		log.Println("Client disconnected")
		conn.Close()
	}()

	for err == nil || readBytes <= 0 {
		readBytes, err = reader.Read(buff)
		if err != nil {
			break
		}

		response := string(buff[:readBytes])
		log.Printf("Client said: %s\n", response)
		if response == TimeRequest {
			_, err = writer.Write(nowInBytes())
			if err != nil {
				break
			}
			err = writer.Flush()
			if err != nil {
				break
			}
		}
	}
}

func nowInBytes() []byte {
	return []byte(strconv.FormatInt(time.Now().UnixNano(), 10))
}