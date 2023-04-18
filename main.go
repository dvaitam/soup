package main

import (
	"fmt"
	"os"
)

func main() {
	path := ""
	fmt.Println("calculated ", path)
	b, err := os.ReadFile(path) // just pass the file name
	if err != nil {
		fmt.Print(err)
	}
	if err != nil {
		fmt.Println()
	}

	//fmt.Println(b) // print the content as 'bytes'

	s := string(b) // convert content to a 'string'
	tags := soup.load(s)
	fmt.Println(tags)
}
