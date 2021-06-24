package cipher

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/hex"
	"errors"
)

func EncryptAES(key, value string) (string, error) {

	if key == "" {
		return "", errors.New("invalid key given")
	}

	if value == "" {
		return "", errors.New("invalid value given")
	}

	bKey, err := hex.DecodeString(key)
	if err != nil {
		return "", err
	}

	bValue := []byte(value)
	bPlaintext := AddPkcs7(bValue, aes.BlockSize)

	if len(bPlaintext)%aes.BlockSize != 0 {
		return "", errors.New("text length is not a multiple of the block size")
	}

	block, err := aes.NewCipher(bKey)
	if err != nil {
		return "", err
	}

	ciphertext := make([]byte, len(bPlaintext))
	bIv := ciphertext[:aes.BlockSize]

	mode := cipher.NewCBCEncrypter(block, bIv)
	mode.CryptBlocks(ciphertext, bPlaintext)
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func DecryptAES(key, value string) (string, error) {
	if key == "" {
		return "", errors.New("invalid key given")
	}

	if value == "" {
		return "", errors.New("invalid value given")
	}

	bKey, err := hex.DecodeString(key)
	if err != nil {
		return "", err
	}

	bValue, err := base64.StdEncoding.DecodeString(value)
	if err != nil {
		return "", err
	}

	bPlaintext := RemovePkcs7(bValue, aes.BlockSize)

	if len(bPlaintext)%aes.BlockSize != 0 {
		return "", errors.New("text length is not a multiple of the block size")
	}

	block, err := aes.NewCipher(bKey)
	if err != nil {
		return "", err
	}

	ciphertext := make([]byte, len(bPlaintext))
	bIv := ciphertext[:aes.BlockSize]

	mode := cipher.NewCBCDecrypter(block, bIv)
	mode.CryptBlocks(ciphertext, bPlaintext)

	return string(ciphertext), nil
}

// AddPkcs7 pads given byte array using pkcs7 padding schema till it has blockSize length in bytes
func AddPkcs7(data []byte, blockSize int) []byte {

	var paddingCount int

	if paddingCount = blockSize - (len(data) % blockSize); paddingCount == 0 {
		paddingCount = blockSize
	}

	return append(data, bytes.Repeat([]byte{byte(paddingCount)}, paddingCount)...)
}

// RemovePkcs7 removes pkcs7 padding from previously padded byte array
func RemovePkcs7(padded []byte, blockSize int) []byte {

	dataLen := len(padded)
	paddingCount := int(padded[dataLen-1])

	if paddingCount > blockSize || paddingCount <= 0 {
		return padded //data is not padded (or not padded correctly), return as is
	}

	padding := padded[dataLen-paddingCount : dataLen-1]

	for _, b := range padding {
		if int(b) != paddingCount {
			return padded //data is not padded (or not padded correcly), return as is
		}
	}

	return padded[:len(padded)-paddingCount] //return data - padding
}
