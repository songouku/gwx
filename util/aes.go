package util

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
)

//@brief: 填充明文
func PKCS5Padding(plaintext []byte, blockSize int) []byte {
	padding := blockSize - len(plaintext)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(plaintext, padText...)
}

//@brief: 去除填充数据
func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unPadding := int(origData[length-1])
	return origData[:(length - unPadding)]
}

//@brief: AES加密
func AesEncrypt(origData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	//AES分组长度为 128 位，所以 blockSize=16 字节
	blockSize := block.BlockSize()
	origData = PKCS5Padding(origData, blockSize)
	//初始向量的长度必须等于块block的长度16字节
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crapped := make([]byte, len(origData))
	blockMode.CryptBlocks(crapped, origData)
	return crapped, nil
}

//@brief:AES解密
func AesDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// AES分组长度为 128 位，所以 blockSize=16 字节
	blockSize := block.BlockSize()
	//初始向量的长度必须等于块block的长度16字节
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	return origData, nil
}
