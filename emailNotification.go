package main

import (
	"fmt"

	"blog.local/interfaces/types"
	"github.com/joho/godotenv"
	"gopkg.in/gomail.v2"
)

type EmailConfig struct {
	SMTPServer   string
	SMTPPort     int
	SMTPUsername string
	SMTPPassword string
	FromEmail    string
	ToEmail      string
}

func sendEmail(article types.IArticleText) (error) {

	envFile, _ := godotenv.Read(".env");

	emailCfg := EmailConfig{
		SMTPServer:   envFile["SMTP_SERVER"],
		SMTPPort:     587,
		SMTPUsername: envFile["SMTP_USERNAME"],
		SMTPPassword: envFile["SMTP_PASSWORD"],
		FromEmail:    envFile["FROM_EMAIL"],
		ToEmail:      envFile["TO_EMAIL"],
	}

		m := gomail.NewMessage();
		m.SetHeader("From", emailCfg.FromEmail);
		m.SetHeader("To", emailCfg.ToEmail);
		m.SetHeader("Subject", "New Article Submission: "+ article.Title);
		
		body := fmt.Sprintf(`
			<h1>New Article Submission</h1>
			<p><strong>Title:</strong> %s</p>
			<h2>Content:</h2>
		`, article.Title);
		
		m.SetBody("text/html", body);
	
		d := gomail.NewDialer(
			emailCfg.SMTPServer,
			emailCfg.SMTPPort,
			emailCfg.SMTPUsername,
			emailCfg.SMTPPassword,
		);
	
		return d.DialAndSend(m);
}