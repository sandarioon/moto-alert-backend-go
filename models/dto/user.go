package dto

import (
	"time"

	"github.com/sandarioon/moto-alert-backend-go/models"
)

type ProfileResponse struct {
	Status  int          `json:"status" example:"200"`
	Data    UserResponse `json:"data" `
	Message string       `json:"message" example:"OK"`
}

type UserResponse struct {
	Id              int             `json:"id"`
	Email           string          `json:"email"`
	FirstName       *string         `json:"firstName"`
	LastName        *string         `json:"lastName"`
	Username        *string         `json:"username"`
	ExpoPushToken   *string         `json:"expoPushToken"`
	Gender          string          `json:"gender"`
	Phone           *string         `json:"phone"`
	Longitude       *float64        `json:"longitude"`
	Latitude        *float64        `json:"latitude"`
	BikeModel       *string         `json:"bikeModel"`
	GeoUpdatedAt    *time.Time      `json:"geoUpdatedAt"`
	CreatedAt       time.Time       `json:"createdAt"`
	AccidentId      *int16          `json:"accidentId"`
	MedicalInfo     UserMedicalInfo `json:"medicalInfo"`
	IsQrCodeEnabled bool            `json:"isQrCodeEnabled"`
	QrCodeUrl       string          `json:"qrCodeUrl"`
}

type UserMedicalInfo struct {
	BloodGroup      string     `json:"bloodGroup"`
	HeightCm        *int16     `json:"heightCm"`
	WeightKg        *int16     `json:"weightKg"`
	DateOfBirth     *time.Time `json:"dateOfBirth"`
	ChronicDiseases *string    `json:"chronicDiseases"`
	Allergies       *string    `json:"allergies"`
	Medications     *string    `json:"medications"`
	HasHypertension string     `json:"hasHypertension"`
	HasHepatitis    string     `json:"hasHepatitis"`
	HasHiv          string     `json:"hasHiv"`
}

type EditUserRequest struct {
	FirstName *string            `json:"firstName"`
	LastName  *string            `json:"lastName"`
	Username  *string            `json:"username"`
	Phone     *string            `json:"phone"`
	BikeModel *string            `json:"bikeModel"`
	Gender    *models.UserGender `json:"gender"`
}

type UpdateLocationRequest struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

type UpdateExpoPushTokenRequest struct {
	ExpoPushToken string `json:"expoPushToken"`
}
