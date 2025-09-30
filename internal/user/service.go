package user

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/sandarioon/moto-alert-backend-go/internal/interfaces"
	"github.com/sandarioon/moto-alert-backend-go/internal/transaction"
	"github.com/sandarioon/moto-alert-backend-go/models"
	"github.com/sandarioon/moto-alert-backend-go/models/dto"
	"github.com/spf13/viper"
)

type service struct {
	transactioner transaction.Transactioner
	userRepo      interfaces.UserRepository
}

type Service interface {
	GetProfile(ctx context.Context, tx transaction.Transaction, userId int) (models.User, error)
	EditUser(ctx context.Context, tx transaction.Transaction, userId int, input dto.EditUserRequest) (models.User, error)
}

func NewService(t transaction.Transactioner, userRepo interfaces.UserRepository) Service {
	return service{transactioner: t, userRepo: userRepo}
}

func (s service) GetProfile(ctx context.Context, tx transaction.Transaction, userId int) (models.User, error) {
	return s.userRepo.GetUserById(ctx, tx, userId)
}

func (s service) EditUser(ctx context.Context, tx transaction.Transaction, userId int, input dto.EditUserRequest) (models.User, error) {
	err := s.userRepo.UpdateUserProfileData(ctx, userId, input)
	if err != nil {
		return models.User{}, err
	}

	user, err := s.userRepo.GetUserById(ctx, tx, userId)
	if err != nil {
		return models.User{}, err
	}

	return user, err
}

func FormatUser(u models.User) dto.UserResponse {
	return dto.UserResponse{
		Id:            u.Id,
		Email:         u.Email,
		FirstName:     nullStringToPtr(u.FirstName),
		LastName:      nullStringToPtr(u.LastName),
		Username:      nullStringToPtr(u.Username),
		ExpoPushToken: nullStringToPtr(u.ExpoPushToken),
		Gender:        string(u.Gender),
		Phone:         nullStringToPtr(u.Phone),
		Longitude:     nullStringToPtr(u.Longitude),
		Latitude:      nullStringToPtr(u.Latitude),
		BikeModel:     nullStringToPtr(u.BikeModel),
		MedicalInfo: dto.UserMedicalInfo{
			BloodGroup:      string(u.BloodGroup),
			HeightCm:        nullInt16ToPtr(u.HeightCm),
			WeightKg:        nullInt16ToPtr(u.WeightKg),
			DateOfBirth:     nullTimeToPtr(u.DateOfBirth),
			ChronicDiseases: nullStringToPtr(u.ChronicDiseases),
			Allergies:       nullStringToPtr(u.Allergies),
			Medications:     nullStringToPtr(u.Medications),
			HasHypertension: u.HasHypertension,
			HasHepatitis:    u.HasHepatitis,
			HasHiv:          u.HasHiv,
		},
		GeoUpdatedAt:    nullTimeToPtr(u.GeoUpdatedAt),
		CreatedAt:       u.CreatedAt,
		AccidentId:      nullInt16ToPtr(u.AccidentId),
		QrCodeUrl:       formatQrCodeUrl(u.Uuid),
		IsQrCodeEnabled: u.IsQrCodeEnabled,
	}
}

func formatQrCodeUrl(uuid string) string {
	env := viper.GetString("general.env")
	switch env {
	case "local":
		return fmt.Sprintf("http://localhost:%s/user/publicProfile/%s", viper.GetString("general.port"), uuid)
	case "production":
		return fmt.Sprintf("https://production.moto-alert.ru/user/publicProfile/%s", uuid)
	case "stage":
		return fmt.Sprintf("https://stage.moto-alert.ru/user/publicProfile/%s", uuid)
	default:
		return fmt.Sprintf("https://production.moto-alert.ru/user/publicProfile/%s", uuid)
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
