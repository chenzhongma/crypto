package main

import (
	"crypto/rsa"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"os"
	"fmt"
)

//需求: 生成并保存私钥，公钥

func generateKeyPair(bits int) error {

	//生成私钥分析：
	//1. GenerateKey函数使用随机数据生成器random生成一对具有指定字位数的RSA密钥。
	//func GenerateKey(random io.Reader, bits int) (priv *PrivateKey, err error)
	//包： rsa
	//- 参数1：随机数, crypto/rand, 随机数生成器
	//- 参数2：秘钥长度
	//- 返回值：私钥
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)

	if err != nil {
		return err
	}

	//
	//2. 要对生成的私钥进行编码处理， x509， 按照规则，进行序列化处理, 生成der编码的数据
	//MarshalPKCS1PrivateKey将公钥序列化为PKCS格式DER编码。

	// MarshalPKCS1PrivateKey converts a private key to ASN.1 DER encoded form.

	//func MarshalPKCS1PrivateKey(key *rsa.PrivateKey) []byte {
	priDerText := x509.MarshalPKCS1PrivateKey(privateKey)

	//3. 创建Block代表PEM编码的结构, 并填入der编码的数据
	//type Block struct {
	//    Type    string            // 得自前言的类型（如"RSA PRIVATE KEY"）
	//    Headers map[string]string // 可选的头项
	//    Bytes   []byte            // 内容解码后的数据，一般是DER编码的ASN.1结构
	//}

	block := pem.Block{
		Type:    "SZ RSA PRIVATE KEY", //随便填写
		Headers: nil,                  //可选信息，包括私钥加密方式等
		Bytes:   priDerText,           //私钥编码后的数据
	}

	//4. 将Pem Block数据写入到磁盘文件
	fileHandler1, err := os.Create(PrivateKeyFile)
	if err != nil {
		return err
	}

	//关闭句柄
	defer fileHandler1.Close()

	//func Encode(out io.Writer, b *Block) error
	err = pem.Encode(fileHandler1, &block)

	if err != nil {
		return err
	}

	fmt.Println("++++++++++++++ 生成公钥 +++++++++++")

	/*
	1. 获取公钥， 通过私钥获取
	2. 要对生成的私钥进行编码处理， x509， 按照规则，进行序列化处理, 生成der编码的数据
	3. 创建Block代表PEM编码的结构, 并填入der编码的数据
	4. 将Pem Block数据写入到磁盘文件
	*/

	//1. 获取公钥， 通过私钥获取
	pubKey := privateKey.PublicKey //注意是对象，而不是地址

	//2. 要对生成的私钥进行编码处理， x509， 按照规则，进行序列化处理, 生成der编码的数据
	pubKeyDerText := x509.MarshalPKCS1PublicKey(&pubKey)

	//3. 创建Block代表PEM编码的结构, 并填入der编码的数据
	block1 := pem.Block{
		Type:    "SZ RSA Public Key",
		Headers: nil,
		Bytes:   pubKeyDerText,
	}

	//4. 将Pem Block数据写入到磁盘文件

	fileHandler2, err := os.Create(PublicKeyFile)
	if err != nil {
		return err
	}

	//关闭句柄
	defer fileHandler2.Close()

	err = pem.Encode(fileHandler2, &block1)
	if err != nil {
		return err
	}

	return nil
}

func main() {

	fmt.Printf("generate rsa private key ...\n")
	err := generateKeyPair(1024)
	if err != nil {
		fmt.Printf("generate rsa private failed, err : %v", err)
	}

	fmt.Printf("generate rsa private key successfully!\n")
}
