package main

import (
	"./netUtils"
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

const (
	connected = 1
	disconnected = -1
)

type clientConnectState struct {
	action int // connected or disconnected
	addr net.Addr
}

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

	connectionChan := make(chan clientConnectState)
	counter := 0
	go clientCounterHandler(counter, connectionChan)

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatalln(err)
		} else {
			go clientHandler(conn, connectionChan)
		}
	}
}

func clientHandler(conn net.Conn, clientCounter chan clientConnectState) {
	var (
		addr		= conn.RemoteAddr()
		buff		= make([]byte, 512)
		reader		= bufio.NewReader(conn)
		writer		= bufio.NewWriter(conn)
		err 		error
		readBytes	int
	)

	defer func () {
		clientCounter <- clientConnectState{disconnected, addr}
		conn.Close()
	}()

	clientCounter <- clientConnectState{connected, addr}

	for err == nil || readBytes <= 0 {
		readBytes, err = reader.Read(buff)
		if err != nil {
			break
		}
		fmt.Printf("Client %s, SAID: %s\n", addr, string(buff[:readBytes]))
		_, err = writer.Write(buff[:readBytes])
		if err != nil {
			break
		}
		err = writer.Flush()
		if err != nil {
			break
		}
	}
}

func clientCounterHandler(counter int, countReceiver chan clientConnectState) {
	var value clientConnectState
	for {
		value = <- countReceiver
		counter += value.action
		switch value.action {
		case connected:
			log.Printf("Client %s connected, total number of connected clients: %d", value.addr, counter)
		case disconnected:
			log.Printf("Client %s disconnected, total number of connected clients: %d", value.addr, counter)
		}
	}
}