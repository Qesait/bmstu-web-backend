package app

import (
	"crypto/sha1"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	// "bmstu-web-backend/internal/app/config"
	// "bmstu-web-backend/internal/app/dsn"
	// "bmstu-web-backend/internal/app/repository"
	"bmstu-web-backend/internal/app/ds"
	"bmstu-web-backend/internal/app/role"
	"encoding/hex"
	"github.com/golang-jwt/jwt"
)

type loginReq struct {
	Login    string `json:"login" binding:"required,max=30"`
	Password string `json:"password" binding:"required,max=30"`
}

type loginResp struct {
	ExpiresIn   time.Duration `json:"expires_in"`
	AccessToken string        `json:"access_token"`
	TokenType   string        `json:"token_type"`
}

func (a *Application) Login(c *gin.Context) {
	JWTConfig := a.config.JWT
	request := &loginReq{}
	if err := c.ShouldBind(request); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	user, err := a.repo.GetUserByLogin(request.Login)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// TODO: сравнить логины?
	if user.Password != generateHashString(request.Password) {
		c.AbortWithStatus(http.StatusForbidden) // отдаем 403 ответ в знак того что доступ запрещен
		return
	}
	// значит проверка пройдена
	// генерируем ему jwt
	token := jwt.NewWithClaims(JWTConfig.SigningMethod, &ds.JWTClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(JWTConfig.ExpiresIn).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "bitop-admin",
		},
		Role: user.Role,
	})
	if token == nil {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("token is nil"))
		return
	}

	strToken, err := token.SignedString([]byte(JWTConfig.Token))
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("cant create str token"))
		return
	}

	c.JSON(http.StatusOK, loginResp{
		ExpiresIn:   JWTConfig.ExpiresIn,
		AccessToken: strToken,
		TokenType:   "Bearer",
	})

}

type registerReq struct {
	Login    string `json:"login" binding:"required,max=30"`
	Password string `json:"password" binding:"required,max=30"`
}

type registerResp struct {
	Ok bool `json:"ok"`
}

func (a *Application) Register(c *gin.Context) {
	request := &registerReq{}
	if err := c.ShouldBind(request); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if request.Password == "" {
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("password is empty"))
		return
	}

	if request.Login == "" {
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("login is empty"))
		return
	}

	if err := a.repo.AddUser(&ds.User{
		Role:     role.Customer,
		Login:    request.Login,
		Password: generateHashString(request.Password),
	}); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, &registerResp{
		Ok: true,
	})
}

func generateHashString(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

// Ping godoc
// @Summary      Show hello text
// @Description  very very friendly response
// @Tags         Tests
// @Produce      json
// @Success      200  {object}  pingResp
// @Router       /ping/{name} [get]
func (a *Application) Ping(gCtx *gin.Context) {
	name := gCtx.Param("name")
	gCtx.String(http.StatusOK, "Hello %s", name)
}

// type pingReq struct{}
// type pingResp struct {
// 	Status string `json:"status"`
// }
