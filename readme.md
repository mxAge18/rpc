
# go 自带rpc包实现远程函数调用

## server_demo/server.go

    //注册rpc服务，维护一个hash表，key值是服务名称，value值是服务的地址
 rpc.RegisterName("HelloService",new(HelloService))
    // net包调用监听某个端口的管道，管道接收客户端发来的请求，并进行处理后返回数据
    listener, err := net.Listen()
    listener.Accept()
    //rpc调用,并返回执行后的数据
    //1.read，获取服务名称和方法名，获取请求数据
    //2.调用对应服务里面的方法，获取传出数据
    //3.write，把数据返回给client
 rpc.ServeConn(conn)

## client_demo/client.go

- rpc.Dial
// 调用远程地址，建立连接，返回一个conn
conn.Call("hello.HelloWorld", "maxu", &res) //第一个参数是用点号链接的RPC服务名字和方法名字，第二和第三个参数分别我们定义RPC方法的两个参数。

## 测试

go run server_demo/server.go

go run client_demo/client.go
