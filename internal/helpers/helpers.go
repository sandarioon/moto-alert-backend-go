package helpers

import (
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"errors"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateGeom(longitude, latitude float32) string {
	// SRID=4326;POINT(20.562519027 54.733065521)

	return fmt.Sprintf("SRID=4326;POINT(%f %f)", longitude, latitude)
}

func GetContextUserId(c *gin.Context) (int, error) {
	value, exists := c.Get("user_id")
	if !exists {
		return 0, errors.New("user id not found")
	}

	return value.(int), nil
}

func GenerateRandomString(length int) (string, error) {
	// Create a byte slice of the specified length
	bytes := make([]byte, length)

	// Read cryptographically secure random bytes into the slice
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	// Encode the random bytes to a base64 string
	return base64.URLEncoding.EncodeToString(bytes), nil
}

func NullStringToPtr(ns sql.NullString) *string {
	if ns.Valid {
		return &ns.String
	}
	return nil
}

func NullInt16ToPtr(n sql.NullInt16) *int16 {
	if n.Valid {
		return &n.Int16
	}
	return nil
}

func NullFloat64ToPtr(n sql.NullFloat64) *float64 {
	if n.Valid {
		return &n.Float64
	}
	return nil
}

func NullTimeToPtr(nt sql.NullTime) *time.Time {
	if nt.Valid {
		return &nt.Time
	}
	return nil
}
