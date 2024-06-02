package main

import (
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) <= 1 {
		log.Println("No arguments given..")
		return
	}

	switch strings.ToUpper(os.Args[1])[0] {
	case 'C': RunClient()
	case 'S': RunServer()
	}
}
