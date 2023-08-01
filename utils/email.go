package utils

import (
	"bytes"
	"crypto/tls"
	"log"
	"os"
	"path/filepath"
	"text/template"

	"github.com/k3a/html2text"
	"github.com/panbhatt/go-gin-crud-gorm/initializers"
	"github.com/panbhatt/go-gin-crud-gorm/models"
	"gopkg.in/gomail.v2"
)

type EmailData struct {
	URL       string
	FirstName string
	Subject   string
}

// Email Template Parser

func ParseTemplateDir(dir string) (*template.Template, error) {

	var paths []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			paths = append(paths, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return template.ParseFiles(paths...)

}

func SendEmail(user *models.User, data *EmailData) {

	configObj := initializers.CO

	// Sender Data
	from := configObj.EmailFrom
	smtpPassword := configObj.SMTPPass
	smtpUser := configObj.SMTPUser
	to := user.Email
	smtpHost := configObj.SMTPHost
	smtpPort := configObj.SMTPPort

	var body bytes.Buffer
	template, err := ParseTemplateDir("templates")
	if err != nil {
		log.Fatal("Could not parse Template", err)
	}

	template.ExecuteTemplate(&body, "verificationCode.html", &data)

	m := gomail.NewMessage()

	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", data.Subject)
	m.SetBody("text/html", body.String())
	m.AddAlternative("text/plain", html2text.HTML2Text(body.String()))

	d := gomail.NewDialer(smtpHost, smtpPort, smtpUser, smtpPassword)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: false}

	// Send Email
	if err := d.DialAndSend(m); err != nil {
		log.Fatal("Cound not send email : ", err)
	}
}
