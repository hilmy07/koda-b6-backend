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
	<div style="font-family: Arial, sans-serif; background-color: #f4f4f4; padding: 20px;">
		<div style="max-width: 500px; margin: auto; background: #ffffff; padding: 30px; border-radius: 10px; text-align: center; box-shadow: 0 4px 10px rgba(0,0,0,0.1);">
			
			<h2 style="color: #333;">🔐 OTP Verification</h2>
			
			<p style="color: #555;">Use the code below to proceed:</p>
			
			<div style="font-size: 32px; font-weight: bold; letter-spacing: 5px; color: #2c7be5; margin: 20px 0;">
				%s
			</div>
			
			<p style="color: #888; font-size: 14px;">
				This code is valid for a few minutes only.<br/>
				Do not share this code with anyone.
			</p>

			<hr style="margin: 20px 0; border: none; border-top: 1px solid #eee;" />

			<p style="font-size: 12px; color: #aaa;">
				If you did not request this OTP, please ignore this email.
			</p>

		</div>
	</div>
	`, otp)

	m.SetBody("text/html", body)

	d := gomail.NewDialer(
		"sandbox.smtp.mailtrap.io",
		2525,
		"2c8b3df5875788", // ganti dari Mailtrap
		"1a26ee3572183e", // ganti dari Mailtrap
	)

	return d.DialAndSend(m)
}