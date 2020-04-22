package main

import (
	"../netUtils"
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	consoleRead := bufio.NewScanner(os.Stdin)

	//fmt.Printf("Would you like to send broadcast messege? Y/N")
	//response := consoleRead.Text()
	//if strings.EqualFold(response, "Y") {
	//	broadcastMessages(consoleRead)
	//	return
	//}

	connectString := resolveUsersChoice(consoleRead)

	fmt.Print("Waiting for connection... ")
	udpAddr , err := net.ResolveUDPAddr("udp", connectString)
	conn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		fmt.Println("Couldn't connect to server")
		fmt.Println(err)
		return
	}
	fmt.Println("Connected")
	defer conn.Close()

	for {
		if !consoleRead.Scan() {
			return
		}
		text := consoleRead.Text()
		text = strings.TrimSpace(text)
		go handleMessage(text, conn)
	}

}

func handleMessage(message string, conn *net.UDPConn) {
	bytesWritten, err := conn.Write([]byte(message))
	if err != nil {
		log.Print(err)
	}
	fmt.Printf("%d bytes written\n", bytesWritten)
	buff := make([]byte, 512)
	bytesRead, addrRecv, _ := conn.ReadFrom(buff)
	fmt.Printf("(%d, btytes read)Server(%s) response: %s\n",
		bytesRead,
		addrRecv.String(),
		string(buff[:bytesRead]))
}

//func broadcastMessages(consoleRead *bufio.Scanner) {
//	connectString := resolveUsersChoice(consoleRead)
//
//
//}

func resolveUsersChoice(consoleRead *bufio.Scanner) string {
	var addr, port, response string
	consoleRead.Scan()
	fmt.Printf("Would you like to use default address(%s:%s)? Y/N:",
		netUtils.DEFAULT_ADDR,
		netUtils.DEFAULT_PORT)
	response = consoleRead.Text()
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

	return strings.Join([]string{strings.TrimSpace(addr), strings.TrimSpace(port)}, ":")
}