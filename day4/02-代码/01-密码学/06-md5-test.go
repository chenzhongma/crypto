package main

import (
	"crypto/md5"
	"io"
	"fmt"
)

//哈希运算，使用go包，有两种调用方式

//方式一

//16bytes, 128bit
func md5Test1(info []byte) []byte {
	//对多量数据进行哈希运算

	//1. 创建一个哈希器
	hasher := md5.New()

	io.WriteString(hasher, "hello ")
	io.WriteString(hasher, "world!")

	//2. 执行Sum操作，得到哈希值
	//hash := hasher.Sum(nil)
	//sum(b), 如果b不是nil， 那么返回的值为b+hash值
	hash := hasher.Sum([]byte("0x"))

	return hash
}

//方式二
func md5Test2(info []byte) []byte {
	hash := md5.Sum(info)

	//将数组转换为切片
	return hash[:]
}

func main() {

	hash := md5Test1(nil)

	fmt.Printf("hash : %x\n", hash)

	fmt.Printf("+++++++++++++++\n")

	src := []byte("hello world!")
	hash2 := md5Test2(src)

	fmt.Printf("hash2 : %x\n", hash2)
}
