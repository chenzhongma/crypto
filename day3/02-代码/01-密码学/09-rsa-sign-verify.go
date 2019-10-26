package main

import (
	"io/ioutil"
	"encoding/pem"
	"crypto/x509"
	"crypto/sha256"
	"crypto/rsa"
	"crypto/rand"
	"crypto"
	"fmt"
)

/*
私钥签名:
1. 提供私钥文件， 解析出私钥内容（decode, parse....）

2. 使用私钥进行数字签名


公钥认证
1. 提供公钥文件， 解析出公钥内容（decode, parse....）

2. 使用公钥进行数字签名认证



*/

//私钥签名: 提供私钥，签名数据，得到数字签名
func rsaSignData(filename string, src []byte) ([]byte, error) {

	//一、 提供私钥文件， 解析出私钥内容（decode, parse....）
	//1. 通过私钥文件，读取私钥信息 ==》 pem encode 的数据
	info, err := ioutil.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	//2. pem decode， 得到block中的der编码数据
	block, _ := pem.Decode(info)
	//返回值1 ：pem.block
	//返回值2：rest参加是未解码完的数据，存储在这里

	//type Block struct {
	//    Type    string            // 得自前言的类型（如"RSA PRIVATE KEY"）
	//    Headers map[string]string // 可选的头项
	//    Bytes   []byte            // 内容解码后的数据，一般是DER编码的ASN.1结构
	//}

	//3. 解码der，得到私钥
	//derText := block.Bytes
	derText := block.Bytes
	privateKey, err := x509.ParsePKCS1PrivateKey(derText)

	if err != nil {
		return nil, err
	}

	//二. 使用私钥进行数字签名

	//1. 获取原文的哈希值
	hash := sha256.Sum256(src) //返回值是[32]byte， 一个数组

	//SignPKCS1v15使用RSA PKCS#1 v1.5规定的RSASSA-PKCS1-V1_5-SIGN签名方案计算签名
	//func SignPKCS1v15(rand io.Reader, priv *PrivateKey, hash crypto.Hash, hashed []byte) (s []byte, err error)

	//2. 执行签名操作
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hash[:])
	if err != nil {
		return nil, err
	}

	return signature, nil
}

//公钥认证
func rsaVerifySignature(sig []byte, src []byte, filename string) error {
	//一. 提供公钥文件， 解析出公钥内容（decode, parse....）
	//1. 通过公钥文件，读取公钥信息 ==》 pem encode 的数据
	info, err := ioutil.ReadFile(filename)

	if err != nil {
		return err
	}

	//2. pem decode， 得到block中的der编码数据
	block, _ := pem.Decode(info)

	//3. 解码der，得到公钥
	//derText := block.Bytes
	derText := block.Bytes
	publicKey, err := x509.ParsePKCS1PublicKey(derText)

	if err != nil {
		return err
	}

	//二. 使用公钥进行数字签名认证

	//1. 获取原文的哈希值
	hash := sha256.Sum256(src) //返回值是[32]byte， 一个数组

	//VerifyPKCS1v15认证RSA PKCS#1 v1.5签名。hashed是使用提供的hash参数对（要签名的）原始数据进行hash的结果。合法的签名会返回nil，否则表示签名不合法。
	//func VerifyPKCS1v15(pub *PublicKey, hash crypto.Hash, hashed []byte, sig []byte) error {
	err = rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hash[:], sig)

	return err
}

func main() {

	src := []byte("hello world!!!!")
	signature, err := rsaSignData(PrivateKeyFile, src)
	if err != nil {
		fmt.Printf("签名失败!, err: %s\n", err)
	}

	fmt.Printf("signature ： %x\n", signature)
	//fmt.Printf("signature ： %s\n", signature)

	fmt.Printf("++++++++++\n")

	src1 := []byte("hello world!!!!=======")

	err = rsaVerifySignature(signature, src1, PublicKeyFile)
	if err != nil {
		fmt.Printf("签名校验失败!, err: %s\n", err)
		return
	}

	fmt.Printf("签名校验成功!\n")
}
