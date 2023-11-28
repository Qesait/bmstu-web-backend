package ds

import (
	"bmstu-web-backend/internal/app/role"
	"github.com/golang-jwt/jwt"
)

type JWTClaims struct {
	jwt.StandardClaims        // все что точно необходимо по RFC
	UserUUID           string `json:"user_uuid"` // наши данные - uuid этого пользователя в базе данных
	Role               role.Role
}
