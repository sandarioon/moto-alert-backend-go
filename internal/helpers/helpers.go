package helpers

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"

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
