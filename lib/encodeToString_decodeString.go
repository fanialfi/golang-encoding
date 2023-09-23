package lib

import (
	"encoding/base64"
	"fmt"
)

func EncodeToStringDecodeString() {
	data := "Fani Alfirdaus"

	encodingString := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println("encoded :", encodingString)

	decodeByte, err := base64.StdEncoding.DecodeString(encodingString)
	if err != nil {
		fmt.Println("error", err.Error())
		return
	}
	decodeString := string(decodeByte)
	fmt.Printf("%T\t%v\n", decodeByte, decodeByte)
	fmt.Println("decoded :", decodeString)
}
