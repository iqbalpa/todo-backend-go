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
	tokenString, err := token.SignedString([]byte(config.GetJwtSecretKey()))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// verify the token if needed
func VerifyToken(tokenString string) error {
	config := config.LoadEnv()
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.GetJwtSecretKey()), nil
	})
	if err != nil {
		return err
	}
	if !token.Valid {
		return fmt.Errorf("token is invalid")
	}
	return nil
}

// extract the userId from jwt token
func ExtractClaimsUserId(tokenString string) (int, error) {
	// parse the jwt token
	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		return -1, fmt.Errorf("failed to parse the token")
	}
	// check the claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		// Use type assertion to convert the value to int
		if userIdFloat, ok := claims["userId"].(float64); ok {
			userId := int(userIdFloat)
			return userId, nil
		} else {
			return -1, fmt.Errorf("userId is not a float64")
		}
	}
	return -1, fmt.Errorf("failed to get the claims")
}