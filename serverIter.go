package main

import (
	"./netUtils"
	"bufio"
	"log"
	"net"
	"os"
	"strings"
)

func main() {

	var response, addr, port string

	consoleRead := bufio.NewScanner(os.Stdin)

	log.Printf("Would you like to use default address(%s:%s)? Y/N:", netUtils.DEFAULT_ADDR, netUtils.DEFAULT_PORT)
	consoleRead.Scan()
	response = consoleRead.Text()
	if strings.EqualFold(response, "Y") {
		addr = netUtils.DEFAULT_ADDR
		port = netUtils.DEFAULT_PORT
	} else {
		for {
			log.Print("Address: ")
			consoleRead.Scan()
			addr = consoleRead.Text()
			log.Print("Port: ")
			consoleRead.Scan()
			port = consoleRead.Text()
			if !netUtils.CheckAddressCorrectness(addr, port) {
				log.Println("Incorrect Address or port")
			} else {
				break
			}
		}
	}

	log.Println("Starting server...")
	listen, err := net.Listen("tcp", strings.Join([]string{addr, port}, ":"))

	if err != nil {
		log.Printf("Socket listen port %s, failed, %s", port, err)
		return
	}

	defer listen.Close()

	log.Printf("Begin listening on port %s...", port)

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatalln(err)
		} else {
			clientHandlerIter(conn)
		}
	}
}

func clientHandlerIter(conn net.Conn) {
	defer conn.Close()

	var (
		addr		= conn.RemoteAddr()
		buff		= make([]byte, 512)
		reader		= bufio.NewReader(conn)
		writer		= bufio.NewWriter(conn)
		err 		error
		readBytes	int
	)

	log.Printf("Client connected at port: %s\n", addr)

	for err == nil || readBytes <= 0 {
		readBytes, err = reader.Read(buff)
		if err != nil {
			break
		}
		log.Printf("Client %s, SAID: %s", addr, string(buff[:readBytes]))
		_, err = writer.Write(buff[:readBytes])
		if err != nil {
			break
		}
		err = writer.Flush()
		if err != nil {
			break
		}
	}

	log.Printf("Client at ported %s, disconnected", addr)
}