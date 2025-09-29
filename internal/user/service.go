package user

import (
	"github.com/sandarioon/moto-alert-backend-go/internal/interfaces"
	"github.com/sandarioon/moto-alert-backend-go/internal/transaction"
)

type service struct {
	transactioner transaction.Transactioner
	userRepo      interfaces.UserRepository
}

type Service interface {
}

func NewService(t transaction.Transactioner, userRepo interfaces.UserRepository) Service {
	return service{transactioner: t, userRepo: userRepo}
}
