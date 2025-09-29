package email

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type service struct{}

func NewService() Service {
	return &service{}
}

type Sender struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
type Recipient struct {
	Email string `json:"email"`
}
type Request struct {
	Sender      Sender            `json:"sender"`
	To          []Recipient       `json:"to"`
	Subject     string            `json:"subject"`
	HTMLContent string            `json:"htmlContent"`
	Params      map[string]string `json:"params"`
}

func (s *service) SendTransactionalEmail(senderName string, senderEmail string, recipientEmail string, subject string, htmlContent string, params map[string]string) error {
	url := "https://api.brevo.com/v3/smtp/email"

	payload := Request{
		Sender: Sender{
			Name:  senderName,
			Email: senderEmail,
		},
		To: []Recipient{
			{
				Email: recipientEmail,
			},
		},
		Subject:     subject,
		HTMLContent: htmlContent,
		Params:      params,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal email payload: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("accept", "application/json")
	req.Header.Set("api-key", os.Getenv("BREVO_ACCESS_TOKEN"))
	req.Header.Set("content-type", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	// fmt.Printf("Brevo Response [%d]: %s\n", resp.StatusCode, string(respBody))

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("send email request to brevo failed with status code: %d, err: %s", resp.StatusCode, string(respBody))
	}

	return nil
}
