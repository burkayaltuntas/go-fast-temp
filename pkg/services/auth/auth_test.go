package auth

import (
	"log"
	"testing"

	"github.com/burkayaltuntas/go-fast-temp/pkg/dto"
	"golang.org/x/crypto/bcrypt"
)

func TestGenerateToken(t *testing.T) {
	u := dto.UserDto{
		Email:   "username@gmail.com",
		Name:    "uername",
		Surname: "Altuntaş",
		Role:    0,
		Id:      "db4530a9-e1d9-48aa-b67d-225fc8ecd242",
	}
	token, e := generateToken(&u)
	if e != nil {
		t.Error(e)
	}
	log.Println("token: ", token)
}

func TestGeneratePasswordHash(t *testing.T) {
	p, err := bcrypt.GenerateFromPassword([]byte("12345678"), bcrypt.DefaultCost)
	if err != nil {
		t.Error(err)
	}
	log.Println(string(p))
}

func TestPasswordHashCompare(t *testing.T) {
	err := bcrypt.CompareHashAndPassword([]byte("token"), []byte("12345678"))
	err2 := bcrypt.CompareHashAndPassword([]byte("token"), []byte("12345678"))
	if err != nil || err2 != nil {
		t.Error(err)
	}
}
