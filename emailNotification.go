package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/url"

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

func sendEmail(article types.IArticleText) error {
	// Load environment variables
	envFile, err := godotenv.Read(".env")
	if err != nil {
		return fmt.Errorf("error loading .env file: %v", err)
	}

	// Set up email configuration
	emailCfg := EmailConfig{
		SMTPServer:   envFile["SMTP_SERVER"],
		SMTPPort:     587,
		SMTPUsername: envFile["SMTP_USERNAME"],
		SMTPPassword: envFile["SMTP_PASSWORD"],
		FromEmail:    envFile["FROM_EMAIL"],
		ToEmail:      envFile["TO_EMAIL"],
	}

	// Get validation token from env
	authToken := envFile["VALIDATION_TOKEN"]

	// Get base URL (default if not specified)
	baseURL := envFile["API_BASE_URL"]
	if baseURL == "" {
		baseURL = "http://localhost:8080" // Default for local development
	}

	// Create article JSON for URL parameters
	articleBytes, err := json.Marshal(article)
	if err != nil {
		return fmt.Errorf("error serializing article: %v", err)
	}
	encodedArticle := base64.StdEncoding.EncodeToString(articleBytes)

	// Create validation URL with parameters
	validationURL := fmt.Sprintf("%s/api/validatePosting?data=%s&auth=%s",
		baseURL, url.QueryEscape(encodedArticle), url.QueryEscape(authToken))

	// Create a new message
	m := gomail.NewMessage()
	m.SetHeader("From", emailCfg.FromEmail)
	m.SetHeader("To", emailCfg.ToEmail)
	m.SetHeader("Subject", "New Article Submission: "+article.Title)

	// Create email body with a button-styled link
	body := fmt.Sprintf(`
		<html>
		<head>
			<style>
				.button {
					display: inline-block;
					background-color: #4CAF50;
					color: white;
					padding: 12px 24px;
					text-align: center;
					text-decoration: none;
					font-size: 16px;
					margin: 10px 0;
					cursor: pointer;
					border-radius: 4px;
					font-weight: bold;
				}
				.article-details {
					background-color: #f8f8f8;
					border-left: 4px solid #ddd;
					padding: 12px;
					margin: 15px 0;
				}
			</style>
		</head>
		<body>
			<h1>New Article Submission</h1>
			<div class="article-details">
				<p><strong>Title:</strong> %s</p>
				<h2>Content Preview:</h2>
				<p>%s</p>
			</div>
			<p>Review this submission and click the button below to validate and publish:</p>
			<a href="%s" class="button">Validate & Publish Article</a>
			<p>Or copy and paste this URL into your browser:</p>
			<p><small>%s</small></p>
		</body>
		</html>
	`, article.Title, truncateContent(article.Content, 200), validationURL, validationURL)

	m.SetBody("text/html", body)

	// Set up email dialer
	d := gomail.NewDialer(
		emailCfg.SMTPServer,
		emailCfg.SMTPPort,
		emailCfg.SMTPUsername,
		emailCfg.SMTPPassword,
	)

	// Send the email
	return d.DialAndSend(m)
}

// Helper function to truncate content for preview
func truncateContent(content string, maxLen int) string {
	if len(content) <= maxLen {
		return content
	}
	return content[:maxLen] + "..."
}