package lib

import (
	"encoding/base64"
	"fmt"
)

func Encode() {
	src := []byte("Fani Alfirdaus")
	dst := make([]byte, base64.StdEncoding.EncodedLen(len(src)))

	base64.StdEncoding.Encode(dst, src)
	fmt.Println(src, string(src))
	fmt.Println(dst, string(dst))
}

func Decode() {
	dataPlan := []byte("Super Hero")
	dataEncode := make([]byte, base64.StdEncoding.EncodedLen(len(dataPlan)))

	base64.StdEncoding.Encode(dataEncode, dataPlan)

	// decoding data ke bentuk semula
	src := dataEncode
	dst := make([]byte, base64.StdEncoding.DecodedLen(len(src)))
	n, err := base64.StdEncoding.Decode(dst, src)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("n, is", n)
	}

	fmt.Println(src, string(src))
	fmt.Println(dst, string(dst))
}
