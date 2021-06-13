package repository

import (
	"context"
	"fmt"
	"time"
	"user_ms/backend/core/internal/constant"
	"user_ms/backend/core/internal/models"

	"github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

// function parse user info from ctx
func ParseUserFromCTX(ctx context.Context) *UserClaims {

	userClaim := ctx.Value(constant.JWTKey)
	if userClaim != nil {
		if userClaim, ok := userClaim.(UserClaims); ok {
			return &userClaim
		}
	}

	return &UserClaims{
		Username: "Anonymous",
		Role:     "anonymous",
	}
}

// function validate user access role
func ValidateAcessRole(ctx context.Context, roles ...string) (*UserClaims, error) {

	userClaim := ctx.Value(constant.JWTKey)
	if userClaim != nil {
		if userClaim, ok := userClaim.(UserClaims); ok {
			for _, r := range roles {
				if userClaim.Role == r || r == "all" {
					return &userClaim, nil
				}
			}
			return nil, status.Errorf(codes.PermissionDenied, "do not have premission to access")

		}
	}

	return nil, status.Errorf(codes.Unauthenticated, "authorization info is not provided")
}
