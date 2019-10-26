package main

import (
	"net/http"
	"log"
	"fmt"
)

func main() {

	//1. 创建http server
	server := http.Server{
		//Addr    string  // TCP address to listen on, ":http" if empty
		Addr: ":8848", //监听端口

		//Handler Handler // handler to invoke, http.DefaultServeMux if nil
		Handler: nil, //填写nil时， 会使用默认的处理器， 还是要自己实现处理逻辑

		//TLSConfig *tls.Config
		TLSConfig: nil,
	}

	//编写处理逻辑
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("HandleFunc called!\n")
		writer.Write([]byte("hello world!!!!!"))
	})

	//2. 启动http server，启动时加载自己的证书， 启动时使用tls
	//- 一般服务器代码时，路径会使用全局变量
	//- 这个全局变量的值会写到配置文件中
	//- 服务器启动时，读取到全局变量
	err := server.ListenAndServeTLS("./server.crt", "./server.key")

	if err != nil {
		log.Fatal(err)
	}
}
