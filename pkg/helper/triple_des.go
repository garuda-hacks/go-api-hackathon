package helper

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"encoding/base64"
)

var (
	sharedkey = []byte{0x01, 0x02, 0x03, 0x05, 0x07, 0x0B, 0x0D, 0x11, 0x12,
		0x11, 0x0D, 0x0B, 0x07, 0x02, 0x04, 0x08, 0x01, 0x02, 0x03, 0x05, 0x07, 0x0B, 0x0D, 0x11}
	sharedvector = []byte{0x01, 0x02, 0x03, 0x05, 0x07, 0x0B, 0x0D, 0x11}
)

func Encrypt(plain string) string {
	key := sharedkey
	iv := sharedvector

	block, _ := des.NewTripleDESCipher(key)
	input := []byte(plain)
	input = PKCS5Padding(input, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, iv)
	crypted := make([]byte, len(input))
	blockMode.CryptBlocks(crypted, input)

	return base64.StdEncoding.EncodeToString(crypted)
}

func Decrypt(secret string) (string, error) {
	key := sharedkey
	iv := sharedvector

	crypted, err := base64.StdEncoding.DecodeString(secret)
	if err != nil {
		return "", err
	}
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return "", err
	}
	blockMode := cipher.NewCBCDecrypter(block, iv)
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	return string(origData), nil
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	// remove the last byte unpadding times
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
