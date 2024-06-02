package main

import (
	"context"
	"encoding/binary"
	"errors"
	"fmt"
	"net"
	"os"
	"syscall"
)

const (
	BROADLAN = "192.168.0.255"
	BROADLOOPBACK = "169.254.255.255"
)

func main() {
	addr := getIPs()[1]
	listenBroad(addr)
}

func listenBroad(addr net.IP) {
	udpA, err := net.ResolveUDPAddr("udp", addr.String() + ":2137")
	checkErr(err)
	fmt.Println("Listen and obey " + udpA.String())

	config := &net.ListenConfig{Control: reusePort}

	conn, err := config.ListenPacket(context.Background(), "udp", udpA.String())
	checkErr(err)
	buff := make([]byte, 512)
	for {
		count, addr, err := conn.ReadFrom(buff)
		checkErr(err)
		fmt.Println(addr.String() + ": " + string(buff[:count]))
	}

	//buff := make([]byte, 512)
	//sock, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_DGRAM, syscall.IPPROTO_UDP)
	//checkErr(err)
	//err = syscall.SetsockoptInt(sock, syscall.SOL_SOCKET, syscall.SO_REUSEADDR, 1)
	//checkErr(err)
	//udpA, err := net.ResolveUDPAddr("udp", ":2137")
	//checkErr(err)
	//err = syscall.Bind(sock, &syscall.SockaddrInet4{Port: udpA.Port})
	//checkErr(err)
	//file := os.NewFile(uintptr(sock), string(sock))
	//conn, err := net.FilePacketConn(file)
	//checkErr(err)
	//fmt.Println("Listen and obey " + conn.LocalAddr().String())
	//err = file.Close()
	//checkErr(err)
	//for {
	//	n, remoteAddr, _ := conn.ReadFrom(buff)
	//	fmt.Printf("%s said : %s\n", remoteAddr.String(), n)
	//}

	//sock, err := windows.Socket(windows.AF_INET, windows.SOCK_DGRAM, windows.IPPROTO_UDP)
	//checkErr(err)
	//err = windows.SetsockoptInt(sock, windows.SOL_SOCKET, windows.SO_REUSEADDR, 1)
	//checkErr(err)
	//udpA, err := net.ResolveUDPAddr("udp", addr.String() + ":2137")
	//checkErr(err)
	//err = windows.Bind(sock, &windows.SockaddrInet4{Port: udpA.Port})
	//checkErr(err)
	//file := os.NewFile(uintptr(sock), string(sock))
	//conn, err := net.FilePacketConn(file)
	//checkErr(err)
	//fmt.Println("Listen and obey " + conn.LocalAddr().String())
	//err = file.Close()
	//checkErr(err)
	//for {
	//	n, remoteAddr, _ := conn.ReadFrom(buff)
	//	fmt.Printf("%s said : %s\n", remoteAddr.String(), n)
	//}
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
		panic(err)
	}
}

func reusePort(network, address string, conn syscall.RawConn) error {
	return conn.Control(func(descriptor uintptr) {
		syscall.SetsockoptInt(syscall.Handle(descriptor), syscall.SOL_SOCKET, syscall.SO_REUSEADDR, 1)
	})
}