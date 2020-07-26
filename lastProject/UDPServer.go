package main

import (
	util "../netUtils"
	"context"
	"log"
	"net"
	"strconv"
)

func RunUDPServer(addr *net.UDPAddr, closingChan chan int) {
	conn, err := ListenUDP(UdpPort)
	defer func() {
		log.Println("Closing UDP connection...")
		conn.Close()
	}()
	checkErr(err)
	log.Println("Listen and Obey " + addr.String())
	buff := make([]byte, 512)
	for {
		count, addr, err := conn.ReadFrom(buff)
		message := buff[:count]
		checkErr(err)
		log.Println(addr.String() + ": " + string(message))
		conn.WriteTo([]byte("kufa?"), addr)
	}
}

func ListenUDP(port int) (net.PacketConn, error) {
	config := &net.ListenConfig{Control: util.ReusePort}
	return config.ListenPacket(context.Background(), "udp", ":" + strconv.Itoa(port))
}