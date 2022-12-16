package common

import (
	"time"

	"github.com/gin-gonic/gin"
)

type ContextUser struct {
	Id        string
	Role      Role
	RoleName  string
	Email     string
	Password  string
	Name      string
	Surname   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func ExtractUser(ctx *gin.Context) ContextUser {
	u, exist := ctx.Get("user")
	if !exist {
		panic("user can't found in the context")
	}
	return u.(ContextUser)
}
