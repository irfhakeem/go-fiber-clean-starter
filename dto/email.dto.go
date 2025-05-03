package dto

type (
	VerificationEmail struct {
		Email            string `json:"email"`
		VerificationLink string `json:"verification_link"`
	}
)
