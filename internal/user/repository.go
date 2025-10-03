package user

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"strings"

	_ "github.com/lib/pq"

	"github.com/sandarioon/moto-alert-backend-go/internal/transaction"
	"github.com/sandarioon/moto-alert-backend-go/models"
	"github.com/sandarioon/moto-alert-backend-go/models/dto"
	postgres "github.com/sandarioon/moto-alert-backend-go/pkg/database"
)

const usersTable = "users"

type userRepository struct {
	db *postgres.DBLogger
}

func NewRepository(db *postgres.DBLogger) *userRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(ctx context.Context, tx transaction.Transaction, user models.User, code int) (int, error) {
	var id int

	query := fmt.Sprintf(
		`INSERT INTO %s (
			email, 
			hashed_password, 
			expo_push_token, 
			username, 
			first_name, 
			last_name, 
			gender, 
			phone, 
			bike_model, 
			latitude, 
			longitude, 
			code, 
			uuid, 
			geom, 
			geo_updated_at) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15) 
		RETURNING id;`,
		usersTable,
	)

	params := []interface{}{
		strings.ToLower(user.Email),
		user.HashedPassword,
		user.ExpoPushToken,
		user.Username,
		user.FirstName,
		user.LastName,
		user.Gender,
		user.Phone,
		user.BikeModel,
		user.Latitude,
		user.Longitude,
		code,
		user.Uuid,
		user.Geom,
		user.GeoUpdatedAt,
	}

	var row *sql.Row
	if tx != nil {
		row = tx.QueryRowContext(ctx, query, params...)
	} else {
		row = r.db.QueryRowContext(ctx, query, params...)
	}

	err := row.Scan(&id)

	if err != nil {
		return id, err
	}

	return id, nil
}

func (r *userRepository) IsUserExistsWithEmail(ctx context.Context, tx transaction.Transaction, email string) (bool, error) {
	var exists bool

	query := fmt.Sprintf(`SELECT EXISTS(SELECT 1 FROM %s WHERE email = $1);`, usersTable)
	params := []interface{}{strings.ToLower(email)}

	var row *sql.Row
	if tx != nil {
		row = tx.QueryRowContext(ctx, query, params...)
	} else {
		row = r.db.QueryRowContext(ctx, query, params...)
	}

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.New("failed to check if user exists with email. Err: " + err.Error())
	}

	return exists, nil
}

func (r *userRepository) IsUserExistsWithPhone(ctx context.Context, tx transaction.Transaction, phone string) (bool, error) {
	var exists bool

	query := fmt.Sprintf(`SELECT EXISTS(SELECT 1 FROM %s WHERE phone = $1);`, usersTable)
	params := []interface{}{phone}

	var row *sql.Row
	if tx != nil {
		row = tx.QueryRowContext(ctx, query, params...)
	} else {
		row = r.db.QueryRowContext(ctx, query, params...)
	}

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.New("failed to check if user exists with phone. Err: " + err.Error())
	}

	return exists, nil
}

func (r userRepository) GetUserById(ctx context.Context, tx transaction.Transaction, id int) (models.User, error) {
	var user models.User

	query := fmt.Sprintf(`SELECT 
		id,
		code,
		email,
		first_name,
		last_name,
		username,
		expo_push_token,
		gender,
		phone,
		longitude,
		latitude,
		bike_model,
		comment,
		last_auth,
		geo_updated_at,
		created_at,
		accident_id,
		blood_group,
		height_cm,
		weight_kg,
		date_of_birth,
		chronic_diseases,
		allergies,
		medications,
		geom,
		is_banned,
		is_verified,
		is_deleted,
		uuid,
		is_qr_code_enabled,
		has_hypertension,
		has_hepatitis,
		has_hiv
	FROM 
		%s
	WHERE 
		id = $1;`, usersTable)
	params := []interface{}{id}

	var row *sql.Row
	if tx != nil {
		row = tx.QueryRowContext(ctx, query, params...)
	} else {
		row = r.db.QueryRowContext(ctx, query, params...)
	}

	err := row.Scan(
		&user.Id,
		&user.Code,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.Username,
		&user.ExpoPushToken,
		&user.Gender,
		&user.Phone,
		&user.Longitude,
		&user.Latitude,
		&user.BikeModel,
		&user.Comment,
		&user.LastAuth,
		&user.GeoUpdatedAt,
		&user.CreatedAt,
		&user.AccidentId,
		&user.BloodGroup,
		&user.HeightCm,
		&user.WeightKg,
		&user.DateOfBirth,
		&user.ChronicDiseases,
		&user.Allergies,
		&user.Medications,
		&user.Geom,
		&user.IsBanned,
		&user.IsVerified,
		&user.IsDeleted,
		&user.Uuid,
		&user.IsQrCodeEnabled,
		&user.HasHypertension,
		&user.HasHepatitis,
		&user.HasHiv,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return models.User{}, errors.New("user with id " + strconv.Itoa(id) + " not found")
		}
		return models.User{}, errors.New("failed to get user by id. Err: " + err.Error())
	}

	return user, nil
}

func (r userRepository) GetUserByEmail(ctx context.Context, tx transaction.Transaction, email string) (models.User, error) {
	var user models.User

	query := fmt.Sprintf(`SELECT 
		id,
		code,
		email,
		first_name,
		last_name,
		username,
		expo_push_token,
		gender,
		phone,
		longitude,
		latitude,
		bike_model,
		comment,
		last_auth,
		geo_updated_at,
		created_at,
		accident_id,
		blood_group,
		height_cm,
		weight_kg,
		date_of_birth,
		chronic_diseases,
		allergies,
		medications,
		geom,
		is_banned,
		is_verified,
		is_deleted,
		uuid,
		is_qr_code_enabled,
		has_hypertension,
		has_hepatitis,
		has_hiv
	FROM 
		%s
	WHERE 
		email = $1;`, usersTable)
	params := []interface{}{email}

	var row *sql.Row
	if tx != nil {
		row = tx.QueryRowContext(ctx, query, params...)
	} else {
		row = r.db.QueryRowContext(ctx, query, params...)
	}

	err := row.Scan(
		&user.Id,
		&user.Code,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.Username,
		&user.ExpoPushToken,
		&user.Gender,
		&user.Phone,
		&user.Longitude,
		&user.Latitude,
		&user.BikeModel,
		&user.Comment,
		&user.LastAuth,
		&user.GeoUpdatedAt,
		&user.CreatedAt,
		&user.AccidentId,
		&user.BloodGroup,
		&user.HeightCm,
		&user.WeightKg,
		&user.DateOfBirth,
		&user.ChronicDiseases,
		&user.Allergies,
		&user.Medications,
		&user.Geom,
		&user.IsBanned,
		&user.IsVerified,
		&user.IsDeleted,
		&user.Uuid,
		&user.IsQrCodeEnabled,
		&user.HasHypertension,
		&user.HasHepatitis,
		&user.HasHiv,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return models.User{}, errors.New("user with email " + email + " not found")
		}
		return models.User{}, errors.New("failed to get user by email. Err: " + err.Error())
	}

	return user, nil
}

func (r userRepository) GetUserHashedPassword(ctx context.Context, tx transaction.Transaction, email string) (string, error) {
	var hashedPassword string

	query := fmt.Sprintf(`SELECT 
		hashed_password
	FROM 
		%s
	WHERE 
		email = $1;`, usersTable)
	params := []interface{}{email}

	var row *sql.Row
	if tx != nil {
		row = tx.QueryRowContext(ctx, query, params...)
	} else {
		row = r.db.QueryRowContext(ctx, query, params...)
	}

	err := row.Scan(&hashedPassword)

	if err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New("hashed password for user with email " + email + " not found")
		}
		return "", errors.New("failed to get hashed password by email. Err: " + err.Error())
	}

	return hashedPassword, nil
}

func (r userRepository) UpdateUserIsVerified(ctx context.Context, id int, isVerified bool) error {

	query := fmt.Sprintf(`UPDATE %s SET is_verified = $1 WHERE id = $2;`, usersTable)

	_, err := r.db.ExecContext(ctx, query, isVerified, id)

	if err != nil {
		return err
	}

	return nil
}

func (r userRepository) UpdateUserPassword(ctx context.Context, email string, hashedPassword string) error {

	query := fmt.Sprintf(`UPDATE %s SET hashed_password = $1 WHERE email = $2;`, usersTable)

	_, err := r.db.ExecContext(ctx, query, hashedPassword, email)

	if err != nil {
		return err
	}

	return nil
}

func (r userRepository) UpdateUserExpoPushToken(ctx context.Context, userId int, expoPushToken *string) error {

	query := fmt.Sprintf(`UPDATE %s SET expo_push_token = $1 WHERE id = $2;`, usersTable)

	_, err := r.db.ExecContext(ctx, query, &expoPushToken, userId)

	if err != nil {
		return err
	}

	return nil
}

func (r userRepository) UpdateUserProfileData(ctx context.Context, userId int, input dto.EditUserRequest) error {
	var (
		queryParts []string
		args       []interface{}
		argIndex   = 1
	)

	// TODO
	// Нужно добавить сюда валидации на поля
	if input.FirstName != nil {
		queryParts = append(queryParts, fmt.Sprintf("first_name = $%d", argIndex))
		args = append(args, *input.FirstName)
		argIndex++
	}
	if input.LastName != nil {
		queryParts = append(queryParts, fmt.Sprintf("last_name = $%d", argIndex))
		args = append(args, *input.LastName)
		argIndex++
	}
	if input.Username != nil {
		queryParts = append(queryParts, fmt.Sprintf("username = $%d", argIndex))
		args = append(args, *input.Username)
		argIndex++
	}
	if input.Phone != nil {
		queryParts = append(queryParts, fmt.Sprintf("phone = $%d", argIndex))
		args = append(args, *input.Phone)
		argIndex++
	}
	if input.BikeModel != nil {
		queryParts = append(queryParts, fmt.Sprintf("bike_model = $%d", argIndex))
		args = append(args, *input.BikeModel)
		argIndex++
	}
	if input.Gender != nil {
		queryParts = append(queryParts, fmt.Sprintf("gender = $%d", argIndex))
		args = append(args, *input.Gender)
		argIndex++
	}

	if len(queryParts) == 0 {
		return nil
	}

	query := fmt.Sprintf(`UPDATE %s SET %s WHERE id = $%d`, usersTable, strings.Join(queryParts, ", "), argIndex)
	args = append(args, userId)

	result, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return errors.New("failed to update user: " + err.Error())
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return errors.New("failed to check affected rows: " + err.Error())
	}
	if rowsAffected == 0 {
		return fmt.Errorf("user with id %d not found", userId)
	}

	return nil
}

func (r userRepository) UpdateUserLocation(ctx context.Context, userId int, input dto.UpdateLocationRequest) error {

	query := fmt.Sprintf(`UPDATE %s SET longitude = $1, latitude = $2 WHERE id = $3;`, usersTable)

	_, err := r.db.ExecContext(ctx, query, input.Longitude, input.Latitude, userId)

	if err != nil {
		return err
	}

	return nil
}
