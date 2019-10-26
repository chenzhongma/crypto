package main

import (
	"os"
	"crypto/sha256"
	"io"
	"fmt"
)

//使用打开文件方式获取哈希

const filename = "/Users/duke/Desktop/给学员的区块链工具包/03_以太坊相关/03-钱包/Ethereum-Wallet-installer-0-9-3.exe"

func main() {

	//1. open 文件
	file, err := os.Open(filename)

	defer file.Close()

	if err != nil {
		panic(err)
	}

	//2. 创建hash
	hasher := sha256.New()

	/*
	type Hash interface {
		io.Writer
		Sum(b []byte) []byte

		Reset()

		Size() int
		BlockSize() int
	}
	*/

	//3. copy句柄
	//func Copy(dst Writer, src Reader) (written int64, err error) {
	length, err := io.Copy(hasher, file)

	if err != nil {
		panic(err)
	}

	fmt.Printf("length : %d\n", length)

	//4. hash sum操作
	hash := hasher.Sum(nil)

	fmt.Printf("hash : %x\n", hash)

}
