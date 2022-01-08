package main

import (
	"bufio"
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)


func Encode(input string, key string, output string) {
	file, err := os.Open(input)
	if err != nil {
		panic(err)
	}

	content, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	k := []byte(key)
	if len(k) < 32 {
		panic(errors.New("key needs at least 32 bytes"))
	}
	k = k[:32]

	encryptCode := aesEncrypt(content, k)

	out := os.Stdout
	if output != "" {
		out, err = os.Create(output)
		if err != nil {
			panic(err)
		}
	}
	_, _ = fmt.Fprintln(out, encryptCode)
}

func Decode(input string, key string, output string) {
	file, err := os.Open(input)
	if err != nil {
		panic(err)
	}

	buf := bufio.NewReader(file)
	content, err := buf.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}

	k := []byte(key)
	if len(k) < 32 {
		panic(errors.New("key needs at least 32 bytes"))
	}
	k = k[:32]

	decryptCode := aesDecrypt(content, k)

	out := os.Stdout
	if output != "" {
		out, err = os.Create(output)
		if err != nil {
			panic(err)
		}
	}

	_, _ = fmt.Fprintln(out, decryptCode)
}

func aesEncrypt(origData []byte, k []byte) string {
	// 转成字节数组
	// origData := []byte(orig)
	// k := []byte(key)

	// 分组秘钥
	block, _ := aes.NewCipher(k)
	// 获取秘钥块的长度
	blockSize := block.BlockSize()
	// 补全码
	origData = padding(origData, blockSize)
	// 加密模式
	blockMode := cipher.NewCBCEncrypter(block, k[:blockSize])
	// 创建数组
	cryted := make([]byte, len(origData))
	// 加密
	blockMode.CryptBlocks(cryted, origData)

	return base64.StdEncoding.EncodeToString(cryted)

}

func aesDecrypt(cryted string, k []byte) string {
	// 转成字节数组
	crytedByte, _ := base64.StdEncoding.DecodeString(cryted)

	// 分组秘钥
	block, _ := aes.NewCipher(k)
	// 获取秘钥块的长度
	blockSize := block.BlockSize()
	// 加密模式
	blockMode := cipher.NewCBCDecrypter(block, k[:blockSize])
	// 创建数组
	orig := make([]byte, len(crytedByte))
	// 解密
	blockMode.CryptBlocks(orig, crytedByte)
	// 去补全码
	orig = unPadding(orig)
	return string(orig)
}

//补码
func padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padText...)
}

//去码
func unPadding(origData []byte) []byte {
	length := len(origData)
	unPad := int(origData[length-1])
	return origData[:(length - unPad)]
}
