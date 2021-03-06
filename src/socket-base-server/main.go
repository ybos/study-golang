package main

import (
	"fmt"
	//	"io/ioutil"
	"net"
	"os"
	"time"
)

func main() {
	service := ":7777"

	// 转换 ip 到 tcpAddr 结构
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	// 监听地址
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		// 接收请求
		conn, err := listener.Accept()

		if err != nil {
			continue
		}

		//		request, err := ioutil.ReadAll(conn)
		//		checkError(err)

		//		fmt.Println(request)

		daytime := time.Now().String()
		fmt.Println(daytime)

		// 写数据
		conn.Write([]byte(daytime))
		conn.Close()
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s\n", err.Error())
		os.Exit(1)
	}
}
