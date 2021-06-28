package main

import (
	"fmt"
	"net"
	"net/rpc"
)

type ThisServer struct {

}
func (this *ThisServer) HelloWorld(name string, resp *string) error {
	*resp = name + " hi, how are u"
	return nil
}
func main() {
	// 注册rpc服务 绑定对象
	err := rpc.RegisterName("hello", new(ThisServer))
	if err != nil {
		fmt.Println("register failed, err=", err)
		return
	}
	// 监听设置
	listener, err := net.Listen("tcp", "0.0.0.0:8811")
	if err != nil {
		fmt.Println("listener set failed, err=", err)
		return
	}
	defer listener.Close()
	// 建立连接
	fmt.Println("start listen tcp conn")
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("listener.Accept() set failed, err=", err)
		return
	}
	defer conn.Close()
	// 绑定服务
	fmt.Println("connect rpc success")
	rpc.ServeConn(conn)
}