// package handlers

// import (
// 	"email-service/config"
// 	"fmt"
// 	"net/http"
// 	"net/smtp"

// 	"github.com/gin-gonic/gin"
// )

// // Define the EmailRequest struct to take only "From", "To", "Subject", and "Body"
// type EmailRequest struct {
// 	From    string `json:"from" binding:"required"`
// 	To      string `json:"to" binding:"required"`
// 	Subject string `json:"subject" binding:"required"`
// 	Body    string `json:"body" binding:"required"` // Plain text body from the request
// }

// // Function to send email
// func SendEmail(c *gin.Context) {
// 	var emailReq EmailRequest

// 	// Bind the incoming JSON request to the struct
// 	if err := c.ShouldBindJSON(&emailReq); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// Set up SMTP authentication
// 	auth := smtp.PlainAuth(
// 		"",
// 		config.AppConfig.SMTPUser,
// 		config.AppConfig.SMTPPassword,
// 		config.AppConfig.SMTPHost,
// 	)

// 	// HTML template that wraps only the email body, which is styled with CSS
// 	htmlTemplate := `
// 		<html>
// 		  <body style="background-color: #f3f4f6; margin: 0; padding: 0; font-family: Arial, sans-serif;">
// 		    <div style="max-width: 600px; margin: 20px auto; padding: 20px; background-color: white; border-radius: 8px; box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);">
// 		    	<div style="font-size: 18px; line-height: 1.6;">
// 		    		%s
// 		    	</div>
// 		    </div>
// 		  </body>
// 		</html>
// 	`

// 	// Format the body for proper HTML injection, ensuring the formatting is preserved
// 	formattedBody := formatBodyForHTML(emailReq.Body)

// 	// Inject the formatted body into the HTML template
// 	htmlBody := fmt.Sprintf(htmlTemplate, formattedBody)

// 	// Construct the full email message including the headers and formatted body
// 	message := []byte(fmt.Sprintf(
// 		"From: %s\r\nTo: %s\r\nSubject: %s\r\n"+
// 			"MIME-Version: 1.0\r\n"+
// 			"Content-Type: text/html; charset=\"UTF-8\"\r\n"+
// 			"\r\n%s",
// 		emailReq.From, emailReq.To, emailReq.Subject, htmlBody))

// 	// Send the email via SMTP
// 	err := smtp.SendMail(
// 		fmt.Sprintf("%s:%s", config.AppConfig.SMTPHost, config.AppConfig.SMTPPort),
// 		auth,
// 		emailReq.From,
// 		[]string{emailReq.To},
// 		message,
// 	)

// 	// Handle errors and response
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Failed to send email", "error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Email sent successfully!"})
// }

// // Helper function to format the body for HTML, escaping newlines and special characters
// func formatBodyForHTML(body string) string {
// 	// Wrap the text in <pre> to preserve newlines and any spacing in the text
// 	return fmt.Sprintf("<pre>%s</pre>", body)
// }

package handlers

import (
	"email-service/config"
	"fmt"
	"net/http"
	"net/smtp"

	"github.com/gin-gonic/gin"
)

// Define the EmailRequest struct to take only "From", "To", "Subject", and "Body"
type EmailRequest struct {
	From    string `json:"from" binding:"required"`
	To      string `json:"to" binding:"required"`
	Subject string `json:"subject" binding:"required"`
	Body    string `json:"body" binding:"required"` // Plain text body from the request
}

// Function to send email
func SendEmail(c *gin.Context) {
	var emailReq EmailRequest

	// Bind the incoming JSON request to the struct
	if err := c.ShouldBindJSON(&emailReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set up SMTP authentication
	auth := smtp.PlainAuth(
		"",
		config.AppConfig.SMTPUser,
		config.AppConfig.SMTPPassword,
		config.AppConfig.SMTPHost,
	)

	// Enhanced HTML template with improved CSS styling
	htmlTemplate := `
		<html>
		  <head>
		    <style>
		      body {
		        background-color: #f3f4f6;
		        margin: 0;
		        padding: 0;
		        font-family: Arial, sans-serif;
		      }
		      .container {
		        max-width: 600px;
		        margin: 20px auto;
		        padding: 20px;
		        background-color: #ffffff;
		        border-radius: 8px;
		        box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
		        border-left: 6px solid #4CAF50; /* Green left border */
		      }
		      .content {
		        font-size: 16px;
		        color: #333;
		        line-height: 1.6;
		        white-space: pre-wrap; /* Preserve formatting */
		        padding: 10px;
		        background-color: #f9f9f9; /* Light gray background for the content */
		        border-radius: 5px; /* Rounded corners for content */
		      }
		      h1 {
		        font-size: 24px;
		        color: #4CAF50; /* Green color for the header */
		      }
		    </style>
		  </head>
		  <body>
		    <div class="container">
		      <h1>Email Notification</h1>
		      <div class="content">
		        %s
		      </div>
		    </div>
		  </body>
		</html>
	`

	htmlBody := fmt.Sprintf(htmlTemplate, emailReq.Body)

	// Construct the full email message including headers
	message := []byte(fmt.Sprintf(
		"From: %s\r\nTo: %s\r\nSubject: %s\r\n"+
			"MIME-Version: 1.0\r\n"+
			"Content-Type: text/html; charset=\"UTF-8\"\r\n"+
			"\r\n%s",
		emailReq.From, emailReq.To, emailReq.Subject, htmlBody))

	// Send the email via SMTP
	err := smtp.SendMail(
		fmt.Sprintf("%s:%s", config.AppConfig.SMTPHost, config.AppConfig.SMTPPort),
		auth,
		emailReq.From,
		[]string{emailReq.To},
		message,
	)

	// Handle errors and response
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Failed to send email", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Email sent successfully!"})
}

// package handlers

// import (
// 	"email-service/config"
// 	"fmt"
// 	"net/http"
// 	"net/smtp"

// 	"github.com/gin-gonic/gin"
// )

// // Define the EmailRequest struct to take only "From", "To", "Subject", and "Body"
// type EmailRequest struct {
// 	From    string `json:"from" binding:"required"`
// 	To      string `json:"to" binding:"required"`
// 	Subject string `json:"subject" binding:"required"`
// 	Body    string `json:"body" binding:"required"` // Plain text body from the request
// }

// // Function to send email
// func SendEmail(c *gin.Context) {
// 	var emailReq EmailRequest

// 	// Bind the incoming JSON request to the struct
// 	if err := c.ShouldBindJSON(&emailReq); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// Set up SMTP authentication
// 	auth := smtp.PlainAuth(
// 		"",
// 		config.AppConfig.SMTPUser,
// 		config.AppConfig.SMTPPassword,
// 		config.AppConfig.SMTPHost,
// 	)

// 	// Enhanced HTML template without footer
// 	htmlTemplate := `
// 		<!DOCTYPE html>
// 		<html lang="en">
// 		<head>
// 			<meta charset="UTF-8">
// 			<meta name="viewport" content="width=device-width, initial-scale=1.0">
// 			<title>Email Notification</title>
// 			<style>
// 				body {
// 					background-color: #f0f2f5;
// 					font-family: Arial, sans-serif;
// 					margin: 0;
// 					padding: 0;
// 				}
// 				.container {
// 					max-width: 600px;
// 					margin: 40px auto;
// 					padding: 20px;
// 					background-color: #ffffff;
// 					border-radius: 8px;
// 					box-shadow: 0 4px 20px rgba(0, 0, 0, 0.1);
// 				}
// 				.header {
// 					background-color: #4CAF50;
// 					color: white;
// 					padding: 10px;
// 					text-align: center;
// 					border-radius: 8px 8px 0 0;
// 				}
// 				.content {
// 					font-size: 16px;
// 					color: #333;
// 					line-height: 1.6;
// 					padding: 20px;
// 					border-radius: 0 0 8px 8px;
// 				}
// 				pre {
// 					background-color: #f9f9f9;
// 					border-left: 4px solid #4CAF50;
// 					padding: 10px;
// 					overflow-x: auto;
// 					white-space: pre-wrap; /* Preserve formatting */
// 					border-radius: 4px;
// 				}
// 			</style>
// 		</head>
// 		<body>
// 			<div class="container">
// 				<div class="header">
// 					<h1>Email Notification</h1>
// 				</div>
// 				<div class="content">
// 					<pre>%s</pre>
// 				</div>
// 			</div>
// 		</body>
// 		</html>
// 	`

// 	// Format the body for proper HTML injection, ensuring the formatting is preserved
// 	formattedBody := formatBodyForHTML(emailReq.Body)

// 	// Inject the formatted body into the HTML template
// 	htmlBody := fmt.Sprintf(htmlTemplate, formattedBody)

// 	// Construct the full email message including the headers and formatted body
// 	message := []byte(fmt.Sprintf(
// 		"From: %s\r\nTo: %s\r\nSubject: %s\r\n"+
// 			"MIME-Version: 1.0\r\n"+
// 			"Content-Type: text/html; charset=\"UTF-8\"\r\n"+
// 			"\r\n%s",
// 		emailReq.From, emailReq.To, emailReq.Subject, htmlBody))

// 	// Send the email via SMTP
// 	err := smtp.SendMail(
// 		fmt.Sprintf("%s:%s", config.AppConfig.SMTPHost, config.AppConfig.SMTPPort),
// 		auth,
// 		emailReq.From,
// 		[]string{emailReq.To},
// 		message,
// 	)

// 	// Handle errors and response
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Failed to send email", "error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Email sent successfully!"})
// }

// // Helper function to format the body for HTML, escaping newlines and special characters
// func formatBodyForHTML(body string) string {
// 	// Wrap the text in <pre> to preserve newlines and any spacing in the text
// 	return fmt.Sprintf("<pre>%s</pre>", body)
// }
