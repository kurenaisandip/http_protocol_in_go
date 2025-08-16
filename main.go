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
	dat, err := os.ReadFile("./message.txt")
	check(err)
	fmt.Println(string(dat))
}
