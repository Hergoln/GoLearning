package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

const (
	NAME	= "Test client"
	PORT	= "2137"
	ADDR	= "127.0.0.1"
	WHITE = "\r\n\t\x20"
)

// dunno how to not send eol using reading methods without modifying read string
func main() {
	consoleRead := bufio.NewScanner(os.Stdin)

	fmt.Print("Address: ")
	consoleRead.Scan()
	addr := consoleRead.Text()
	fmt.Print("Port: ")
	consoleRead.Scan()
	port := consoleRead.Text()

	connectString := strings.Join([]string{strings.TrimSpace(string(addr)), strings.TrimSpace(string(port))}, ":")

	fmt.Println("Waiting for connection ...")
	//connectString := strings.Join([]string{ADDR, PORT}, ":")
	conn, err := net.Dial("tcp", connectString)
	if err != nil {
		fmt.Println("Couldn't connect to server")
		fmt.Println(err)
		return
	}

	defer conn.Close()

	connReader := bufio.NewReader(conn)
	connWriter := bufio.NewWriter(conn)

	buff := make([]byte, 512)
	for {
		log.Println("Enter data")
		if !consoleRead.Scan() {
			return
		}
		text := consoleRead.Text()
		text = strings.TrimSpace(text)
		connWriter.Write([]byte(text))
		connWriter.Flush()
		n, _ := connReader.Read(buff)
		log.Print(string(buff[:n]))
	}
}

