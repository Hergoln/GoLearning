package netUtils

import (
	"regexp"
)

const (
	DEFAULT_PORT	= "7"
	DEFAULT_ADDR	= "127.0.0.1"
	WHITESPACE 		= "\r\n\t\x20"
)

// simple address correctness check
func CheckAddressCorrectness(address, port string) bool {
	addressMatch, _ := regexp.MatchString(`\b\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}\b`, address)
	portMatch, _ := regexp.MatchString(`\b\d{1,4}`, port)
	return addressMatch && portMatch
}
