package model

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Users struct {
	ID           uuid.UUID    `json:"id"`
	NIK          string       `json:"nik" validate:"required,len=16,numeric"`
	Fullname     string       `json:"fullname" validate:"required,min=3"`
	LegalName    string       `json:"legal_name" validate:"required"`
	PlaceOfBirth string       `json:"place_of_birth" validate:"required"`
	DateOfBirth  string       `json:"date_of_birth" validate:"required"`
	Salary       float64      `json:"salary" validate:"required,gt=0"`
	Password     string       `json:"password" validate:"required,min=8"`
	Role         string       `json:"role" validate:"required,oneof=User Admin"`
	CreatedAt    time.Time    `json:"created_at"`
	UpdatedAt    sql.NullTime `json:"updated_at"`
}

type UserDocument struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	DocType   string    `json:"doc_type"`
	DocURL    string    `json:"doc_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserLimit struct {
	ID          uuid.UUID `json:"id"`
	UserID      uuid.UUID `json:"user_id"`
	Tenor       int       `json:"tenor"`
	LimitAmount float64   `json:"limit_amount"`
	UsedAmount  float64   `json:"used_amount"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type UserLoginRequest struct {
	Nik      string `json:"nik" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserLogin struct {
	AccessToken          string    `json:"access_token"`
	AccessTokenExpiresAt time.Time `json:"access_token_expires_at"`
	RefreshToken         string    `json:"refresh_token"`
	RefreshTokenExpiryAt time.Time `json:"refresh_token_expiry_at"`
}
