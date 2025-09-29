package notification

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/sandarioon/moto-alert-backend-go/internal/transaction"
	"github.com/sandarioon/moto-alert-backend-go/models"
	postgres "github.com/sandarioon/moto-alert-backend-go/pkg/database"
)

const emailTemplatesTable = "email_templates"
const pushNotificationTemplatesTable = "push_notification_templates"

type NotificationRepository interface {
	GetEmailTemplateByType(ctx context.Context, tx transaction.Transaction, templateType string) (models.EmailTemplate, error)
}

type repository struct {
	db *postgres.DBLogger
}

func NewRepository(db *postgres.DBLogger) *repository {
	return &repository{db: db}
}

func (r repository) GetEmailTemplateByType(ctx context.Context, tx transaction.Transaction, templateType string) (models.EmailTemplate, error) {
	var template models.EmailTemplate

	query := fmt.Sprintf("SELECT * FROM %s WHERE type = $1;", emailTemplatesTable)
	params := []any{templateType}

	var row *sql.Row
	if tx != nil {
		row = tx.QueryRowContext(ctx, query, params...)
	} else {
		row = r.db.QueryRowContext(ctx, query, params...)
	}

	err := row.Scan(&template.Id, &template.Type, &template.Title, &template.Body)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.EmailTemplate{}, errors.New(fmt.Sprintf("failed to find email template by type %s. Err: ", templateType) + err.Error())
		}
		return models.EmailTemplate{}, err
	}

	return template, nil
}
