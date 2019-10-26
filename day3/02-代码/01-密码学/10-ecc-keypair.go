package main

import (
	"crypto/elliptic"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"os"
	"fmt"
)

//生成私钥公钥

func generateEccKeypair() {
	//- 选择一个椭圆曲线（在elliptic包）
	//type Curve
	//func P224() Curve
	//func P256() Curve
	//func P384() Curve
	//func P521() Curve
	curve := elliptic.P521()

	//- 使用ecdsa包，创建私钥 //ecdsa椭圆曲线数字签名
	//GenerateKey函数生成秘钥对
	//func GenerateKey(c elliptic.Curve, rand io.Reader) (priv *PrivateKey, err error)
	privateKey, err := ecdsa.GenerateKey(curve, rand.Reader)

	checkErr("generate key failed!", err)

	//- 使用x509进行编码
	//MarshalECPrivateKey将ecdsa私钥序列化为ASN.1 DER编码。
	//func MarshalECPrivateKey(key *ecdsa.PrivateKey) ([]byte, error)
	derText, err := x509.MarshalECPrivateKey(privateKey)
	checkErr("MarshalECPrivateKey", err)

	//- 写入pem.Block中
	block1 := pem.Block{
		Type:    "ECC PRIVATE KEY",
		Headers: nil,
		Bytes:   derText,
	}

	//- pem.Encode
	fileHander, err := os.Create(EccPrivateKeyFile)
	checkErr("os.Create Failed", err)

	defer fileHander.Close()

	err = pem.Encode(fileHander, &block1)
	checkErr("pem Encode failed", err)

	fmt.Printf("++++++++++++++++++++++\n")
	//获取公钥
	publicKey := privateKey.PublicKey

	//- 使用x509进行编码
	//通用的序列化方式
	//derText2, err := x509.MarshalPKIXPublicKey(publicKey)
	derText2, err := x509.MarshalPKIXPublicKey(&publicKey)
	//传递地址

	checkErr("MarshalPKIXPublicKey", err)

	//- 写入pem.Block中
	block2 := pem.Block{
		Type:    "ECC PUBLICK KEY",
		Headers: nil,
		Bytes:   derText2,
	}

	//- pem.Encode
	fileHander2, err := os.Create(EccPublicKeyFile)
	checkErr("public key os.Create Failed", err)

	defer fileHander2.Close()

	err = pem.Encode(fileHander2, &block2)
	checkErr("public key pem Encode failed", err)

}

func main() {
	generateEccKeypair()
}
