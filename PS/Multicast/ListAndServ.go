package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/ipv4"
	"log"
	"net"
	"os"
	"time"
)

const (
	groupAddr = "224.0.0.10"
	BUFFSIZE = 512
)

func main() {
	group, err := net.ResolveUDPAddr("udp", groupAddr + ":7")
	if err != nil {
		log.Fatal(err)
	}

	// listening
	conn, err := net.ListenPacket("udp4", group.String())
	checkErr(err)
	defer conn.Close()

	packetConn := ipv4.NewPacketConn(conn)
	if err := packetConn.JoinGroup(nil, group); err != nil {
		checkErr(err)
	}
	defer packetConn.LeaveGroup(nil, group)

	go ping(group)
	fmt.Println(packetConn.LocalAddr().String() + " Listen and obey")
	buff := make([]byte, BUFFSIZE)
	for {
		nBytes, controlMessage, src, err := packetConn.ReadFrom(buff)
		checkErr(err)
		if controlMessage != nil {
			fmt.Println(controlMessage)
		}
		fmt.Printf("%s says: %s\n", src, string(buff[0:nBytes]))
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func ping(addr *net.UDPAddr) {
	var response string
	conn, err := net.DialUDP("udp", nil, addr)
	checkErr(err)
	consoleRead := bufio.NewScanner(os.Stdin)
	fmt.Print("Say: ")
	for {
		consoleRead.Scan()
		response = consoleRead.Text()
		_, err = conn.Write([]byte(response))
		checkErr(err)
		time.Sleep(1 * time.Second)
	}
}