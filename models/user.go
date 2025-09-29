package models

import (
	"database/sql"
	"time"
)

type UserGender string
type BloodGroup string

const (
	MALE   UserGender = "MALE"
	FEMALE UserGender = "FEMALE"
)

type User struct {
	Id              int            `json:"id"`
	Code            string         `json:"code"`
	Email           string         `json:"email"`
	HashedPassword  string         `json:"hashedPassword"`
	FirstName       sql.NullString `json:"firstName"`
	LastName        sql.NullString `json:"lastName"`
	Username        sql.NullString `json:"username"`
	ExpoPushToken   sql.NullString `json:"expoPushToken"`
	Gender          UserGender     `json:"gender"`
	Phone           sql.NullString `json:"phone"`
	Longitude       sql.NullString `json:"longitude"`
	Latitude        sql.NullString `json:"latitude"`
	BikeModel       sql.NullString `json:"bikeModel"`
	Comment         sql.NullString `json:"comment"`
	LastAuth        sql.NullTime   `json:"lastAuth"`
	GeoUpdatedAt    sql.NullTime   `json:"geoUpdatedAt"`
	CreatedAt       time.Time      `json:"createdAt"`
	AccidentId      sql.NullInt16  `json:"accidentId"`
	BloodGroup      BloodGroup     `json:"bloodGroup"`
	HeightCm        sql.NullInt16  `json:"heightCm"`
	WeightKg        sql.NullInt16  `json:"weightKg"`
	DateOfBirth     sql.NullTime   `json:"dateOfBirth"`
	ChronicDiseases sql.NullString `json:"chronicDiseases"`
	Allergies       sql.NullString `json:"allergies"`
	Medications     sql.NullString `json:"medications"`
	Geom            sql.NullString `json:"geom"`
	IsBanned        bool           `json:"isBanned"`
	IsVerified      bool           `json:"isVerified"`
	IsDeleted       bool           `json:"isDeleted"`
	Uuid            string         `json:"uuid"`
	IsQrCodeEnabled bool           `json:"isQrCodeEnabled"`
	HasHypertension string         `json:"hasHypertension"`
	HasHepatitis    string         `json:"hasHepatitis"`
	HasHiv          string         `json:"hasHiv"`
}
