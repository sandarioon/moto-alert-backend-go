package user

import (
	"context"
	"fmt"

	h "github.com/sandarioon/moto-alert-backend-go/internal/helpers"
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
	UpdateLocation(ctx context.Context, tx transaction.Transaction, userId int, input dto.UpdateLocationRequest) (models.User, error)
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

func (s service) UpdateLocation(ctx context.Context, tx transaction.Transaction, userId int, input dto.UpdateLocationRequest) (models.User, error) {
	err := s.userRepo.UpdateUserLocation(ctx, userId, input)
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
		FirstName:     h.NullStringToPtr(u.FirstName),
		LastName:      h.NullStringToPtr(u.LastName),
		Username:      h.NullStringToPtr(u.Username),
		ExpoPushToken: h.NullStringToPtr(u.ExpoPushToken),
		Gender:        string(u.Gender),
		Phone:         h.NullStringToPtr(u.Phone),
		Longitude:     h.NullFloat64ToPtr(u.Longitude),
		Latitude:      h.NullFloat64ToPtr(u.Latitude),
		BikeModel:     h.NullStringToPtr(u.BikeModel),
		MedicalInfo: dto.UserMedicalInfo{
			BloodGroup:      string(u.BloodGroup),
			HeightCm:        h.NullInt16ToPtr(u.HeightCm),
			WeightKg:        h.NullInt16ToPtr(u.WeightKg),
			DateOfBirth:     h.NullTimeToPtr(u.DateOfBirth),
			ChronicDiseases: h.NullStringToPtr(u.ChronicDiseases),
			Allergies:       h.NullStringToPtr(u.Allergies),
			Medications:     h.NullStringToPtr(u.Medications),
			HasHypertension: u.HasHypertension,
			HasHepatitis:    u.HasHepatitis,
			HasHiv:          u.HasHiv,
		},
		GeoUpdatedAt:    h.NullTimeToPtr(u.GeoUpdatedAt),
		CreatedAt:       u.CreatedAt,
		AccidentId:      h.NullInt16ToPtr(u.AccidentId),
		QrCodeUrl:       formatQrCodeUrl(u.Uuid),
		IsQrCodeEnabled: u.IsQrCodeEnabled,
	}
}

func formatQrCodeUrl(uuid string) string {
	switch env := viper.GetString("general.env"); env {
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
