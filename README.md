# encoding

go memiliki package `encoding/base64` berisi function function untuk keperluan **encode** dan **decode** data ke base64 dan sebaliknya, dimana data yang akan di encode harus bertipe `[]byte`.

## penerapan function `*Encoding.EncodeToString(src []byte) string` dan `*Encoding.DecodeString(s string) ([]byte, error)`

function `*Encoding.EncodeToString(src []byte)` digunakan untuk encoding data dari bentuk `[]byte` kedalam _base64_ dari _src_, dimana nanti bentuk dari _base64_ adalah `string`.

```go
package main

import (
  "encoding/base64"
  "fmt"
)

func main(){
  name := "Fani Alfirdaus"
  encodingString := base64.StdEncoding.EncodeToString([]byte(name))
  fmt.Println("encoded :",encodingString)
}
```

sedangkan function `*Encoding.DecodeString(s string) ([]byte, error)` merupakan kebalikan dari `EncodeToString()`, dimana function tersebut mengubah dari _base64 `string`_ ke bentuk data asli-nya (`[]byte`), function ini juga mengembalikan error jika proses decoding-nya gagal.

```go
package main

import (
  "encoding/base64"
  "fmt"
)

func main(){
  name := []byte("Fani Alfirdaus")
  encodingString := base64.StdEncoding.EncodeToString(name)
  
  decodeByte, err := base64.StdEncoding.DecodeString(encodingString)
  if err != nil {
    fmt.Println("error",err.Error())
    return
  }

  fmt.Println(string(decodeByte))
}
```

## penerapan function `*Encoding.Encode(dst, src []byte)` dan `*Encoding.Decode(dst, src []byte) (n int, err error)`

kedua function ini kegunaanya sama seperti kedua function sebelumnya (`*.Encoding.EncodeToString(src []byte) string`, dan `*.Encoding.DecodeString(s string) ([]byte, error)`). 
Salah satu pembedan dari function sebelumnya adalah data yang akan di decode maupun di encode keduanya bertipe `[]byte`.

Penggunaan cara ini lumayan panjang, karena variabel hasil dari encode maupun decode harus di siapkan terlebih dahulu, dan harus memiliki lebar data yang sesuai dengan hasil yang akan di tampung ( nilainya bisa dicari dengan menggunakan function `EncodedLen()` dan `DecodedLen()` )

```go
package main

import (
  "encoding/base64"
  "fmt"
)

func main(){
  Encode()
  Decode()
}

func Encode(){
  src := []byte("Fani Alfirdaus")
	dst := make([]byte, base64.StdEncoding.EncodedLen(len(src)))

	base64.StdEncoding.Encode(dst, src)
	fmt.Println(src, string(src))
	fmt.Println(dst, string(dst))
}

func Decode(){
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
```

## Encode dan Decode data URL

khusus data string yang isinya berupa url, lebih efektif menggunakan `URLEncoding` dibanding `StdEncoding`. 
Cara penerapannya pun sama seperti diatas.

```go
package main

import (
  "encoding/base64"
  "fmt"
)

func main(){
  EncodeUrl()
  DecodeUrl()
}

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
```