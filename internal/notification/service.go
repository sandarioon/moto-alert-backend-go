package notification

import (
	"context"
	"fmt"

	"github.com/sandarioon/moto-alert-backend-go/internal/transaction"
	"github.com/sandarioon/moto-alert-backend-go/pkg/notifications/email"
	"github.com/sirupsen/logrus"
)

type service struct {
	repo         NotificationRepository
	emailService email.Service
}

type Service interface {
	SendEmail(ctx context.Context, tx transaction.Transaction, templateType string, recipientEmail string, params map[string]string) error
}

func NewService(repo NotificationRepository, emailSvc email.Service) Service {
	return service{repo: repo, emailService: emailSvc}
}

func (s service) SendEmail(ctx context.Context, tx transaction.Transaction, templateType string, recipientEmail string, params map[string]string) error {
	logrus.Infof("Sending email with type %s to %s", templateType, recipientEmail)

	template, err := s.repo.GetEmailTemplateByType(ctx, tx, templateType)
	if err != nil {
		fmt.Errorf("Failed to get email template by type: %v", err)
		return err
	}

	senderName := "Мото ДТП"
	senderEmail := "no-reply@moto-alert.ru"

	err = s.emailService.SendTransactionalEmail(senderName, senderEmail, recipientEmail, template.Title, template.Body, params)
	if err != nil {
		return err
	}

	return nil
}
