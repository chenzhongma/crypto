package main

import (
	"io/ioutil"
	"encoding/pem"
	"crypto/x509"
	"crypto/sha256"
	"crypto/ecdsa"
	"crypto/rand"
	"math/big"
	"fmt"
	"github.com/pkg/errors"
)

//自己定义的签名结构
type Signature struct {
	r *big.Int
	s *big.Int
}

//使用私钥签名

func eccSignData(filename string, src []byte) (Signature, error) {
	//1. 读取私钥，解码
	info, err := ioutil.ReadFile(filename)

	if err != nil {
		return Signature{}, err
	}

	//2. pem decode， 得到block中的der编码数据
	block, _ := pem.Decode(info)
	//返回值1 ：pem.block
	//返回值2：rest参加是未解码完的数据，存储在这里

	//3. 解码der，得到私钥
	//derText := block.Bytes
	derText := block.Bytes
	privateKey, err := x509.ParseECPrivateKey(derText)

	if err != nil {
		return Signature{}, err
	}

	//2. 对原文生成哈希
	hash /*[32]byte*/ := sha256.Sum256(src)

	//3. 使用私钥签名
	//使用私钥对任意长度的hash值（必须是较大信息的hash结果）进行签名，返回签名结果（一对大整数）。私钥的安全性取决于密码读取器
	//func Sign(rand io.Reader, priv *PrivateKey, hash []byte) (r, s *big.Int, err error)
	r, s, err := ecdsa.Sign(rand.Reader, privateKey, hash[:])

	if err != nil {
		return Signature{}, err
	}

	sig := Signature{r, s}

	return sig, nil
}

//使用公钥校验
func eccVerifySig(filename string, src []byte, sig Signature) error {

	//1. 读取公钥，解码
	info, err := ioutil.ReadFile(filename)

	if err != nil {
		return err
	}

	//2. pem decode， 得到block中的der编码数据
	block, _ := pem.Decode(info)

	//3. 解码der，得到公钥
	//derText := block.Bytes
	derText := block.Bytes
	//publicKey, err := x509.ParsePKCS1PublicKey(derText)

	publicKeyInterface, err := x509.ParsePKIXPublicKey(derText)

	if err != nil {
		return err
	}

	publicKey, ok := publicKeyInterface.(*ecdsa.PublicKey)
	if !ok {
		return errors.New("断言失败，非ecds公钥!\n")
	}

	//2. 对原文生成哈希

	hash := sha256.Sum256(src)
	//3. 使用公钥验证

	//使用公钥验证hash值和两个大整数r、s构成的签名，并返回签名是否合法。
	//func Verify(pub *PublicKey, hash []byte, r, s *big.Int) bool

	isValid := ecdsa.Verify(publicKey, hash[:], sig.r, sig.s)

	if !isValid {
		return errors.New("校验失败!")
	}

	return nil
}

func main() {

	src := []byte("Golang，不支持加解密，支持ECC签名")

	sig, err := eccSignData(EccPrivateKeyFile, src)

	if err != nil {
		fmt.Printf("err : %s\n", err)
	}

	fmt.Printf("signature : %s\n", sig)
	fmt.Printf("signature hex : %x\n", sig)

	fmt.Printf("++++++++++++++++++=\n")

	src1 := []byte("Golang，不支持加解密，支持ECC签名!!!!!!!!!!!")
	err = eccVerifySig(EccPublicKeyFile, src1, sig)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("签名校验成功!\n")
}
