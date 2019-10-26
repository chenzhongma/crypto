package main

import (
	"crypto/aes"
	"crypto/cipher"
	"bytes"
	"fmt"
)

/*
需求： 使用aes， ctr

aes :
- 分组长度： 16
- 秘钥：16

ctr:
- 不需要填充
- 需要提供一个数字


1. 创建一个cipher.Block接口。参数key为密钥，长度只能是16、24、32字节，用以选择AES-128、AES-192、AES-256。
func NewCipher(key []byte) (cipher.Block, error)
- 包：aes
- 秘钥
- cipher.Block接口


2. 选择分组模式：ctr
返回一个计数器模式的、底层采用block生成key流的Stream接口，初始向量iv的长度必须等于block的块尺寸。
func NewCTR(block Block, iv []byte) Stream
- block
- iv
- 秘钥流

3. 加密操作
type Stream interface {
    // 从加密器的key流和src中依次取出字节二者xor后写入dst，src和dst可指向同一内存地址
    XORKeyStream(dst, src []byte)
}

*/

func aesCTREncrypt(src, key []byte) []byte {
	fmt.Printf("明文： %s\n", src)

	//1. 创建一个cipher.Block接口。
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	fmt.Println("aes block size : ", block.BlockSize())

	iv := bytes.Repeat([]byte("1"), block.BlockSize())

	//2. 选择分组模式：ctr
	stream := cipher.NewCTR(block, iv)

	//3. 加密操作
	stream.XORKeyStream(src /*密文*/ , src /*明文*/)

	return src
}

func aesCTRDecrypt(cipherData, key []byte) []byte {

	//1. 创建一个cipher.Block接口。
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	iv := bytes.Repeat([]byte("1"), block.BlockSize())

	//2. 选择分组模式：ctr
	stream := cipher.NewCTR(block, iv)

	//3. 解密操作
	stream.XORKeyStream(cipherData /*明文*/ , cipherData)

	return cipherData
}

func main() {

	src := []byte("不是一番寒彻骨，哪得梅花扑鼻香!!! 123456734523452345	")
	key := []byte("1234567887654321")

	cipherData := aesCTREncrypt(src, key)

	fmt.Printf("cipherData : %x\n", cipherData)

	fmt.Printf("+++++++++++++++++++++++++\n")

	plainText := aesCTRDecrypt(cipherData, key)
	fmt.Printf("plainText ： %s\n", plainText)
}
