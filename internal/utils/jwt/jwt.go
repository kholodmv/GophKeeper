package jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/kholodmv/GophKeeper/internal/models"
	"log"
	"time"
)

type Claims struct {
	jwt.RegisteredClaims
	UserID uint
	Login  string
}

const SecretKey = "bvaEFBtr5e"

// NewToken creates new JWT token for given user and app.
func NewToken(user *models.User, duration time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
		},
		UserID: user.ID,
		Login:  user.Login,
	})

	// Sign the token using the application's secret key
	tokenString, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GetUserIDFromToken(tokenString string) (uint, error) {
	// create an instance of a structure with statements
	claims := &Claims{}
	_, err := jwt.ParseWithClaims(tokenString, claims,
		func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			}
			return []byte(SecretKey), nil
		})
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return claims.UserID, nil
}
