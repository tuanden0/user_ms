package repository

import (
	"fmt"
	"time"
	"user_ms/backend/core/internal/models"

	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
)

type JWTManager struct {
	secretKey     string
	tokenDuration time.Duration
	db            *gorm.DB
}

func NewJWTManager(secretKey string, tokenDuration time.Duration, db *gorm.DB) *JWTManager {
	return &JWTManager{secretKey, tokenDuration, db}
}

type UserClaims struct {
	jwt.StandardClaims
	Username string `json:"username"`
	Role     string `json:"role"`
}

func (manager *JWTManager) Verify(accessToken string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(
		accessToken,
		&UserClaims{},
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, fmt.Errorf("unexpected token signing method")
			}

			return []byte(manager.secretKey), nil
		},
	)

	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	return claims, nil
}

func (manager *JWTManager) Login(username string) (*models.User, error) {
	user := models.User{}

	if err := manager.db.First(&user, "username = ?", username).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (manager *JWTManager) GenerateJWTToken(u *models.User) (string, error) {

	claims := UserClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(manager.tokenDuration).Unix(),
		},
		Username: u.Username,
		Role:     u.Role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(manager.secretKey))
}
