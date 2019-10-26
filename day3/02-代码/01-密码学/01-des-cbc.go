package main

import (
	"fmt"
	"crypto/des"
	"crypto/cipher"
	"bytes"
)

/*
需求：算法：des ， 分组模式：CBC

des :
秘钥：8bytes
分组长度：8bytes

cbc:
1. 提供初始化向量，长度与分组长度相同，8bytes
2. 需要填充


加密分析

1. 创建并返回一个使用DES算法的cipher.Block接口。

	func NewCipher(key []byte) (cipher.Block, error)
	- 包名：des
	- 参数：秘钥，8bytes
	- 返回值：一个cipher.Block接口

	type Block interface {
		// 返回加密字节块的大小
		BlockSize() int
		// 加密src的第一块数据并写入dst，src和dst可指向同一内存地址
		Encrypt(dst, src []byte)
		// 解密src的第一块数据并写入dst，src和dst可指向同一内存地址
		Decrypt(dst, src []byte)
	}

2. 进行数据填充
//TODO


3. 引入CBC模式, 返回一个密码分组链接模式的、底层用b加密的BlockMode接口，初始向量iv的长度必须等于b的块尺寸。
	func NewCBCEncrypter(b Block, iv []byte) BlockMode
	- 包名：cipher
	- 参数1：cipher.Block
	- 参数2：iv， initialize vector
	- 返回值：分组模式，里面提供加解密方法

	type BlockMode interface {
		// 返回加密字节块的大小
		BlockSize() int
		// 加密或解密连续的数据块，src的尺寸必须是块大小的整数倍，src和dst可指向同一内存地址
		CryptBlocks(dst, src []byte)
	}

解密分析
1. 创建并返回一个使用DES算法的cipher.Block接口。

	func NewCipher(key []byte) (cipher.Block, error)
	- 包名：des
	- 参数：秘钥，8bytes
	- 返回值：一个cipher.Block接口

	type Block interface {
		// 返回加密字节块的大小
		BlockSize() int
		// 加密src的第一块数据并写入dst，src和dst可指向同一内存地址
		Encrypt(dst, src []byte)
		// 解密src的第一块数据并写入dst，src和dst可指向同一内存地址
		Decrypt(dst, src []byte)
	}


2. 返回一个密码分组链接模式的、底层用b解密的BlockMode接口，初始向量iv必须和加密时使用的iv相同。
	func NewCBCDecrypter(b Block, iv []byte) BlockMode
	- 包名：cipher
	- 参数1：cipher.Block
	- 参数2：iv， initialize vector
	- 返回值：分组模式，里面提供加解密方法

	type BlockMode interface {
		// 返回加密字节块的大小
		BlockSize() int
		// 加密或解密连续的数据块，src的尺寸必须是块大小的整数倍，src和dst可指向同一内存地址
		CryptBlocks(dst, src []byte)
	}

3. 解密操作

4. 去除填充
//TODO
*/

//输入明文，秘钥，输出密文
func desCBCEncrypt(src, key []byte) []byte {
	fmt.Printf("加密开始，输入的数据为：%s\n", src)

	//1. 创建并返回一个使用DES算法的cipher.Block接口。
	//NewCipher(key []byte) (cipher.Block, error)
	block, err := des.NewCipher(key)

	fmt.Printf("block size : %d\n", block.BlockSize())

	if err != nil {
		panic(err)
	}

	//2. 进行数据填充
	src = paddingInfo(src, block.BlockSize())

	//3. 引入CBC模式, 返回一个密码分组链接模式的、底层用b加密的BlockMode接口，初始向量iv的长度必须等于b的块尺寸。
	//func NewCBCEncrypter(b Block, iv []byte) BlockMode

	iv := bytes.Repeat([]byte("1"), block.BlockSize())

	blockMode := cipher.NewCBCEncrypter(block, iv)

	//4. 加密操作
	blockMode.CryptBlocks(src /*加密后的密文*/ , src /*明文*/)

	fmt.Printf("加密结束，加密数据为%x\n", src)
	return src
}

//填充函数, 输入明文, 分组长度, 输出：填充后的数据
func paddingInfo(src []byte, blockSize int) []byte {
	//1. 得到明文的长度
	length := len(src)

	//2. 需要填充的数量

	remains := length % blockSize        //3
	paddingNumber := blockSize - remains //5

	//3. 把填充的数值转换为字符
	s1 := byte(paddingNumber) // '5'

	fmt.Printf("s1 : %x\n", s1)

	//4. 把字符拼成数组
	s2 := bytes.Repeat([]byte{s1}, paddingNumber) //[]byte{0x35, '5', '5', '5, '5'}

	fmt.Printf("s2 : %x\n", s2)
	//5. 把拼成的数组追加到src后面
	srcNew := append(src, s2...)

	//6. 返回新的数组
	return srcNew
}

//输入密文，秘钥，得到明文
func desCBCDecrypt(cipherData, key []byte) []byte {
	fmt.Printf("解密开始，输入的数据为：%x\n", cipherData)

	//1. 创建并返回一个使用DES算法的cipher.Block接口。
	//NewCipher(key []byte) (cipher.Block, error)
	block, err := des.NewCipher(key)

	fmt.Printf("block size : %d\n", block.BlockSize())

	if err != nil {
		panic(err)
	}

	//3. 引入CBC模式

	iv := bytes.Repeat([]byte("1"), block.BlockSize())

	blockMode := cipher.NewCBCDecrypter(block, iv)

	//4. 解密操作
	blockMode.CryptBlocks(cipherData /*解密后的明文*/ , cipherData /*密文*/)

	fmt.Printf("解密结束，解密数据为%s\n", cipherData)

	//5. 去除填充
	plainText := unpaddingInfo(cipherData /*解密后的数据*/)

	return plainText
}

//去除填充
func unpaddingInfo(plainText []byte) []byte {
	//1. 获取长度
	length := len(plainText)

	if length == 0 {
		return []byte{}
	}

	//2. 获取最后一个字符
	lastByte := plainText[length-1]

	//3. 将字符转换成数字
	unpaddingNumber := int(lastByte)

	//4. 切片获取需要的数据
	return plainText[:length-unpaddingNumber]

}

func main() {
	src := []byte("不是一番寒彻骨，哪得梅花扑鼻香!!!")
	key := []byte("12345678")

	cipherData := desCBCEncrypt(src, key)

	fmt.Printf("cipherData : %x\n", cipherData)
	fmt.Printf("+++++++++++++++++++++++++\n")

	plainText := desCBCDecrypt(cipherData, key)
	fmt.Printf("plainText str: %s\n", plainText)
	fmt.Printf("plainText hex: %x\n", plainText)
}
