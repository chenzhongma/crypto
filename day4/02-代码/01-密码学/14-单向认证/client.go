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
	//- 读取ca证书, 我们的证书是自签名的，server.crt能够认证自己，server.crt当成CA证书
	caCerInfo /*pem格式*/ , err := ioutil.ReadFile("./server.crt")
	if err != nil {
		log.Fatal(err)
	}

	//- 把ca的证书添加到ca池中
	//- 创建ca池
	cerPool := x509.NewCertPool()

	//- 将ca添加到ca池
	cerPool.AppendCertsFromPEM(caCerInfo)

	//
	//2. 配置tls
	// RootCAs defines the set of root certificate authorities
	// that clients use when verifying server certificates.
	// If RootCAs is nil, TLS uses the host's root CA set.
	//RootCAs *x509.CertPool

	//将我们承认ca池配置给tls
	cfg := tls.Config{
		RootCAs: cerPool,
	}

	//fmt.Printf("cfg : %s", cfg)

	//3.创建http client
	client := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &cfg,
			//TLSClientConfig: nil,
		},
	}

	//4. client发起请求
	response, err := client.Get("https://localhost:8848")
	if err != nil {
		log.Fatal(err)
	}

	//5. 打印返回值
	bodyInfo, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	//勿忘
	response.Body.Close()

	//body
	fmt.Printf("body : %s\n", bodyInfo)

	//状态码
	fmt.Printf("status code : %s\n", response.Status)
}
