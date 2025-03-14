package middleware

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

var (
	signedKey = []byte("secret")
)

type Payload struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func NewPayload(userID string, role string, duration time.Duration) (*Payload, error) {
	usrEmail, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	timeNow := time.Now()
	payload := &Payload{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(timeNow.Add(duration)),
			IssuedAt:  jwt.NewNumericDate(timeNow),
			NotBefore: jwt.NewNumericDate(timeNow),
			Issuer:    "user_login",
			Subject:   "multifinance",
			ID:        usrEmail.String(),
		},
	}
	return payload, nil
}

func CreateRefreshToken(userID string, role string, refreshTokenExpiry time.Duration) (string, *Payload, error) {
	return createToken(userID, role, refreshTokenExpiry)
}

func CreateAccessToken(userID string, role string, tokenExpiry time.Duration) (string, *Payload, error) {
	return createToken(userID, role, tokenExpiry)
}

func createToken(userID string, role string, tokenExpiry time.Duration) (string, *Payload, error) {
	payload, err := NewPayload(userID, role, tokenExpiry)
	if err != nil {
		return "", nil, err
	}
	tokenWithClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	tokens, err := tokenWithClaims.SignedString(signedKey)
	if err != nil {
		return "", nil, err
	}

	return tokens, payload, nil
}

func VerifyToken(tokenString string) (*Payload, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return signedKey, nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token")
	}

	userId := claims["user_id"].(string)
	role := claims["role"].(string)

	payload := &Payload{
		UserID: userId,
		Role:   role,
	}

	return payload, nil
}
