package controllers

import (
	"github.com/birjemin/gin-structure/datasource"
	pb "github.com/birjemin/gin-structure/grpc/pb"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)


func SayHello(g *gin.Context) {
	name := cast.ToString(g.Query("name"))
	// Contact the server and print out its response.
	client := pb.NewGreeterClient(datasource.GetGRPC())
	res, err := client.SayHello(g, &pb.HelloRequest{Name: name})
	if err != nil {
		JsonReturn(g, 0, "获取失败", err)
		return
	}
	JsonReturn(g, 0, "", res)
	return
}
