package handler

import (
	"api_gateway/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AccountInterface interface {
	GetAccount(*gin.Context)
	CreateAccount(*gin.Context)
	UpdateAccount(*gin.Context)
	DeleteAccount(*gin.Context)
}

type accountImplement struct{}

// func NewAccount() AccountInterface {
// 	 return &accountImplement{}
// }

func (a *accountImplement) AccountAuth(g *gin.Context) {
	queryParam := g.Request.URL.Query()

	name := queryParam.Get("name")

	g.JSON(http.StatusOK, gin.H{
		"message": "Get account successfully",
		"data":    name,
	})
}

type BodyPayloadAccount struct {
	AccountID string
	Name      string
	Address   string
}

func (a *accountImplement) CreateAuth(g *gin.Context) {
	bodyPayload := BodyPayloadAccount{}

	err := g.BindJSON(&bodyPayload)
	if err != nil {
		g.AbortWithError(http.StatusBadRequest, err)
	}

	g.JSON(http.StatusOK, gin.H{
		"message": "Post account successfully",
		"data":    bodyPayload,
	})
}

func (a *accountImplement) UpdateAccount(g *gin.Context) {
	queryParam := g.Request.URL.Query()

	name := queryParam.Get("name")

	g.JSON(http.StatusOK, gin.H{
		"message": "Get account successfully",
		"data":    name,
	})
}

func (a *accountImplement) DeleteAccount(g *gin.Context) {
	queryParam := g.Request.URL.Query()

	id := queryParam.Get("id")

	g.JSON(http.StatusOK, gin.H{
		"message": "Delete account successfully",
		"data":    id,
	})
}

type BalanceInterface interface {
	BalanceAuth(*gin.Context)
}

type BalanceImplement struct{}

type BalancePayloadAccount struct {
	Username string
	Password string
}

func NewBalance() BalanceInterface {
	return &BalanceImplement{}
}

func (a *BalanceImplement) BalanceAuth(g *gin.Context) {
	bodyPayload := BalancePayloadAccount{}

	err := g.BindJSON(&bodyPayload)
	if err != nil {
		g.JSON(http.StatusUnauthorized, gin.H{
			"message": "anda gagal login",
		})
	}

	if usecase.NewLogin("veha", "password").Autentikasi(bodyPayload.Username, bodyPayload.Password) {
		g.JSON(http.StatusOK, gin.H{
			"message": "Anda berhasil login",
			"data":    bodyPayload,
		})
	} else {
		g.JSON(http.StatusUnauthorized, gin.H{
			"message": "Anda gagal login",
			"data":    bodyPayload,
		})
	}
}
