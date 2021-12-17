package middleware

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"time"
	"vaccine-app-be/exceptions"
)

type jwtCustomClaims struct {
	UserId int    `json:"user_id"`
	Name   string `json:"name"`
	jwt.StandardClaims
}

type ConfigJWT struct {
	SecretJWT string
	ExpiredIn int
}

func (jwtConf *ConfigJWT) Init() middleware.JWTConfig {
	return middleware.JWTConfig{
		Claims:     &jwtCustomClaims{},
		SigningKey: []byte(jwtConf.SecretJWT),
	}
}

//generate new token
func (jwtConf *ConfigJWT) GenerateToken(id int, name string) string {
	claims := &jwtCustomClaims{
		id,
		name,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(int64(jwtConf.ExpiredIn))).Unix(),
		},
	}
	//membuat token dari claims yang isinya data" tersebut
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(jwtConf.SecretJWT))
	exceptions.PanicIfError(err)

	return t
}

func GetUserId(c echo.Context) int {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*jwtCustomClaims)
	id := claims.UserId
	return id
}

func GetUserName(c echo.Context) string {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*jwtCustomClaims)
	name := claims.Name
	return name
}

