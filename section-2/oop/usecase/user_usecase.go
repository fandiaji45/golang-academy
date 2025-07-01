package usecase

import "fmt"

// UserUsecase bertanggung jawab untuk proses registrasi user dan pengiriman email
type UserUsecase struct {
	sender ISmtpSenderUsecase // dependency untuk mengirim email
}

// NewUserUsecase adalah constructor untuk membuat instance UserUsecase
func NewUserUsecase(sender ISmtpSenderUsecase) *UserUsecase {
	return &UserUsecase{
		sender: sender,
	}
}

// IUserUsecase adalah interface yang mendefinisikan kontrak fungsi RegistrasiUser
type IUserUsecase interface {
	RegistrasiUser(email string)
}

// RegistrasiUser adalah implementasi dari interface IUserUsecase
// Fungsi ini mencetak pesan sukses dan mengirimkan email sambutan
func (u *UserUsecase) RegistrasiUser(email string) {
	fmt.Println("Registrasi user berhasil:", email)
	u.sender.KirimEmail(email, "Selamat datang di platform kami!")
}
