package main

import (
	"fmt"
	"golang-encoding/lib"
)

func main() {
	lib.EncodeToStringDecodeString()
	fmt.Println()
	lib.Encode()
	fmt.Println()
	lib.Decode()
	fmt.Println()
	lib.EncodeUrl()
	fmt.Println()
	lib.DecodeUrl()
}
