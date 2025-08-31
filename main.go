package main

import (
	"fmt"
	"io"
	"net"
	_ "net"
	_ "os"
	"strings"
)

// method to catch the error
func check(err error) {
	if err != nil {
		panic(err)
	}
}

// Reads 8 bytes at a time, aggregates lines, and sends them to a channel
func getLinesChannel(f io.ReadCloser) <-chan string {
	ch := make(chan string)   // channel to return
	buffer := make([]byte, 8) // 8-byte buffer
	var currentLine string    // holds incomplete line across chunks

	go func() {
		defer close(ch) // close channel when done
		//defer f.Close() // close file when done

		for {
			n, err := f.Read(buffer)
			if n > 0 {
				chunk := string(buffer[:n])
				parts := strings.Split(chunk, "\n")

				for i := 0; i < len(parts)-1; i++ {
					// send full line to channel
					ch <- currentLine + parts[i]
					currentLine = ""
				}

				// last part may be incomplete, keep for next iteration
				currentLine += parts[len(parts)-1]
			}

			if err != nil {
				if err != io.EOF {
					fmt.Println("Error reading file:", err)
				}
				break
			}
		}

		// send any leftover line after EOF
		if len(currentLine) > 0 {
			ch <- currentLine
		}
	}()

	return ch
}

func main() {

	//read file from the disk --v1
	//dat, err := os.ReadFile("./message.txt") // this method also work but i cannot control how to parse the file
	//check(err)
	//fmt.Println(string(dat))

	//for more control i am going to use .open method  --v2
	//file, err := os.Open("./message.txt")
	//check(err)
	//
	//for line := range getLinesChannel(file) {
	//	fmt.Printf("read: %s\n", line)
	//}

	// this is for creating a server to read the file
	ln, err := net.Listen("tcp", ":42069")
	check(err)

	defer ln.Close()
	for {
		conn, err := ln.Accept()
		check(err)
		go handleConnection(conn)

	}

}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Println("new connection has been accepted:", conn.RemoteAddr())

	for line := range getLinesChannel(conn) { // read lines from the client
		fmt.Printf("received: %s\n", line)
	}
}
