package auth

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"math/rand/v2"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/sandarioon/moto-alert-backend-go/internal/helpers"
	"github.com/sandarioon/moto-alert-backend-go/internal/interfaces"
	"github.com/sandarioon/moto-alert-backend-go/internal/notification"
	"github.com/sandarioon/moto-alert-backend-go/internal/transaction"
	"github.com/sandarioon/moto-alert-backend-go/models"
	"github.com/sandarioon/moto-alert-backend-go/models/dto"
	"golang.org/x/crypto/bcrypt"
)

type service struct {
	transactioner   transaction.Transactioner
	userRepo        interfaces.UserRepository
	notificationSvc notification.Service
}

type Service interface {
	CreateUser(ctx context.Context, input dto.CreateUserRequest) (int, error)
	VerifyCode(ctx context.Context, input dto.VerifyCodeRequest) (string, error)
	VerifyEmail(ctx context.Context, input dto.VerifyEmailRequest) error
	ForgotPassword(ctx context.Context, input dto.ForgotPasswordRequest) error
	ResendCode(ctx context.Context, input dto.ResendCodeRequest) error
	Login(ctx context.Context, input dto.LoginRequest) (string, error)
	Logout(ctx context.Context, userId int) error
}

func NewService(t transaction.Transactioner, userRepo interfaces.UserRepository, notificationSvc notification.Service) Service {
	return service{transactioner: t, userRepo: userRepo, notificationSvc: notificationSvc}
}

func (s service) CreateUser(ctx context.Context, input dto.CreateUserRequest) (int, error) {
	var id int

	tx, err := s.transactioner.BeginTx(ctx)
	if err != nil {
		return 0, err
	}

	defer tx.Rollback()

	existsWithEmail, err := s.userRepo.IsUserExistsWithEmail(ctx, tx, input.Email)
	if err != nil {
		return 0, err
	}
	if existsWithEmail {
		return 0, errors.New("Пользователь с таким email уже существует")
	}

	if input.Phone != nil {
		existsWithPhone, err := s.userRepo.IsUserExistsWithPhone(ctx, tx, *input.Phone)
		if err != nil {
			return 0, err
		}
		if existsWithPhone {
			return 0, errors.New("Пользователь с таким телефоном уже существует")
		}
	}

	hashedPassword, err := s.hashPassword(input.Password)
	if err != nil {
		return 0, err
	}

	code, err := s.generateCode()
	if err != nil {
		return 0, err
	}

	var geom = sql.NullString{}
	var geoUpdatedAt = sql.NullTime{}
	if input.Longitude != nil && input.Latitude != nil {
		geom.Valid = true
		geom.String = helpers.CreateGeom(*input.Longitude, *input.Latitude)

		geoUpdatedAt.Valid = true
		geoUpdatedAt.Time = time.Now().UTC()
	} else {
		geom.Valid = false
		geoUpdatedAt.Valid = false
	}

	var expoPushToken = sql.NullString{}
	if input.ExpoPushToken != nil {
		expoPushToken.Valid = true
		expoPushToken.String = *input.ExpoPushToken
	}

	var username = sql.NullString{}
	if input.Username != nil {
		username.Valid = true
		username.String = *input.Username
	}

	var firstName = sql.NullString{}
	if input.FirstName != nil {
		firstName.Valid = true
		firstName.String = *input.FirstName
	}

	var lastName = sql.NullString{}
	if input.LastName != nil {
		lastName.Valid = true
		lastName.String = *input.LastName
	}

	var phone = sql.NullString{}
	if input.Phone != nil {
		phone.Valid = true
		phone.String = *input.Phone
	}

	var bikeModel = sql.NullString{}
	if input.BikeModel != nil {
		bikeModel.Valid = true
		bikeModel.String = *input.BikeModel
	}

	user := models.User{
		Email:          input.Email,
		HashedPassword: hashedPassword,
		ExpoPushToken:  expoPushToken,
		Username:       username,
		FirstName:      firstName,
		LastName:       lastName,
		Gender:         input.Gender,
		Phone:          phone,
		BikeModel:      bikeModel,
		Uuid:           uuid.New().String(),
	}

	id, err = s.userRepo.CreateUser(ctx, tx, user, code)
	if err != nil {
		return 0, err
	}

	emailParams := map[string]string{"code": fmt.Sprintf("%d", code)}

	err = s.notificationSvc.SendEmail(ctx, tx, string(models.SEND_CODE), input.Email, emailParams)
	if err != nil {
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s service) hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Error hashing password:", err)
		return "", err
	}

	return string(hashedPassword), nil
}

func (s service) checkPassword(hashedPassword, password string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		fmt.Println("Password comparison failed:", err)
		return false, err
	}

	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) == nil, nil
}

func (s service) generateCode() (int, error) {
	min := 100000
	max := 999999
	randomNumber := rand.IntN(max-min+1) + min
	return randomNumber, nil
}

func (s service) VerifyCode(ctx context.Context, input dto.VerifyCodeRequest) (string, error) {
	var token string

	user, err := s.userRepo.GetUserByEmail(ctx, nil, input.Email)
	if err != nil {
		return "", errors.New("user not found")
	}

	if user.IsVerified {
		return "", errors.New("user already verified")
	}

	if user.Code != input.Code {
		return "", errors.New("invalid code")
	}

	token, err = s.generateJwtToken(user.Id)
	if err != nil {
		return "", err
	}

	err = s.userRepo.UpdateUserIsVerified(ctx, user.Id, true)
	if err != nil {
		return "", err

	}

	return token, nil
}

func (s service) VerifyEmail(ctx context.Context, input dto.VerifyEmailRequest) error {
	exists, err := s.userRepo.IsUserExistsWithEmail(ctx, nil, input.Email)

	if exists {
		return errors.New("user already exists")
	}
	if err != nil {
		return err
	}

	return nil
}

func (s service) ForgotPassword(ctx context.Context, input dto.ForgotPasswordRequest) error {
	exists, err := s.userRepo.IsUserExistsWithEmail(ctx, nil, input.Email)

	if !exists {
		return errors.New("user not found")
	}
	if err != nil {
		return err
	}

	newPass, err := helpers.GenerateRandomString(12)
	if err != nil {
		return errors.New("failed to generate random string. Err: " + err.Error())
	}

	hashedPassword, err := s.hashPassword(newPass)
	if err != nil {
		return errors.New("failed to hash password. Err: " + err.Error())
	}

	err = s.userRepo.UpdateUserPassword(ctx, input.Email, hashedPassword)
	if err != nil {
		return errors.New("failed to update user password. Err: " + err.Error())
	}

	err = s.notificationSvc.SendEmail(ctx, nil, string(models.SEND_NEW_PASSWORD), input.Email, map[string]string{
		"password": newPass,
	})
	if err != nil {
		return errors.New("failed to send email. Err: " + err.Error())
	}

	return nil
}

func (s service) generateJwtToken(userId int) (string, error) {
	now := time.Now()

	claims := Claims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(time.Duration(time.Hour * 12))),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func (s service) ResendCode(ctx context.Context, input dto.ResendCodeRequest) error {
	user, err := s.userRepo.GetUserByEmail(ctx, nil, input.Email)
	if err != nil {
		return err
	}
	if user.Id == 0 {
		return errors.New("Пользователь с таким email уже существует")
	}

	emailParams := map[string]string{"code": user.Code}

	err = s.notificationSvc.SendEmail(ctx, nil, string(models.SEND_CODE), input.Email, emailParams)
	if err != nil {
		return err
	}

	return nil
}

func (s service) Login(ctx context.Context, input dto.LoginRequest) (string, error) {
	var token string
	user, err := s.userRepo.GetUserByEmail(ctx, nil, input.Email)
	if err != nil {
		return "", err
	}
	if user.Id == 0 {
		return "", errors.New("user not found")
	}

	if !user.IsVerified {
		return "", errors.New("user not verified")
	}

	hashedPassword, err := s.userRepo.GetUserHashedPassword(ctx, nil, input.Email)
	if err != nil {
		return "", err
	}

	isValid, err := s.checkPassword(hashedPassword, input.Password)
	if err != nil {
		return "", errors.New("failed to check password")
	}

	if !isValid {
		return "", errors.New("invalid password")
	}

	token, err = s.generateJwtToken(user.Id)
	if err != nil {
		return "", errors.New("failed to create jwt token. Err: " + err.Error())
	}

	return token, nil
}

func (s service) Logout(ctx context.Context, userId int) error {
	err := s.userRepo.UpdateUserExpoPushToken(ctx, userId, nil)
	if err != nil {
		return err
	}

	return nil
}
