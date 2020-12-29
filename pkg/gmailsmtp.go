package gmailsmtp

import (
	"fmt"
	"net/smtp"
	"os"
)

var emailAuth smtp.Auth

func SendEmailSMTP(to []string, data interface{}, template string) (bool, error) {
	emailHost := "smtp.gmail.com"
	emailFrom := os.Getenv("GMAILUSER")
	emailPassword := os.Getenv("GMAILPASS")
	fmt.Println(emailFrom)
	emailPort := "587"

	emailAuth = smtp.PlainAuth("", emailFrom, emailPassword, emailHost)

	emailBody := "TEST"
	//emailBody := "<!DOCTYPE html><html lang=\"en\"><head><meta charset=\"utf-8\" /><meta http-equiv=\"x-ua-compatible\" content=\"ie=edge\" /><meta name=\"viewport\" content=\"width=device-width, initial-scale=1\" /><title></title><link rel=\"stylesheet\" href=\"css/main.css\" /><link rel=\"icon\" href=\"images/favicon.png\" /></head><body><h1>Congratualtions! You won!!</h1><img src=\"http://thegalley.com/galleyblog/wp-content/uploads/sites/2/2013/12/995589_728552843822962_1742475431_n.png\" alt=\"LA\" style=\"width:100%\"><div class=\"\"><a href=\"https://www.youtube.com/watch?v=dQw4w9WgXcQ\">Only limited trme left! Click here!</a></div></body></html>"
	/*
		emailBody, err := parseTemplate(template, data)
		if err != nil {
			return false, errors.New("unable to parse email template")
		}
	*/

	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	subject := "Subject: " + "TEST!\n"
	msg := []byte(subject + mime + "\n" + emailBody)
	addr := fmt.Sprintf("%s:%s", emailHost, emailPort)

	if err := smtp.SendMail(addr, emailAuth, emailFrom, to, msg); err != nil {
		return false, err
	}
	return true, nil
}

// func parseTemplate(templateFileName string, data interface{}) (string, error) {
// 	templatePath, err := filepath.Abs(fmt.Sprintf("gomail/email_templates/%s", templateFileName))
// 	if err != nil {
// 		return "", errors.New("invalid template name")
// 	}
// 	t, err := template.ParseFiles(templatePath)
// 	if err != nil {
// 		return "", err
// 	}
// 	buf := new(bytes.Buffer)
// 	if err = t.Execute(buf, data); err != nil {
// 		return "", err
// 	}
// 	body := buf.String()
// 	return body, nil
// }
