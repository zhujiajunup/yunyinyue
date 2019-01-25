package encrypt

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func TestRandom(t *testing.T) {
	fmt.Println(len(Random(16)))
}

func TestEncryptData(t *testing.T) {
	params := make(map[string]string)
	params["rid"] = "R_SO_4_376635"
	params["offset"] = "0"
	params["totail"] = "true"
	params["limit"] = "20"
	params["csrf_token"] = ""
	result := EncryptData(params)
	fmt.Println(result)
}

func TestAesEncryptAndDecrypt(t *testing.T) {
	aeskey := "321423u9y8d2fwfl"
	pass := "zczdfff"
	xpass, err := aesEncrypt(pass, aeskey)
	if err != nil {
		fmt.Println(err)
		return
	}

	pass64 := base64.StdEncoding.EncodeToString([]byte(xpass))
	fmt.Printf("加密后:%v\n",pass64)

	bytesPass, err := base64.StdEncoding.DecodeString(pass64)
	if err != nil {
		fmt.Println(err)
		return
	}

	tpass, err := AesDecrypt(bytesPass, []byte(aeskey))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("解密后:%s\n", tpass)
}
