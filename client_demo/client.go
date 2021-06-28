package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	// rpc 连接服务器
	conn, err := rpc.Dial("tcp", "127.0.0.1:8811")
	if err != nil {
		fmt.Println("dial falut")
	}
	defer conn.Close()
	// rpc 调用远程
	var res string
	err = conn.Call("hello.HelloWorld", "maxu", &res)
	if err != nil {
		fmt.Println("conn.Call falut, err=", err)
	}
	fmt.Println(res)
}