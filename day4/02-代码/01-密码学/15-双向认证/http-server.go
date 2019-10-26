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

	//1. 注册client ca证书
	//- 读取client的ca证书, client的证书也是自签名的，自己认证自己
	caInfo, err := ioutil.ReadFile("./client.crt")
	if err != nil {
		log.Fatal(err)
	}

	//- 创建ca池
	caCertPool := x509.NewCertPool()

	//- 把client 的 ca 添加到ca池
	caCertPool.AppendCertsFromPEM(caInfo)

	//2. 配置tls ==> cfg
	cfg := tls.Config{
		// 我们要认证client, 需要两个字段
		// ClientAuth determines the server's policy for
		// TLS Client Authentication. The default is NoClientCert.
		//ClientAuth ClientAuthType

		//	const (
		//	NoClientCert ClientAuthType = iota
		//	RequestClientCert
		//	RequireAnyClientCert
		//	VerifyClientCertIfGiven
		//	RequireAndVerifyClientCert
		//)

		//我们设置服务器认证客户端
		ClientAuth: tls.RequireAndVerifyClientCert,

		// ClientCAs defines the set of root certificate authorities
		// that servers use if required to verify a client certificate
		// by the policy in ClientAuth.
		//ClientCAs *x509.CertPool
		ClientCAs: caCertPool, //客户端的ca池填充在这里
	}

	//3. 创建http server， 使用cfg
	server := http.Server{
		//三个字段Addr, Handler, TLSConfig
		Addr: ":8848",
		//Handler:   nil,
		Handler:   &myhandler{},
		TLSConfig: &cfg,
	}

	fmt.Printf("准备启动服务器...\n")
	//4. 启动http server，启动时加载自己的证书， 启动时使用tls
	err = server.ListenAndServeTLS("./server.crt", "./server.key")
	if err != nil {
		log.Fatal(err)
	}
}

type myhandler struct {
}

func (h myhandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("ServeHTTP called!\n")
	w.Write([]byte("hello world!!!!"))
}
