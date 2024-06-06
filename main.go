package main

import (
	"api_gateway/handler"
	"api_gateway/proto"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	//"github.com/go-micro/generator/cmd/protoc-gen-micro/plugin/micro"
	"github.com/go-micro/plugins/v4/client/grpc"
	micro "go-micro.dev/v4"
	"go-micro.dev/v4/client"
	// "go-micro.dev/v4/util/grpc"
)

func main() {
	r := gin.Default()
	addrServiceTransactionOpt := client.WithAddress(":9000")
	clientSrvTransaction := grpc.NewClient()

	srvTransaction := micro.NewService(
		micro.Client(clientSrvTransaction),
	)

	srvTransaction.Init(
		micro.Name("service-transaction"),
		micro.Version("latest"),
	)

	accountRoute := r.Group("/account")
	accountRoute.GET("/get", handler.NewAccount().GetAccount)
	accountRoute.POST("/create", handler.NewAccount().CreateAccount)
	accountRoute.PATCH("/update/:id", handler.NewAccount().UpdateAccount)
	accountRoute.DELETE("/delete/:id", handler.NewAccount().DeleteAccount)
	accountRoute.GET("/balance", handler.NewAccount().GetBalance)

	transactionRoute := r.Group("/transaction")
	transactionRoute.POST("/transferbank", handler.NewTransaction().TransferBank)
	transactionRoute.POST("/transfer-bank", handler.NewTransaction().CreateTransaction)
	transactionRoute.GET("/get", func(g *gin.Context) {
		ClientResponse, err := proto.NewServiceTransactionService("service-transaction",
			srvTransaction.Client()).Call(context.Background(), &proto.CallRequest{
			Name: "Lupi",
		}, addrServiceTransactionOpt)
		if err != nil {
			g.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
		g.JSON(http.StatusOK, gin.H{
			"data": ClientResponse,
		})
	})

	authRoute := r.Group("/auth")
	authRoute.POST("/login", handler.NewAuth().Login)
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
