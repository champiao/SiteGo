package mail

import (
	"fmt"
	"log"
	"os"

	strconv "strconv"

	"github.com/joho/godotenv"
	mail "gopkg.in/mail.v2"
)

func SendEmail(name string, email string, message string, phone string) error {
	m := mail.NewMessage()
	m.SetHeader("From", GetEnv("EMAIL_USER", "champiao@champiao.com.br"))
	m.SetHeader("To", GetEnv("EMAIL_USER", "champiao@champiao.com.br"))
	m.SetHeader("Subject", "nome do cliente: "+name+"  Email: "+email+"  Telefone: "+phone)
	m.SetBody("text/html", message)
	port, errPort := strconv.Atoi(GetEnv("EMAIL_PORT", "587"))
	if errPort != nil {
		log.Printf("smtp error: %s", errPort)
		return errPort
	}
	host := GetEnv("EMAIL_HOST", "smtp.zoho.com")
	mailUser := GetEnv("EMAIL_USER", "champiao@champiao.com.br")
	mailPass := GetEnv("EMAIL_PASSWORD", "")
	d := mail.NewDialer(host, port, mailUser, mailPass)
	d.TLSConfig = nil

	if err := d.DialAndSend(m); err != nil {
		log.Printf("smtp error: %s", err)
		return err
	}

	return nil
}
func GetEnv(key string, fallback string) string {
	errEnv := godotenv.Load()
	if errEnv != nil {
		log.Fatal("env file not found")
	}
	value := os.Getenv(key)
	if len(value) == 0 {
		fmt.Printf("erro ao carregar variável")
		value = fallback
	}
	return value
}
