package consumers

import (
	"bytes"
	"encoding/json"
	"log"
	"os"
	"text/template"

	"github.com/irfhakeem/go-fiber-clean-starter/dto"
	"github.com/irfhakeem/go-fiber-clean-starter/helpers/utils"
)

func HandleVerificationEmail(msg []byte) {
	var payload dto.VerificationEmail
	if err := json.Unmarshal(msg, &payload); err != nil {
		log.Printf("Failed to unmarshal email payload: %v", err)
		return
	}

	readHtml, err := os.ReadFile("helpers/utils/email-templates/verification-email.html")
	if err != nil {
		log.Printf("Failed to read email template: %v", err)
		return
	}

	tmpl, err := template.New("custom").Parse(string(readHtml))
	if err != nil {
		log.Printf("Failed to parse template: %v", err)
		return
	}

	var strMail bytes.Buffer
	if err := tmpl.Execute(&strMail, payload); err != nil {
		log.Printf("Failed to execute template: %v", err)
		return
	}

	if err := utils.SendEmail(payload.Email, strMail.String(), "Go Fiber Template - Verify Email"); err != nil {
		log.Printf("Failed to send email: %v", err)
		return
	}

	log.Printf("Verification email sent to %s", payload.Email)
}
