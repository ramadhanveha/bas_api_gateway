package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthInterface interface {
	LoginAuth(*gin.Context)
}

type authImplement struct{}

func NewAuth() AuthInterface {
	return &authImplement{}
}

type AuthPayloadAccount struct {
	Username string
	Password string
}

func (a *authImplement) LoginAuth(g *gin.Context) {
	bodyPayload := AuthPayloadAccount{}

	err := g.BindJSON(&bodyPayload)
	if err != nil {
		g.JSON(http.StatusUnauthorized, gin.H{
			"message": "anda gagal login",
		})
	}

	g.JSON(http.StatusOK, gin.H{
		"message": "anda berhasil login",
		"data":    bodyPayload,
	})

}

// func (a *Implement) TransactionAuth(g *gin.Context) {
// 	bodyPayload := AuthPayloadAccount{}

// 	err := g.BindJSON(&bodyPayload)
// 	if err != nil {
// 		g.JSON(http.StatusUnauthorized, gin.H{
// 			"message": "anda gagal login",
// 		})
// 	}

// 	g.JSON(http.StatusOK, gin.H{
// 		"message": "anda berhasil login",
// 		"data":    bodyPayload,
// 	})

// }
