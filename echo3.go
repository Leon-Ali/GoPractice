package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	str := strings.Join(os.Args[1:], " ")
	splitStr := strings.Split(os.Args[0], "/")
	fmt.Println(splitStr[len(splitStr)-1])
	fmt.Println(str)
}