package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
)

var (
	SecretKey = []byte("secret")
)

// GenerateToken 生成一个 jwt 令牌，并将用户名分配Token一并返回
func GenerateToken(username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	// 创建claims
	claims := token.Claims.(jwt.MapClaims)
	// 设置 claims
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		log.Fatal("Error in Generating  Key!")
		return "", err
	}
	return tokenString, nil
}

func ParseToken(tokenStr string) (string, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) { return SecretKey, nil })
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		username := claims["username"].(string)
		return username, nil
	} else {
		return "", err
	}
}
