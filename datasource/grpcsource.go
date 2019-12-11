package datasource

import (
	"github.com/birjemin/gin-structure/conf"
	"google.golang.org/grpc"
	"log"
)

var grpcConn *grpc.ClientConn

func GetGRPC() *grpc.ClientConn {
	return grpcConn
}

// 关闭db
func CloseGRPC() error {
	if grpcConn != nil {
		return grpcConn.Close()
	} else {
		return nil
	}
}

func init() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(conf.String("rpc.server"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	grpcConn = conn
}
