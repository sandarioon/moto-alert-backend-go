package interfaces

import (
	"context"

	"github.com/sandarioon/moto-alert-backend-go/internal/transaction"
	"github.com/sandarioon/moto-alert-backend-go/models"
	"github.com/sandarioon/moto-alert-backend-go/models/dto"
)

type UserRepository interface {
	CreateUser(ctx context.Context, tx transaction.Transaction, user models.User, code int) (int, error)
	GetUserById(ctx context.Context, tx transaction.Transaction, id int) (models.User, error)
	GetUserByEmail(ctx context.Context, tx transaction.Transaction, email string) (models.User, error)
	UpdateUserIsVerified(ctx context.Context, id int, isVerified bool) error
	IsUserExistsWithPhone(ctx context.Context, tx transaction.Transaction, phone string) (bool, error)
	IsUserExistsWithEmail(ctx context.Context, tx transaction.Transaction, email string) (bool, error)
	UpdateUserPassword(ctx context.Context, email string, hashedPassword string) error
	UpdateUserExpoPushToken(ctx context.Context, userId int, expoPushToken *string) error
	UpdateUserProfileData(ctx context.Context, userId int, input dto.EditUserRequest) error
	GetUserHashedPassword(ctx context.Context, tx transaction.Transaction, email string) (string, error)
}
