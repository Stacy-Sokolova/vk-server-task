package service

import (
	"context"
	"crypto/sha1"
	"errors"
	"fmt"
	"time"
	"vk-server-task/internal/models"
	"vk-server-task/internal/storage"

	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

const (
	salt       = "hjqrhjqw124617ajfhajs"
	signingKey = "qrkjk#4#%35FSFJlja#4353KSFjH"
	tokenTTL   = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

type AuthService struct {
	storage storage.Auth
}

func NewAuthService(s storage.Auth) *AuthService {
	return &AuthService{storage: s}
}

func (s *AuthService) generateToken(userId int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		userId,
	})

	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, nil
}

func (s *AuthService) CreateUser(ctx context.Context, login, password string) (*models.User, string, error) {
	if _, err := s.storage.GetUserByLogin(ctx, login); err == nil {
		return nil, "", errors.New("user already exists")
	}

	hashed_password := generatePasswordHash(password)

	id, err := s.storage.CreateUser(ctx, login, hashed_password)
	if err != nil {
		logrus.Errorf("failed to create user: %s", err.Error())
		return nil, "", err
	}

	token, err := s.generateToken(id)
	if err != nil {
		logrus.Errorf("failed to generate token: %s", err.Error())
		return nil, "", err
	}

	logrus.Infof("user %d successfully registered", id)
	return &models.User{Id: id, Login: login, Password: hashed_password}, token, nil
}

func (s *AuthService) LoginUser(ctx context.Context, login, password string) (*models.User, string, error) {
	user, err := s.storage.GetUserByLogin(ctx, login)
	if err != nil {
		return nil, "", errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, "", errors.New("invalid credentials")
	}

	token, err := s.generateToken(user.Id)
	if err != nil {
		logrus.Errorf("failed to generate token: %s", err.Error())
		return nil, "", err
	}

	logrus.Infof("user %d successfully logged in", user.Id)
	return user, token, nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
