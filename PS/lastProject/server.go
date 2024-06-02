package main

import (
	util "../netUtils"
	"fmt"
	"math/rand"
	"net"
)

const (
	ShutdownMessage = "Internet interfaces not found, shutting down application"
	EndSignal       = -1
	TimeRequest     = "TIME"
	UdpPort         = 7
	DiscoverMessage = "DISCOVER"
	OfferPrefix     = "OFFER ADDRESS PORT"
	RedColor        = "\033[31m"
	ResetColor      = "\033[0m"
)

func RunServer() {
	var (
		err           error
		closingChan   = make(chan int)
		pickablePorts = []int{7312, 7321, 7123, 1723, 7231, 2137}
		tcpPort       = -1
		addrs         []string
	)

	interfacesAddrs := ListOfPublicEthernetInterfacesIPs()
	if interfacesAddrs == nil {
		fmt.Print(ShutdownMessage)
		return
	}
	broad, err := util.LastAddrInNetwork(interfacesAddrs[0])
	if err != nil {
		fmt.Println(err)
		fmt.Print(ShutdownMessage)
		return
	}
	udpAddr, err := net.ResolveUDPAddr("udp", util.StringAddr(broad, UdpPort))
	if err != nil {
		fmt.Println(err)
		fmt.Print(ShutdownMessage)
		return
	}

	for _, addr := range interfacesAddrs[1:] {
		tcpPort, pickablePorts = randPort(pickablePorts)
		tcpAddr := util.StringAddr(*addr, tcpPort)
		addrs = addAddrsToTable(addrs, tcpAddr)
		go RunTCPServer(tcpAddr, closingChan)
	}
	defer func() {
		closingChan <- EndSignal
	}()
	RunUDPServer(udpAddr, addrs, closingChan)
}

func randPort(pickablePorts []int) (int, []int) {
	if pickablePorts == nil || len(pickablePorts) <= 0 {
		return -1, nil
	}
	drawn := rand.Intn(len(pickablePorts))
	return pickablePorts[drawn], append(pickablePorts[:drawn], pickablePorts[drawn:]...)
}

func ListOfPublicEthernetInterfacesIPs() []*net.Addr {
	IPs := make([]*net.Addr, 0)
	inters, _ := net.Interfaces()
	for _, inter := range inters {
		if inter.Flags&net.FlagUp != 0 &&
			inter.Flags&net.FlagLoopback == 0 {
			addr, _ := inter.Addrs()
			for _, each := range addr {
				IP, _, err := net.ParseCIDR(each.String())
				checkErr(err)
				if ipv4 := IP.To4(); ipv4 != nil {
					IPs = append(IPs, &each)
				}
			}
		}
	}
	return IPs
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
