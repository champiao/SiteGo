package main

import (
	"fmt"
	"net/http"
	"os"

	SendEmail "champiao/func/mail"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Static("/public", "./public")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.GET("/about", func(c *gin.Context) {
		c.HTML(http.StatusOK, "about.html", nil)
	})

	r.GET("/contact", func(c *gin.Context) {
		c.HTML(http.StatusOK, "contact.html", nil)
	})
	r.GET("/aluraCerts", func(c *gin.Context) {
		pdfBytes, err := os.ReadFile("public/certs/Certificado_Alura.pdf")
		if err != nil {
			c.String(http.StatusInternalServerError, "Erro ao exibir o certificado Alura")
			return
		}

		// Defina os cabeçalhos apropriados para enviar um PDF
		c.Header("Content-Disposition", "inline; filename=Certificado_Alura.pdf")
		c.Data(http.StatusOK, "application/pdf", pdfBytes)
	})
	r.GET("/certificates", func(c *gin.Context) {
		c.HTML(http.StatusOK, "certificados.html", nil)
	})

	r.POST("/contact", func(c *gin.Context) {
		name := c.PostForm("name")
		email := c.PostForm("email")
		message := c.PostForm("message")
		phone := c.PostForm("phone")
		sendEmail(c, name, email, message, phone)
	})

	r.LoadHTMLGlob("public/views/*")

	r.Run(":8080")
}

func sendEmail(c *gin.Context, name string, email string, message string, phone string) {
	send := SendEmail.SendEmail(c, name, email, message, phone)
	if send != nil {
		c.HTML(http.StatusBadRequest, "contact.html", gin.H{
			"error":   "Ocorreu um erro ao tentar enviar o formulário",
			"success": "",
		})
		fmt.Printf("erro: %s", send)
		return
	} else if send == nil {
		c.Redirect(http.StatusMovedPermanently, "/contact")
		c.HTML(http.StatusOK, "contact.html", gin.H{
			"error":   "",
			"success": "Email enviado com sucesso",
		})
		return
	}
}
