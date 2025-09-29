package user

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	_ "github.com/lib/pq"

	"github.com/sandarioon/moto-alert-backend-go/internal/transaction"
	"github.com/sandarioon/moto-alert-backend-go/models"
)

const usersTable = "users"

type userRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *userRepository {
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

	params := []any{
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
	params := []any{strings.ToLower(email)}

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

	// logs.PrintSql(true, query, params...)
	return exists, nil
}

func (r *userRepository) IsUserExistsWithPhone(ctx context.Context, tx transaction.Transaction, phone string) (bool, error) {
	var exists bool

	query := fmt.Sprintf(`SELECT EXISTS(SELECT 1 FROM %s WHERE phone = $1);`, usersTable)
	params := []any{phone}

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

func (r userRepository) GetUserByEmail(email string) (models.User, error) {
	var user models.User
	query := fmt.Sprintf("SELECT * FROM %s WHERE email = ?;", usersTable)

	err := r.db.QueryRow(query, strings.ToLower(email)).Scan(&user)

	if err != nil {
		if err == sql.ErrNoRows {
			return models.User{}, errors.New("user with email " + email + " not found")
		}
		return models.User{}, errors.New("failed to get user by email. Err: " + err.Error())
	}

	return user, nil
}

func (r userRepository) UpdateUserIsVerified(id int, isVerified bool) error {

	query := fmt.Sprintf(`UPDATE %s SET is_verified = ? WHERE id = ?;`, usersTable)

	_, err := r.db.Exec(query, isVerified, id)

	if err != nil {
		return err
	}

	return nil
}
