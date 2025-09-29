package dto

import (
	"time"
)

type ProfileResponse struct {
	Status  int          `json:"status" example:"200"`
	Data    UserResponse `json:"data" `
	Message string       `json:"message" example:"OK"`
}

type UserResponse struct {
	Id              int        `json:"id"`
	Code            string     `json:"code"`
	Email           string     `json:"email"`
	FirstName       *string    `json:"firstName,omitempty"`
	LastName        *string    `json:"lastName,omitempty"`
	Username        *string    `json:"username,omitempty"`
	ExpoPushToken   *string    `json:"expoPushToken,omitempty"`
	Gender          string     `json:"gender"`
	Phone           *string    `json:"phone,omitempty"`
	Longitude       *string    `json:"longitude,omitempty"`
	Latitude        *string    `json:"latitude,omitempty"`
	BikeModel       *string    `json:"bikeModel,omitempty"`
	Comment         *string    `json:"comment,omitempty"`
	LastAuth        *time.Time `json:"lastAuth,omitempty"`
	GeoUpdatedAt    *time.Time `json:"geoUpdatedAt,omitempty"`
	CreatedAt       time.Time  `json:"createdAt"`
	AccidentId      *int16     `json:"accidentId,omitempty"`
	BloodGroup      string     `json:"bloodGroup"`
	HeightCm        *int16     `json:"heightCm,omitempty"`
	WeightKg        *int16     `json:"weightKg,omitempty"`
	DateOfBirth     *time.Time `json:"dateOfBirth,omitempty"`
	ChronicDiseases *string    `json:"chronicDiseases,omitempty"`
	Allergies       *string    `json:"allergies,omitempty"`
	Medications     *string    `json:"medications,omitempty"`
	IsBanned        bool       `json:"isBanned"`
	IsVerified      bool       `json:"isVerified"`
	Uuid            string     `json:"uuid"`
	IsQrCodeEnabled bool       `json:"isQrCodeEnabled"`
	HasHypertension string     `json:"hasHypertension"`
	HasHepatitis    string     `json:"hasHepatitis"`
	HasHiv          string     `json:"hasHiv"`
}
