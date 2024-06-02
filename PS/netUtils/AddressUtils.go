package netUtils

import (
	"encoding/binary"
	"errors"
	"log"
	"net"
	"regexp"
	"strconv"
	"syscall"
)

const (
	DEFAULT_PORT	= "7"
	DEFAULT_ADDR	= "127.0.0.1"
	WHITESPACES		= "\r\n\t\x20"
)

// simple address correctness check
func CheckAddressCorrectness(address, port string) bool {
	addressMatch, _ := regexp.MatchString(`\b\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}\b`, address)
	portMatch, _ := regexp.MatchString(`\b\d{1,4}`, port)
	return addressMatch && portMatch
}

func ReusePort(network, address string, conn syscall.RawConn) error {
	return conn.Control(func(descriptor uintptr) {
		syscall.SetsockoptInt(syscall.Handle(descriptor), syscall.SOL_SOCKET, syscall.SO_REUSEADDR, 1)
	})
}


func LastAddr(n *net.Addr) (net.Addr, error) {
	IP, netIP, _ := net.ParseCIDR((*n).String())
	if IP.To4() == nil {
		return nil, errors.New("does not support IPv6 addresses")
	}
	ip := make(net.IP, len(IP.To4()))
	binary.BigEndian.PutUint32(ip, binary.BigEndian.Uint32(IP.To4())|^binary.BigEndian.Uint32(net.IP(netIP.Mask).To4()))
	number, _ := netIP.Mask.Size()
	var result net.Addr = MyAddr{
		network: "udp",
		addr: ip.String() + "/"+ strconv.Itoa(number),
	}
	return result, nil
}

func StringAddr(addr net.Addr, port int) string {
	IP, _, _ := net.ParseCIDR(addr.String())
	log.Println(IP.String() + ":" + strconv.Itoa(port))
	return IP.String() + ":" + strconv.Itoa(port)
}