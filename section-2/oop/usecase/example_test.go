package usecase_test

import (
	"fmt"

	"github.com/fandiaji45/golang-academy/section-2/oop/usecase"
)

// mock sender biar gak kirim email beneran
type dummySender struct{}

func (d *dummySender) KirimEmail(email, pesan string) {
	fmt.Printf("[SMTP] Simulasi kirim ke %s: %s\n", email, pesan)
}

func ExampleUserUsecase_RegistrasiUser() {
	sender := &dummySender{}
	uuc := usecase.NewUserUsecase(sender)
	uuc.RegistrasiUser("andi@mail.com")

}
