package function

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"text/template"

	"gopkg.in/gomail.v2"
)

func SendVerificationEmail(email, verifyToken, reference_code string) error {
	mailerHost := os.Getenv("MAILER_HOST")
	mailerPort, _ := strconv.Atoi(os.Getenv("MAILER_PORT"))
	mailerUsername := os.Getenv("MAILER_USERNAME")
	mailerPassword := os.Getenv("MAILER_PASSWORD")

	tmpl, err := template.ParseFiles("./email-template/verification-email-template.html")
	if err != nil {
		return fmt.Errorf("error reading email template: %w", err)
	}

	data := struct {
		VerifyToken   string
		ReferenceCode string
	}{
		VerifyToken:   verifyToken,
		ReferenceCode: reference_code,
	}

	var body bytes.Buffer
	err = tmpl.Execute(&body, data)
	if err != nil {
		return fmt.Errorf("error executing email template: %w", err)
	}

	m := gomail.NewMessage()
	m.SetHeader("From", mailerUsername)
	m.SetHeader("To", email)
	m.SetHeader("Subject", "ยืนยันอีเมลของคุณ - Verify Your Email")
	m.SetBody("text/html", body.String())

	d := gomail.NewDialer(mailerHost, mailerPort, mailerUsername, mailerPassword)
	if err := d.DialAndSend(m); err != nil {
		return fmt.Errorf("could not send email: %w", err)
	}

	return nil
}

func SendRequestResetPwdEmail(email, resetPwdToken string) error {
	mailerHost := os.Getenv("MAILER_HOST")
	mailerPort, _ := strconv.Atoi(os.Getenv("MAILER_PORT"))
	mailerUsername := os.Getenv("MAILER_USERNAME")
	mailerPassword := os.Getenv("MAILER_PASSWORD")

	tmpl, err := template.ParseFiles("./email-template/reset-password-template.html")
	if err != nil {
		return fmt.Errorf("error reading email template: %w", err)
	}

	data := struct {
		string
		resetPwdToken string
	}{

		resetPwdToken: resetPwdToken,
	}

	var body bytes.Buffer
	err = tmpl.Execute(&body, data)
	if err != nil {
		return fmt.Errorf("error executing email template: %w", err)
	}

	m := gomail.NewMessage()
	m.SetHeader("From", mailerUsername)
	m.SetHeader("To", email)
	m.SetHeader("Subject", "เปลี่ยนรหัสผ่านของคุณ - Reset Your Password")
	m.SetBody("text/html", body.String())

	d := gomail.NewDialer(mailerHost, mailerPort, mailerUsername, mailerPassword)
	if err := d.DialAndSend(m); err != nil {
		return fmt.Errorf("could not send email: %w", err)
	}

	return nil
}
