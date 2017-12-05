package main

import (
	"fmt"
	"os"
)



func main() {
	fmt.Println("Test")
	fmt.Println(os.Args[0])
	skip := false
	if len(os.Args) > 1 && os.Args[1] == "skip" {
		fmt.Printf("%T %+v\n", os.Args[1],os.Args[1])
		skip = true
	}

	fmt.Printf("%T %+v\n", skip, skip)


}
