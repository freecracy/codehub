package sendmail

import (
	"fmt"
	"log"
	"net/smtp"
	"strconv"
	"time"
)

var (
	// Username ...
	Username string
	password string
	host     string
	port     int
	from     string
	to       string
)

// Sendmail ...
func Sendmail() {
	auth := smtp.PlainAuth("cn", username, password, host)

	to := []string{to}

	local, _ := time.LoadLocation("Asia/Shanghai")
	subject := fmt.Sprintf("Subject:%s\n", time.Now().In(local).Format(time.RFC850))
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	content := "test email"
	msg := []byte(subject + mime + content)
	if err := smtp.SendMail(host+":"+strconv.Itoa(port), auth, from, to, msg); err != nil {
		log.Println(err.Error())
	}
}
