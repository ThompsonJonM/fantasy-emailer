package email

import (
	"bytes"
	"fmt"
	"github.com/ThompsonJonM/fantasy-emailer/src/players/players"
	"html/template"
	"log"
	"net/smtp"
	"strings"
)

type smtpServer struct {
	host string
	port string
}

func (s *smtpServer) Address() string {
	return s.host + ":" + s.port
}

func getMessageString(fromEmail, To, Subject, emailBody string) []byte {
	return []byte("From: " + fromEmail + "\r\n" + "To: " + To + "\r\n" + "Subject: " + Subject + "\r\n" + "MIME-Version: 1.0\r\n" + "Content-Type: text/html; charset=\"utf-8\"\r\n\r\n" + emailBody + "\r\n")
}

func parseTemplate(templateName string, data interface{}) (string, error) {
	t, err := template.ParseFiles(templateName)
	if err != nil {
		log.Fatalln("Could not parse template", err)
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func SendEmail(from, pass, file, subject, template string) {
	ps := players.ImportPlayers(file)

	var i int

	for _, v := range ps {
		to := []string{
			v.Email,
		}
		f := strings.Split(v.Name, " ")
		first := strings.Join(f[:1], "")

		s := smtpServer{host: "smtp.gmail.com", port: "587"}
		ms, err := parseTemplate(template, first)
		if err != nil {
			log.Fatalln("Could not parse template", err)
		}

		b := getMessageString(from, to[0], subject, ms)

		auth := smtp.PlainAuth("", from, pass, s.host)

		err = smtp.SendMail(s.Address(), auth, from, to, b)
		if err != nil {
			log.Fatalln("Could not send e-mail", err)
		}

		i++

		fmt.Printf("%d of %d emails sent.\n", i, len(ps))
	}
}
