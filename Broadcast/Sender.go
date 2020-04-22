package main

import (
	"encoding/binary"
	"errors"
	"fmt"
	"net"
	"os"
)

const (
	BROADLAN = "192.168.0.255"
	BROADLOOPBACK = "169.254.255.255"
)

func main() {
	addr := getIPs()[1]
	//fmt.Println(addr)
	//sendBroadMessage(addr)

	addrA := &net.IPNet{IP: addr, Mask: addr.DefaultMask()}
	broadAddr, _ := lastAddr(addrA)
	fmt.Println(broadAddr)
	sendBroadMessage(broadAddr)
}

func sendBroadMessage(addr net.IP) {
	udpA, err := net.ResolveUDPAddr("udp", addr.String() + ":2137")
	checkErr(err)

	fmt.Print(udpA)
	conn, err := net.DialUDP("udp", nil, udpA)
	checkErr(err)

	_, err = conn.Write([]byte("anything"))
	defer conn.Close()
	checkErr(err)
}

func getIPs() []net.IP {
	toReturn := make([]net.IP, 0)
	host, _ := os.Hostname()
	addrs, _ := net.LookupIP(host)
	for _, addr := range addrs {
		if ipv4 := addr.To4(); ipv4 != nil && !ipv4.IsLoopback() {
			toReturn = append(toReturn, ipv4)
		}
	}
	return toReturn
}

// this one is not mine but its relatively clean way of creating broadcast IP from an IP address
func lastAddr(n *net.IPNet) (net.IP, error) {
	if n.IP.To4() == nil {
		return net.IP{}, errors.New("does not support IPv6 addresses")
	}
	ip := make(net.IP, len(n.IP.To4()))
	binary.BigEndian.PutUint32(ip, binary.BigEndian.Uint32(n.IP.To4())|^binary.BigEndian.Uint32(net.IP(n.Mask).To4()))
	return ip, nil
}

func checkErr(err error) {
	if err != nil {
		fmt.Print(err)
	}
}