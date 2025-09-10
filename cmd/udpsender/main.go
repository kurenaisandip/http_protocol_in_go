package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

// to run this first run the net cat in another terminal and run the main function of this file then, enter some message and press enter and see the message in the terminal that net cat is running.
func main() {
	addr, err := net.ResolveUDPAddr("udp", "localhost:42069")
	if err != nil {
		log.Fatal(err)
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Println(err)
			continue
		}

		_, err = conn.Write([]byte(line))
		if err != nil {
			log.Println(err)
		}
	}
}
