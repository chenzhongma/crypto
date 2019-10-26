package main

import (
	"io/ioutil"
	"log"
	"crypto/x509"
	"crypto/tls"
	"net/http"
	"fmt"
)

func main() {

	//1. 注册给服务器颁发证书的ca
	//- 读取ca证书
	caCertInfo, err := ioutil.ReadFile("./server.crt")

	if err != nil {
		log.Fatal(err)
	}

	//- 把ca的证书添加到ca池中
	//- 创建ca pool
	caCertPool := x509.NewCertPool()
	//添加caCert
	caCertPool.AppendCertsFromPEM(caCertInfo)

	//
	//1.5 加载客户端的证书和秘钥 ==> clientCert(修改了)
	//func LoadX509KeyPair(certFile, keyFile string) (Certificate, error) {
	clientCert, err := tls.LoadX509KeyPair("./client.crt", "./client.key")

	if err != nil {
		log.Fatal(err)
	}

	//
	//2. 配置tls, ==》 增加clientCert(修改了)
	//- RootCAs
	//Certificates
	cfg := tls.Config{
		//服务器的ca池
		RootCAs: caCertPool,

		//客户端证书
		Certificates: []tls.Certificate{clientCert},
		//ClientAuth:   tls.RequireAndVerifyClientCert,
	}

	//
	//3.创建http client
	client := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &cfg,
		},
	}

	//4. client发起请求
	response, err := client.Get("https://localhost:8848")
	if err != nil {
		log.Fatal(err)
	}

	bodyInfo, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	//5. 打印返回值
	fmt.Printf("body info : %s\n", bodyInfo)
	fmt.Printf("status code : %s\n", response.Status)
}
