package controller

import (
	"fmt"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CheckUserLogin(c *gin.Context) {

}

func CheckUserLogout(c *gin.Context) {

}

func Login(ctx *gin.Context) Resp {
	f := "Login()-->"
	var authParam Claims
	if err := ctx.BindJSON(&authParam); err != nil {
		// parse param failed
		return Resp{Message: fmt.Sprintf("%s, parse param failed :%s", f, err.Error()), Code: code.ParamInvalid}
	}

	token, err := login(authParam.Username, authParam.Password)
	if err != nil {
		// login failed
		return Resp{Message: fmt.Sprintf("%s, login failed :%s", f, err.Error()), Code: code.ParamInvalid}
	}
	return Resp{Data: token}
}

func login(userName, Pwd string) (string, error) {
	if err := loginService.Login(userName, Pwd); err != nil {
		return "", err
	}

	return GenerateToken(userName, Pwd)
}

func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		claims, err := ParseToken(token)
		if err != nil {
			c.JSON(http.StatusOK, "please login")
			c.Abort()
			return
		}

		if err := claims.Valid(); err != nil {
			c.JSON(http.StatusOK, "place login")
			c.Abort()
			return
		}

		c.Set("user", claims.Username)
		c.Next()
	}
}

type Claims struct {
	Username string `json:"username"  binding:"required"`
	Password string `json:"password"  binding:"required"`
	jwt.StandardClaims
}

func GenerateToken(username, password string) (string, error) {
	expireTime := time.Now().Add(time.Duration(config.Conf.Login.TokenEffectiveHour) * time.Hour)
	claims := Claims{username, password,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "ipalfish-db-injection",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return tokenClaims.SignedString([]byte(config.Conf.Login.TokenSecret))
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Conf.Login.TokenSecret), nil
	})
	if err != nil {
		return nil, err
	}

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
		return nil, fmt.Errorf("parse token failed, not a claims ins")
	}
	return nil, fmt.Errorf("get nil token claims")
}
