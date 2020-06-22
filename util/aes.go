package util

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
)

// PadLength calculates padding length, from github.com/vgorin/cryptogo
func PadLength(sliceLength, blockSize int) (padLen int) {
	padLen = blockSize - sliceLength%blockSize
	if padLen == 0 {
		padLen = blockSize
	}
	return padLen
}

//from github.com/vgorin/cryptogo
func PKCS7Pad(message []byte, blockSize int) (padded []byte) {
	// block size must be bigger or equal 2
	if blockSize < 1<<1 {
		panic("block size is too small (minimum is 2 bytes)")
	}
	// block size up to 255 requires 1 byte padding
	if blockSize < 1<<8 {
		// calculate padding length
		padLen := PadLength(len(message), blockSize)

		// define PKCS7 padding block
		padding := bytes.Repeat([]byte{byte(padLen)}, padLen)

		// apply padding
		padded = append(message, padding...)
		return padded
	}
	// block size bigger or equal 256 is not currently supported
	panic("unsupported block size")
}

func AesEncrypt(plainData []byte, aesKey []byte) ([]byte, error) {
	k := len(aesKey)
	if len(plainData)%k != 0 {
		plainData = PKCS7Pad(plainData, k)
	}

	block, err := aes.NewCipher(aesKey)
	if err != nil {
		return nil, err
	}

	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	cipherData := make([]byte, len(plainData))
	blockMode := cipher.NewCBCEncrypter(block, iv)
	blockMode.CryptBlocks(cipherData, plainData)

	return cipherData, nil
}

func AesDecrypt(cipherData []byte, aesKey []byte) ([]byte, error) {
	k := len(aesKey) //PKCS#7
	if len(cipherData)%k != 0 {
		return nil, errors.New("crypto/cipher: ciphertext size is not multiple of aes key length")
	}

	block, err := aes.NewCipher(aesKey)
	if err != nil {
		return nil, err
	}

	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, iv)
	plainData := make([]byte, len(cipherData))
	blockMode.CryptBlocks(plainData, cipherData)
	return plainData, nil
}
