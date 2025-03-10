package users

import (
	"errors"
	"time"

	model "github.com/RamadanRangkuti/multifinance-api/models"
	"github.com/RamadanRangkuti/multifinance-api/repository/users"
	"github.com/RamadanRangkuti/multifinance-api/util/middleware"
	"github.com/google/uuid"
)

type svc struct {
	userStore users.UserRepository
}

func NewUserSvc(userStore users.UserRepository) *svc {
	return &svc{
		userStore: userStore,
	}
}

type UserSvc interface {
	UserRegister(req model.Users) (*uuid.UUID, error)
	UserLogin(req model.UserLoginRequest) (*model.UserLogin, error)
}

func (s *svc) UserRegister(req model.Users) (*uuid.UUID, error) {
	user, err := s.userStore.GetUserDetail(req)
	if err != nil {
		return nil, err
	}

	//Validasi NIK
	if user.NIK == req.NIK {
		return nil, errors.Join(errors.New("user already exists"))
	}

	// menghashing password
	salt, err := middleware.GenerateSalt(16)
	if err != nil {
		return nil, err
	}

	isPassword, err := middleware.HashPassword(req.Password, salt)
	if err != nil {
		return nil, err
	}

	req.Password = isPassword

	userID, err := s.userStore.UserRegister(req)
	if err != nil {
		return nil, err
	}

	return userID, nil
}

func (s *svc) UserLogin(req model.UserLoginRequest) (*model.UserLogin, error) {
	user, err := s.userStore.GetUserDetail(model.Users{
		NIK: req.Nik,
	})
	if err != nil {
		return nil, err
	}

	if user.NIK != req.Nik {
		return nil, errors.Join(errors.New("invalid NIK or Password"))
	}

	verifyPassword, err := middleware.VerifyPassword(req.Password, user.Password)
	if err != nil {
		return nil, err
	}

	if !verifyPassword {
		return nil, errors.Join(errors.New("invalid NIK or Password"))
	}

	tokenExpiry := time.Minute * 20
	accessToken, payload, err := middleware.CreateAccessToken(user.ID.String(), user.Role, tokenExpiry)
	if err != nil {
		return nil, err
	}

	refreshTokenExpiry := time.Hour * 72
	refreshToken, refreshTokenPayload, err := middleware.CreateRefreshToken(user.ID.String(), user.Role, refreshTokenExpiry)
	if err != nil {
		return nil, err
	}

	return &model.UserLogin{
		AccessToken:          accessToken,
		AccessTokenExpiresAt: payload.ExpiresAt.Time,
		RefreshToken:         refreshToken,
		RefreshTokenExpiryAt: refreshTokenPayload.ExpiresAt.Time,
	}, nil
}
