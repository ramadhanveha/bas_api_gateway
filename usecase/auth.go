package usecase

// import (
// 	"api_gateway/handler"

// 	"github.com/gin-gonic/gin"
// )

type LoginInterface interface {
	Autentikasi(username string, password string) bool
}

type Login struct {
	validUsername string
	validPassword string
}

func NewLogin(username, password string) LoginInterface {
	return &Login{
		validUsername: username,
		validPassword: password,
	}
}

func (a *Login) Autentikasi(username string, password string) bool {
	return username == a.validUsername && password == a.validPassword
}

// func auth() {
// 	r := gin.Default()
// 	accountRoute := r.Group("/auth")
// 	accountRoute.POST("/create", handler.NewAuth().LoginAuth)
// }
