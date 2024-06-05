package usecase

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
