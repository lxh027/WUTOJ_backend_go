package rpcconn

import (
	"log"

	"google.golang.org/grpc"
)

//RPCPORT 爬虫的端口
const RPCPORT = ":50051"

//RPCAddress RPC服务器地址
const RPCAddress = "127.0.0.1"

//RPCConn gRpc连接
var RPCConn *grpc.ClientConn

func init() {
	var err error
	//客户端连接服务端
	RPCConn, err = grpc.Dial(RPCAddress+RPCPORT, grpc.WithInsecure())
	if err != nil {
		log.Panic("network error", err)
	}
}
