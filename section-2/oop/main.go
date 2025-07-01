package main

import (
	"github.com/fandiaji45/golang-academy/section-2/oop/usecase"
)

func main() {
	// Inisialisasi usecase untuk pengiriman email
	senderUsecase := usecase.NewSmtpSenderUsecase()

	// Inisialisasi usecase untuk user dengan menyuntikkan dependency senderUsecase
	userUsecase := usecase.NewUserUsecase(senderUsecase)

	// Registrasi user baru dengan email tertentu
	userUsecase.RegistrasiUser("fndwbw688@gmail.com")
}
