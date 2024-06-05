package main

import (
	"api_gateway/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	accountRoute := r.Group("/auth")
	accountRoute.POST("/login", handler.NewAuth().LoginAuth)

	balanceRoute := r.Group("/account")
	balanceRoute.POST("/balance", handler.NewBalance().BalanceAuth)

	transactionRoute := r.Group("/transaction")
	transactionRoute.POST("/transfer-bank", handler.NewTransaction().TransactionAuth)
	// r := gin.Default()

	// accountRoute := r.Group("/account")
	// accountRoute.GET("/get", handler.NewAccount().GetAccount)
	// accountRoute.POST("/create", handler.NewAccount().CreateAccount)
	// accountRoute.DELETE("/delete/:id", handler.NewAccount().DeleteAccount)
	// accountRoute.PATCH("/update/:id", handler.NewAccount().UpdateAccount)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
