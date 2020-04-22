package main

import (
	"fmt"
	"golang.org/x/net/ipv4"
	"log"
	"net"
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
	en0, err := net.InterfaceByName("en0") // da hell is that?



	conn, err := net.ListenPacket("udp4", group.String())
	checkErr(err)
	defer conn.Close()

	packetConn := ipv4.NewPacketConn(conn)
	packetConn.SetMulticastLoopback(false)
	if err := packetConn.JoinGroup(en0, group); err != nil {
		checkErr(err)
	}
	defer packetConn.LeaveGroup(en0, group)

	go ping(group)

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
	conn, err := net.DialUDP("udp", nil, addr)
	checkErr(err)
	for {
		message := "Ping Pong"
		_, err = conn.Write([]byte(message))
		checkErr(err)
		time.Sleep(1 * time.Second)
	}
}