package main

import (
	util "../netUtils"
	"context"
	"log"
	"net"
	"strconv"
	"strings"
)

func RunUDPServer(addr *net.UDPAddr, tcpAddrs []string, closingChan chan int) {
	var (
		err       error
		message   string
		conn net.PacketConn
	)
	conn, err = ListenUDP(UdpPort)
	defer func() {
		log.Println("Closing UDP connection...")
		conn.Close()
	}()
	checkErr(err)
	log.Println("Listen and Obey " + addr.String())
	for {
		buff := make([]byte, 512)
		count, _, err := conn.ReadFrom(buff)
		message = string(buff[:count])
		checkErr(err)
		log.Println(message)
		if strings.HasPrefix(message, DiscoverMessage) {
			for _, each := range tcpAddrs {
				conn.WriteTo([]byte(OfferPrefix + " " + each), addr)
			}
		}
	}
}

func ListenUDP(port int) (net.PacketConn, error) {
	config := &net.ListenConfig{Control: util.ReusePort}
	return config.ListenPacket(context.Background(), "udp", ":" + strconv.Itoa(port))
}