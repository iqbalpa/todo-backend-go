package utils

import (
	"fmt"
	"time"

	"main/app/models"
	"main/config"

	jwt "github.com/golang-jwt/jwt/v5"
)

// generate token for login
func CreateToken(user models.User) (string, error) {
	config := config.LoadEnv()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, 
		jwt.MapClaims{
			"username": user.Username,
			"name": user.Name,
			"userId": user.ID,
			"expired": time.Now().Add(time.Hour * 1).Unix(),
		},
	)
	tokenString, err := token.SignedString(config.GetJwtExpiration())
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// verify the token if needed
func VerifyToken(tokenString string) error {
	config := config.LoadEnv()
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return config.GetJwtSecretKey(), nil
	})
	if err != nil {
		return err
	}
	if !token.Valid {
		return fmt.Errorf("token is invalid")
	}
	return nil
}