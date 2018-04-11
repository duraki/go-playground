package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	/** heredoc syntax **/
	data := `
		This is rather a long list of strings or so to say letters
		combined.
`

	/** encryption **/
	sEnc := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)

	/** decryption **/
	sDec, _ := base64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec))
	fmt.Println()

	/** url encoding & decoding **/
	uEnc := base64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)

	uDec, _ := base64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))
}
