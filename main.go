package main

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"

	"gopkg.in/gomail.v2"
)

func sendMailSimple(subject string, body string, to []string) {

	auth := smtp.PlainAuth(
		"",
		"senderEmail",
		"Send-Email-Password-Created-on-Email",
		"smtp.gmail.com",
	)

	msg := "Subject :" + subject + " \n" + body + ""

	err := smtp.SendMail("smtp.gmail.com:587",
		auth,
		"sender-email",
		to,
		[]byte(msg),
	)

	if err != nil {
		fmt.Println(err)
	}

}

func sendMailSimpleHtml(subject string, templatePath string, to []string) {

	//PlainAuth accepts four arguments of string type identity(It should be an empty string to act as username),
	//the username(sender mail address), password (sender mail password), and port of SMTP server. PlainAuth returns an Auth,
	//an implementation of an SMTP authentication mechanism. To authenticate to the host, the returned Auth uses the given username
	// and password and acts as an identity.

	auth := smtp.PlainAuth(
		"",
		"sender-email",
		"sender-email-password-generated-by-email",
		"smtp.gmail.com",
	)

	var body bytes.Buffer

	t, err := template.ParseFiles(templatePath)

	t.Execute(&body, struct{ Name string }{Name: "any text"})

	headers := "MIME-version :1.0;\nContent-Type: text/html; charset=\"UTF-8\";"

	msg := "Subject :" + subject + " \n" + headers + "\n\n" + body.String() + ""

	err = smtp.SendMail("smtp.gmail.com:587",
		auth,
		"sender-email",
		to,
		[]byte(msg),
	)

	if err != nil {
		fmt.Println(err)
	}
}

func sendGoMail(templatePath string) {
	var body bytes.Buffer

	t, err := template.ParseFiles(templatePath)
	t.Execute(&body, struct{ Name string }{Name: "any-text"})
	if err != nil {
		fmt.Println(err)
	}

	//Send With Go Mail
	m := gomail.NewMessage()
	m.SetHeader("From", "sender-email")
	m.SetHeader("To", "recevier-1-email", "reciever-2-email")
	//m.SetAddressHeader("Cc","sad")

	m.SetBody("text/html", body.String())
	m.Attach("./72401417-hello-vector-isolated-illustration-brush-calligraphy-hand-lettering-inspirational-typography.webp")

	d := gomail.NewDialer("smtp.gmail.com", 587, "sender-email", "sender-email-password")

	err = d.DialAndSend(m)
	if err != nil {
		panic(err)
	}

}

func main() {
	sendMailSimple("Subject Set From Argument", "Body Set from Arg", []string{"receiver-email"})

	sendMailSimpleHtml(
		"Sending Html as Body",
		"./index.html", []string{"reciever-email"},
	)
	sendGoMail("./index.html")
}
