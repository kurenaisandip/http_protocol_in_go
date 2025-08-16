package main

import (
	"fmt"
	"os"
	_ "os"
)

// method to catch the error
func check(err error) {
	if err != nil {
		panic(err)
	}
}
func main() {

	//read file from the disk
	//dat, err := os.ReadFile("./message.txt") // this method also work but i cannot control how to parse the file
	//check(err)
	//fmt.Println(string(dat))

	//for more control i am going to use .open method
	file, err := os.Open("./message.txt")
	check(err)

	buffer := make([]byte, 8)

	for {
		n1, err := file.Read(buffer)
		if n1 > 0 {
			fmt.Printf("%d bytes: %s\n", n1, string(buffer[:n1]))
		}

		if err != nil {
			break
		}
	}

}
