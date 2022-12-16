package auth

import (
	"fmt"
	"time"

	"github.com/burkayaltuntas/go-fast-temp/pkg/data/models"
	"github.com/burkayaltuntas/go-fast-temp/pkg/dto"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var _db *gorm.DB
var _jwtSecret string

type AuthService struct {
}

func NewAuthService(db *gorm.DB, jwtsecret string) *AuthService {
	_db = db
	_jwtSecret = jwtsecret
	return &AuthService{}
}

func (*AuthService) Login(email string, password string) (*dto.UserDto, string, error) {
	userVM := &dto.UserDto{}

	u := &models.User{}
	_db.First(u, `"Email" = ? and "DeletedAt" is NULL`, email)
	if u.Email == "" {
		return nil, "", fmt.Errorf("user not found")
	}

	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))

	if err != nil {
		return userVM, "", fmt.Errorf(": %v", err)
	}

	t, err := generateToken(ToUserDto(u))
	if err != nil {
		return userVM, "", fmt.Errorf(": %v", err)
	}
	return ToUserDto(u), t, nil
}

func (*AuthService) Register(userVM *dto.UserDto) (*dto.UserDto, error) {
	p, err := bcrypt.GenerateFromPassword([]byte(userVM.Password), bcrypt.DefaultCost)
	if err != nil {
		return userVM, fmt.Errorf("error hashing password: %v", err)
	}
	u := &models.User{
		Email:     userVM.Email,
		Password:  string(p),
		Role:      1,
		Name:      userVM.Name,
		Surname:   userVM.Surname,
		CreatedAt: userVM.CreatedAt,
		UpdatedAt: userVM.UpdatedAt,
	}
	_db.Create(u)
	return ToUserDto(u), nil
}

func generateToken(u *dto.UserDto) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":      u.Id,
		"email":   u.Email,
		"role":    u.Role,
		"name":    u.Name,
		"surname": u.Surname,
		"isLead":  u.IsLead,
		"iat":     time.Now().Unix(),
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := t.SignedString([]byte(_jwtSecret))

	if err != nil {
		return "", err
	}

	return token, nil
}

func (*AuthService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, nil
		}
		return []byte(_jwtSecret), nil
	})
}
