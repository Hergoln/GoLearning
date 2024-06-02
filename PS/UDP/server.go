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

	buff := make([]byte, 512)
	udp, err := net.ListenPacket("udp", strings.Join([]string{addr, port}, ":"))

	if err != nil {
		log.Fatal("IP address and host unavailable")
	}

	defer udp.Close()

	for {
		bytes, addr, err := udp.ReadFrom(buff)
		if err != nil {
			log.Fatal("Something gone wrong")
		}
		fmt.Printf("Client %s, SAID: %s\n", addr.String(), string(buff[:bytes]))
		go handlePack(buff[:bytes], udp, addr)
	}
}

func handlePack(buff []byte, conn net.PacketConn, addr net.Addr) {
	_, err := conn.WriteTo(buff, addr)

	if err != nil {
		log.Printf("Could not send file back to address %s\n", addr.String())
	}

}
