package main

import (
	"io/ioutil"
	"encoding/pem"
	"crypto/x509"
	"crypto/rsa"
	"crypto/rand"
	"fmt"
)

func rsaPubEncrypt(filename string, plainText []byte) (error, []byte) {
	//1. 通过公钥文件，读取公钥信息 ==》 pem encode 的数据
	info, err := ioutil.ReadFile(filename)

	if err != nil {
		return err, nil
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

	//3. 解码der，得到公钥
	//derText := block.Bytes
	derText := block.Bytes
	publicKey, err := x509.ParsePKCS1PublicKey(derText)

	if err != nil {
		return err, nil
	}

	//4. 公钥加密
	//EncryptPKCS1v15使用PKCS#1 v1.5规定的填充方案和RSA算法加密msg。
	//func EncryptPKCS1v15(rand io.Reader, pub *PublicKey, msg []byte) (out []byte, err error)

	cipherData, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, plainText)

	if err != nil {
		return err, nil
	}

	return nil, cipherData
}

func rsaPriKeyDecrypt(filename string, cipherData []byte) (error, []byte) {
	//1. 通过私钥文件，读取私钥信息 ==》 pem encode 的数据
	info, err := ioutil.ReadFile(filename)

	if err != nil {
		return err, nil
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
		return err, nil
	}

	//4. 私钥解密
	//DecryptPKCS1v15使用PKCS#1 v1.5规定的填充方案和RSA算法解密密文。如果random不是nil，函数会注意规避时间侧信道攻击。
	//func DecryptPKCS1v15(rand io.Reader, priv *PrivateKey, ciphertext []byte) (out []byte, err error)
	plainText, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, cipherData)

	if err != nil {
		return err, nil
	}

	return nil, plainText
}

func main() {
	src := []byte("祝班主任节日快乐!祝班主任节日快乐祝班主任节日快乐祝班主任节日快乐祝班主任节日快乐祝班主任节日快乐祝班主任节日快乐祝班主任节日快乐祝班主任节日快乐祝班主任节日快乐祝班主任节日快乐祝班主任节日快乐祝班主任节日快乐祝班主任节日快乐祝班主任节日快乐祝班主任节日快乐祝班主任节日快乐祝班主任节日快乐祝班主任节日快乐祝班主任节日快乐")
	err, cipherData := rsaPubEncrypt(PublicKeyFile, src)

	if err != nil {
		fmt.Println("公钥加密失败!, err :", err)
	}

	fmt.Printf("cipherData : %x\n", cipherData)
	fmt.Println("++++++++++++++++++++++++++++++")

	err, plainText := rsaPriKeyDecrypt(PrivateKeyFile, cipherData)
	if err != nil {
		fmt.Println("私钥解密失败!, err : ", err)
	}
	fmt.Printf("plainText : %s\n", plainText)
}
