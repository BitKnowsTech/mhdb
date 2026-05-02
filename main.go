package main

import (
	"fmt"
	"os"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		panic("Could not get working diectory")
	}
	fmt.Println(wd)
}
