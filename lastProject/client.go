package main

import (
	"bufio"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

func RunClient() {
	var (
		err           error
		broadLoopback = "192.168.56.255"
	)

	udpAddr, err := net.ResolveUDPAddr("udp", broadLoopback+":"+strconv.Itoa(UdpPort))
	checkErr(err)

	conn, err := ListenUDP(UdpPort)
	checkErr(err)
	defer func() {
		log.Println("Closing UDP client side connection...")
		conn.Close()
	}()
	HandleCommunication(udpAddr, conn)
}

func HandleCommunication(udpAddr *net.UDPAddr, udpConn net.PacketConn) {
	var (
		err       error
		message   string
		addr      string
		addrs     []string
		addrIndex int
		buff      = make([]byte, 512)
	)

	for {
		_, err = udpConn.WriteTo([]byte(DiscoverMessage), udpAddr)
		checkErr(err)
		count, _, err := udpConn.ReadFrom(buff)
		checkErr(err)
		message = string(buff[:count])
		addr = handleResponse(message)
		if addr != "" {
			addrs = addAddrsToTable(addrs, addr)
			addrIndex = chooseAddr(addrs)
			if addrIndex >= 0 {
				err = handleTcpConnection(addrs[addrIndex])
				if err != nil {
					log.Println("Connection lost")
				} else {
					log.Println("Connection closed")
				}
			} else {
				log.Println("Waiting 10s")
				time.Sleep(10 * time.Second)
			}
		}
	}
}

func handleResponse(response string) string {
	responseCopy := response[:]
	if !strings.HasPrefix(responseCopy, OfferPrefix) {
		return ""
	}
	responseCopy = responseCopy[len(OfferPrefix)+1:]
	_, err := net.ResolveTCPAddr("tcp", responseCopy)
	if err != nil {
		return ""
	}
	return responseCopy
}

func addAddrsToTable(table []string, addr ...string) []string {
	for _, each := range addr {
		if !stringArrayContains(table, each) {
			table = append(table, each)
		}
	}
	return table
}

func stringArrayContains(array []string, el string) bool {
	for _, each := range array {
		if each == el {
			return true
		}
	}
	return false
}

func chooseAddr(addrs []string) int {
	log.Printf("%d Wait", 0)
	for index, addr := range addrs {
		log.Printf("%d %s\n", index+1, addr)
	}
	choice := readValueInRange(0, len(addrs))
	return choice - 1
}

func readValueInRange(min, max int) int {
	var (
		consoleRead = bufio.NewScanner(os.Stdin)
		response    string
		choice      = -1
	)
	log.Printf("Values between %d - %d", min, max)
	consoleRead.Scan()
	for {
		response = consoleRead.Text()
		choice = validateNumber(response, min, max)
		if choice != -1 {
			return choice
		}
		log.Printf("Incorrect value (min: %d, max: %d)\n", min, max)
	}
}

func validateNumber(number string, minValue int, maxValue int) int {
	ind, err := strconv.Atoi(number)
	if err != nil || ind < minValue || ind > maxValue {
		return -1
	}
	return ind
}

func handleTcpConnection(addr string) error {
	const (
		MaxRefreshRate = 1000
		MinRefreshRate = 10
		TimeRequest    = "TIME"
	)
	var (
		refreshRate int
		connReader  *bufio.Reader
		connWriter  *bufio.Writer
		T1          int64
		Tserv       int64
		Tcli        int64
		delta       int64
		buff        = make([]byte, 256)
	)
	conn, err := net.Dial("tcp", addr)
	defer conn.Close()
	if err != nil {
		log.Println("Couldn't connect to server")
		return err
	}
	log.Printf("Connected to %s", addr)

	refreshRate = readValueInRange(MinRefreshRate, MaxRefreshRate)

	connReader = bufio.NewReader(conn)
	connWriter = bufio.NewWriter(conn)

	for {
		T1 = time.Now().UnixNano()
		_, err = connWriter.Write([]byte(TimeRequest))
		if err != nil {
			return err
		}
		connWriter.Flush()
		if err != nil {
			return err
		}
		bytesRead, _ := connReader.Read(buff)
		Tserv, err = int64FromResponse(buff, bytesRead)
		if err != nil {
			return err
		}
		Tcli = time.Now().UnixNano()
		delta = Tserv + (Tcli - T1)/2 - Tcli
		log.Printf("Time: %s, delta: %d", time.Unix(0, Tcli+delta), delta)
		if readUserInput() {
			return nil
		}
		time.Sleep(time.Duration(refreshRate) * time.Millisecond)
	}
}

func int64FromResponse(buff []byte, bytesRead int) (int64, error) {
	stringRes := string(buff[:bytesRead])
	num, err := strconv.ParseInt(stringRes, 10, 64)
	return num, err
}

func readUserInput() bool {
	consoleRead := bufio.NewScanner(os.Stdin)
	log.Print("Do you want to exit? (Y/N)")
	consoleRead.Scan()
	input := consoleRead.Text()
	if input[0] == 'Y' || input[0] == 'y' {
		return true
	}
	return false
}
