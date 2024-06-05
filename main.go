package main

import (
	"api_gateway/usecase"
	"fmt"
)

func main() {
	// Membuat objek login baru
	login := usecase.NewLogin("admin", "admin123")

	// Melakukan autentikasi
	if login != nil && login.Autentikasi("admin", "admin123") {
		fmt.Println("Login Success")
	} else {
		fmt.Println("Login Fail")
	}
}
