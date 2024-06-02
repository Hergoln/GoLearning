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

func main() {
	var addr, port string
	var bytesWritten int
	consoleRead := bufio.NewScanner(os.Stdin)

	fmt.Printf("Would you like to use default address(%s:%s)? Y/N:",
				netUtils.DEFAULT_ADDR,
				netUtils.DEFAULT_PORT)
	consoleRead.Scan()
	response := consoleRead.Text()
	if strings.EqualFold(response, "Y") {
		addr = netUtils.DEFAULT_ADDR
		port = netUtils.DEFAULT_PORT
	} else {
		for {
			fmt.Print("Address: ")
			consoleRead.Scan()
			addr = consoleRead.Text()
			fmt.Print("Port: ")
			consoleRead.Scan()
			port = consoleRead.Text()
			if !netUtils.CheckAddressCorrectness(addr, port) {
				log.Println("Incorrect Address or port")
			} else {
				break
			}
		}
	}

	connectString := strings.Join([]string{strings.TrimSpace(string(addr)), strings.TrimSpace(string(port))}, ":")

	fmt.Print("Waiting for connection... ")
	conn, err := net.Dial("tcp", connectString)
	if err != nil {
		fmt.Println("Couldn't connect to server")
		fmt.Println(err)
		return
	}
	fmt.Println("Connected")
	defer conn.Close()

	connReader := bufio.NewReader(conn)
	connWriter := bufio.NewWriter(conn)

	for {
		fmt.Print("To server: ")
		if !consoleRead.Scan() {
			return
		}
		text := consoleRead.Text()
		text = strings.TrimSpace(text)
		bytesWritten, err = connWriter.Write([]byte(text))
		if err != nil {

			break
		}
		fmt.Printf("%d bytes written\n", bytesWritten)
		err = connWriter.Flush()
		if err != nil {
			break
		}
		buff := make([]byte, 512)
		bytesRead, _ := connReader.Read(buff)
		fmt.Printf("(%d, btytes read)Server response: %s\n", bytesRead, string(buff[:bytesRead]))
	}
	log.Println(err)
	fmt.Println("Client disconnected")
}

