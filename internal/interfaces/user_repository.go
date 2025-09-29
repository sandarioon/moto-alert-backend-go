package interfaces

import (
	"context"

	"github.com/sandarioon/moto-alert-backend-go/internal/transaction"
	"github.com/sandarioon/moto-alert-backend-go/models"
)

type UserRepository interface {
	CreateUser(ctx context.Context, tx transaction.Transaction, user models.User, code int) (int, error)
	GetUserByEmail(ctx context.Context, tx transaction.Transaction, email string) (models.User, error)
	UpdateUserIsVerified(ctx context.Context, id int, isVerified bool) error
	IsUserExistsWithPhone(ctx context.Context, tx transaction.Transaction, phone string) (bool, error)
	IsUserExistsWithEmail(ctx context.Context, tx transaction.Transaction, email string) (bool, error)
}
