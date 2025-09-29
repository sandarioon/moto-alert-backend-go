package dto

import "github.com/sandarioon/moto-alert-backend-go/models"

type CreateUserRequest struct {
	Email         string            `json:"email" binding:"required"`
	Password      string            `json:"password" binding:"required,min=8,max=25"`
	ExpoPushToken *string           `json:"expoPushToken"`
	Username      *string           `json:"username"`
	FirstName     *string           `json:"firstName"`
	LastName      *string           `json:"lastName"`
	Gender        models.UserGender `json:"gender"`
	Phone         *string           `json:"phone"`
	BikeModel     *string           `json:"bikeModel"`
	Latitude      *float32          `json:"latitude"`
	Longitude     *float32          `json:"longitude"`
}

type VerifyCodeRequest struct {
	Email string `json:"email" binding:"required"`
	Code  string `json:"code" binding:"required,min=6,max=6"`
}
type VerifyCodeResponse struct {
	Status  int      `json:"status" example:"200"`
	Data    JwtToken `json:"data"`
	Message string   `json:"message" example:"OK"`
}

type VerifyEmailRequest struct {
	Email string `json:"email" binding:"required"`
}

type ResendCodeRequest struct {
	Email string `json:"email" binding:"required"`
}

type ForgotPasswordRequest struct {
	Email string `json:"email" binding:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type LoginResponse struct {
	Status  int      `json:"status" example:"200"`
	Data    JwtToken `json:"data"`
	Message string   `json:"message" example:"OK"`
}

type JwtToken struct {
	Token string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"`
}
