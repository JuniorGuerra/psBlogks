package main

import (
	"net/smtp"
	"strconv"
)

func email(mail string, code int) bool {

	// Sender data.
	from := "blogbookresponse@gmail.com"
	password := "xovhjjlhfigxphcg"

	// Receiver email address.
	to := []string{
		mail,
	}

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Message.
	message := []byte(string("👋Hola un gusto saludarte, este es tu codigo de verificacion " + strconv.Itoa(code)))

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)

	return err == nil
}

func recuperate_password_email(mail string, code int) bool {

	// Sender data.
	from := "blogbookresponse@gmail.com"
	password := "xovhjjlhfigxphcg"

	// Receiver email address.
	to := []string{
		mail,
	}

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Message.
	message := []byte(string("👋Hola un gusto saludarte, Tu codigo para canviar contraseña " + strconv.Itoa(code) + " Si este correo no lo pediste tu hacer caso omiso al mismo") )

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)

	return err == nil
}
