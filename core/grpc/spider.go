package grpc

import (
	"fmt"

	"google.golang.org/grpc"
)

//PORT 跟爬虫的端口保持一致
const PORT = ":50051"

//InitService 初始化grpc连接
func InitService( /* port string */ ) {
	conn, err := grpc.Dial("127.0.0.1"+PORT, grpc.WithInsecure())
	if err != nil {
		fmt.Println("network error", err)
	}

	//网络延迟关闭
	defer conn.Close()

}
