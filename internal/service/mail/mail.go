package service

import (
	"fmt"
	"math/rand"
	"time"

	"gopkg.in/gomail.v2"
)

const verifyCodeLength = 6

func SendEmail(verifyCode int, userEmail string) error {
	msg := gomail.NewMessage()
	msg.SetHeader("From", "miller.blanda80@ethereal.email")
	msg.SetHeader("To", "19120390@student.hcmus.edu.vn")
	msg.SetHeader("Subject", "<paste the subject of the mail>")
	msg.SetBody("text/html", "<b>This is the body of the mail</b>")

	n := gomail.NewDialer("smtp.gmail.com", 587, "miller.blanda80@ethereal.email", "XjKxJGPdJC42ZUhzBb")

	// Send the email
	if err := n.DialAndSend(msg); err != nil {
		fmt.Println(err)
	}
	return nil
}

func GenerateVerifyCode() int {
	rand.Seed(time.Now().UnixNano())
	a := rand.Int() % 10
	for a == 0 {
		a = rand.Int() % 10
	}

	for i := 1; i < verifyCodeLength; i++ {
		b := rand.Int() % 10
		a = a*10 + b
	}
	return int(a)
}
