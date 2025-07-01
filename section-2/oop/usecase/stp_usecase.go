package usecase

import "fmt"

// SmtpSenderUsecase adalah struct kosong yang merepresentasikan pengirim email via SMTP
type SmtpSenderUsecase struct{}

// NewSmtpSenderUsecase adalah constructor untuk membuat instance baru dari SmtpSenderUsecase
func NewSmtpSenderUsecase() *SmtpSenderUsecase {
	return &SmtpSenderUsecase{}
}

// ISmtpSenderUsecase adalah interface untuk mengirim email
type ISmtpSenderUsecase interface {
	KirimEmail(penerima, pesan string)
}

// KirimEmail mengimplementasikan fungsi kirim email dengan output ke console
func (s *SmtpSenderUsecase) KirimEmail(penerima, pesan string) {
	fmt.Printf("[SMTP] Mengirim email ke %s: %s\n", penerima, pesan)
}
