package gmail
import (
    "net/smtp"
    "log")

func Send(to, from, subject, msg, pwd string) {
    body := "To: " + to + "\r\nSubject: " + subject + "\r\n\r\n" + msg
    auth := smtp.PlainAuth("",from,pwd,"smtp.gmail.com")
    err := smtp.SendMail("smtp.gmail.com:587",auth, from, []string{to},[]byte(body))
    if err != nil {
        log.Fatal(err)
    }
}
