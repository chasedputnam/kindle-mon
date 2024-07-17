package mailer

import (
	"log"
	"os"
	"time"

	gomail "gopkg.in/mail.v2"
)

var (
	server   = ""
	smtpPort = 587
	password = ""
	sender   = "chasedputnam@gmail.com"
	receiver = " " //kindle inbox address
)

func Send(files []string, timeout int) {
	//cfg := config.GetInstance()
	msg := gomail.NewMessage()
	msg.SetHeader("From", sender)
	msg.SetHeader("To", receiver)

	msg.SetBody("text/plain", "")

	attachedFiles := make([]string, 0)
	for _, file := range files {
		_, err := os.Stat(file)
		if err != nil {
			log.Printf("Couldn't find the file %s : %s \n", file, err)
			continue
		} else {
			msg.Attach(file)
			attachedFiles = append(attachedFiles, file)
		}
	}
	if len(attachedFiles) == 0 {
		log.Println("No files to send")
		return
	}

	dialer := gomail.NewDialer(server, smtpPort, sender, password)
	dialer.Timeout = time.Duration(timeout) * time.Second
	log.Println("Sending mail")
	log.Println("Mail timeout : ", dialer.Timeout.String())
	log.Println("Following files will be sent :")
	for i, file := range attachedFiles {
		log.Printf("%d. %s\n", i+1, file)
	}

	if err := dialer.DialAndSend(msg); err != nil {
		log.Println("Error sending mail : ", err)
		return
	} else {
		log.Printf("Mailed %d files to %s", len(attachedFiles), receiver)
	}

}
