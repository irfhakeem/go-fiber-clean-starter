package service

import (
	"context"
	"encoding/json"

	"github.com/irfhakeem/go-fiber-clean-starter/dto"
	"github.com/irfhakeem/go-fiber-clean-starter/rabbitmq"
)

func SendVerificationEmail(ctx context.Context, email string, verificationLink string) error {
	payload := dto.VerificationEmail{
		Email:            email,
		VerificationLink: verificationLink,
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	rabbitmq.PublishEmail("email_queue", jsonPayload)
	return nil
}
