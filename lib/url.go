package lib

import (
	"encoding/base64"
	"fmt"
)

func EncodeUrl() {
	url := "https://www.fanialfi.space?name=fani&email=faialfirdaus@fanialfi.space"

	encodeStr := base64.URLEncoding.EncodeToString([]byte(url))
	fmt.Println(encodeStr)

	dst := make([]byte, base64.URLEncoding.EncodedLen(len([]byte(url))))
	base64.URLEncoding.Encode(dst, []byte(url))
	fmt.Println(string(dst))

	// jika tidak menggunakan url encoding
	encodeStrNotWithurl := base64.StdEncoding.EncodeToString([]byte(url))
	fmt.Printf("\nwithout URLEncoding %s\n", encodeStrNotWithurl)
	decodeByte, err := base64.StdEncoding.DecodeString(encodeStrNotWithurl)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("\nwithout URLEncoding %s\n", string(decodeByte))
}

func DecodeUrl() {
	url := "https://www.fanialfi.space?name=fani&email=faialfirdaus@fanialfi.space"
	encodeStr := base64.URLEncoding.EncodeToString([]byte(url))

	decodeByte, err := base64.URLEncoding.DecodeString(encodeStr)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(decodeByte))

	// jika tidak menggunakan url encoding
	encodeStrNotUrl := base64.StdEncoding.EncodeToString([]byte(url))
	decodeByteNoturl, err := base64.StdEncoding.DecodeString(encodeStrNotUrl)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(decodeByteNoturl))
}
