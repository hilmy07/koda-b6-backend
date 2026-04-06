package lib

import (
	"fmt"

	"gopkg.in/gomail.v2"
)

func SendOTPEmail(to string, otp string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "no-reply@example.com")
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Kode OTP Kamu")

	body := fmt.Sprintf(`
		<h2>OTP Verifikasi</h2>
		<p>Kode OTP kamu adalah:</p>
		<h1>%s</h1>
		<p>Berlaku beberapa menit saja.</p>
	`, otp)

	m.SetBody("text/html", body)

	d := gomail.NewDialer(
		"sandbox.smtp.mailtrap.io",
		2525,
		"SMTP_USER", // ganti dari Mailtrap
		"SMTP_PASS", // ganti dari Mailtrap
	)

	return d.DialAndSend(m)
}