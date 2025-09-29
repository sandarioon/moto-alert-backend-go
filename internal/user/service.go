package user

import (
	"context"
	"database/sql"
	"time"

	"github.com/sandarioon/moto-alert-backend-go/internal/interfaces"
	"github.com/sandarioon/moto-alert-backend-go/internal/transaction"
	"github.com/sandarioon/moto-alert-backend-go/models"
	"github.com/sandarioon/moto-alert-backend-go/models/dto"
)

type service struct {
	transactioner transaction.Transactioner
	userRepo      interfaces.UserRepository
}

type Service interface {
	GetProfile(ctx context.Context, tx transaction.Transaction, id int) (models.User, error)
}

func NewService(t transaction.Transactioner, userRepo interfaces.UserRepository) Service {
	return service{transactioner: t, userRepo: userRepo}
}

func (s service) GetProfile(ctx context.Context, tx transaction.Transaction, id int) (models.User, error) {
	return s.userRepo.GetUserById(ctx, tx, id)
}

func FormatUser(u models.User) dto.UserResponse {
	return dto.UserResponse{
		Id:              u.Id,
		Code:            u.Code,
		Email:           u.Email,
		FirstName:       nullStringToPtr(u.FirstName),
		LastName:        nullStringToPtr(u.LastName),
		Username:        nullStringToPtr(u.Username),
		ExpoPushToken:   nullStringToPtr(u.ExpoPushToken),
		Gender:          string(u.Gender),
		Phone:           nullStringToPtr(u.Phone),
		Longitude:       nullStringToPtr(u.Longitude),
		Latitude:        nullStringToPtr(u.Latitude),
		BikeModel:       nullStringToPtr(u.BikeModel),
		Comment:         nullStringToPtr(u.Comment),
		LastAuth:        nullTimeToPtr(u.LastAuth),
		GeoUpdatedAt:    nullTimeToPtr(u.GeoUpdatedAt),
		CreatedAt:       u.CreatedAt,
		AccidentId:      nullInt16ToPtr(u.AccidentId),
		BloodGroup:      string(u.BloodGroup),
		HeightCm:        nullInt16ToPtr(u.HeightCm),
		WeightKg:        nullInt16ToPtr(u.WeightKg),
		DateOfBirth:     nullTimeToPtr(u.DateOfBirth),
		ChronicDiseases: nullStringToPtr(u.ChronicDiseases),
		Allergies:       nullStringToPtr(u.Allergies),
		Medications:     nullStringToPtr(u.Medications),
		IsBanned:        u.IsBanned,
		IsVerified:      u.IsVerified,
		Uuid:            u.Uuid,
		IsQrCodeEnabled: u.IsQrCodeEnabled,
		HasHypertension: u.HasHypertension,
		HasHepatitis:    u.HasHepatitis,
		HasHiv:          u.HasHiv,
	}
}

func nullStringToPtr(ns sql.NullString) *string {
	if ns.Valid {
		return &ns.String
	}
	return nil
}

func nullInt16ToPtr(n sql.NullInt16) *int16 {
	if n.Valid {
		return &n.Int16
	}
	return nil
}

func nullTimeToPtr(nt sql.NullTime) *time.Time {
	if nt.Valid {
		return &nt.Time
	}
	return nil
}
