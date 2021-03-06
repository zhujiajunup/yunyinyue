/*
Package encrypt provides encrypt algorithm such as rsa & aes
*/
package encrypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"math/big"
	"math/rand"
	"time"
)

// generate string for given size
func RandomStr(size int) (result []byte) {
	s := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	strBytes := []byte(s)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < size; i++ {
		result = append(result, strBytes[r.Intn(len(strBytes))])
	}
	return
}

func AesEncrypt(sSrc string, sKey string, aseKey string) (string, error) {
	iv := []byte(aseKey)
	block, err := aes.NewCipher([]byte(sKey))
	if err != nil {
		return "", err
	}
	padding := block.BlockSize() - len([]byte(sSrc))%block.BlockSize()
	src := append([]byte(sSrc), bytes.Repeat([]byte{byte(padding)}, padding)...)

	model := cipher.NewCBCEncrypter(block, iv)
	cipherText := make([]byte, len(src))
	model.CryptBlocks(cipherText, src)
	return base64.StdEncoding.EncodeToString(cipherText), nil
}

func RsaEncrypt(key string, pubKey string, modulus string) string {
	rKey := ""
	for i := len(key) - 1; i >= 0; i-- { // reserve key
		rKey += key[i : i+1]
	}
	hexRKey := ""
	for _, char := range []rune(rKey) {
		hexRKey += fmt.Sprintf("%x", int(char))
	}
	bigRKey, _ := big.NewInt(0).SetString(hexRKey, 16)
	bigPubKey, _ := big.NewInt(0).SetString(pubKey, 16)
	bigModulus, _ := big.NewInt(0).SetString(modulus, 16)
	bigRs := bigRKey.Exp(bigRKey, bigPubKey, bigModulus)
	hexRs := fmt.Sprintf("%x", bigRs)
	return addPadding(hexRs, modulus)
}

func addPadding(encText string, modulus string) string {
	ml := len(modulus)
	for i := 0; ml > 0 && modulus[i:i+1] == "0"; i++ {
		ml--
	}
	num := ml - len(encText)
	prefix := ""
	for i := 0; i < num; i++ {
		prefix += "0"
	}
	return prefix + encText
}
