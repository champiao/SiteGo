package mail

import (
	"fmt"
	"log"
	"os"

	strconv "strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	mail "gopkg.in/mail.v2"
)

func SendEmail(c *gin.Context, name string, email string, message string, phone string) error {
	errEnv := godotenv.Load()
	if errEnv != nil {
		log.Fatal("env file not found")
	}
	m := mail.NewMessage()
	m.SetHeader("From", getEnv("EMAIL_USER", "champiao@champiao.com.br"))
	m.SetHeader("To", getEnv("EMAIL_USER", "champiao@champiao.com.br"))
	m.SetHeader("Subject", "nome do cliente: "+name+"  Email: "+email+"  Telefone: "+phone)
	m.SetBody("text/html", message)
	port, errPort := strconv.Atoi(getEnv("EMAIL_PORT", "587"))
	if errPort != nil {
		log.Printf("smtp error: %s", errPort)
		return errPort
	}
	host := getEnv("EMAIL_HOST", "smtp.zoho.com")
	mailUser := getEnv("EMAIL_USER", "champiao@champiao.com.br")
	mailPass := getEnv("EMAIL_PASSWORD", "")
	d := mail.NewDialer(host, port, mailUser, mailPass)
	d.TLSConfig = nil

	if err := d.DialAndSend(m); err != nil {
		log.Printf("smtp error: %s", err)
		return err
	}

	return nil
}
func getEnv(key string, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		fmt.Printf("erro ao carregar variável")
		value = fallback
	}
	return value
}
