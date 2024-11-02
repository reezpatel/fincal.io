package lib

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"os"
	"time"
)

var (
	secret = os.Getenv("JWT_SECRET")
)

func init() {
	if len(secret) == 0 {
		secret = "my_super_secret_key"
	}
}

type AppClaim struct {
	UserId int64 `json:"user_id"`
	jwt.StandardClaims
}

func CreateJWTToken(userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, AppClaim{
		userId,
		jwt.StandardClaims{
			ExpiresAt: int64(time.Hour * 24 * 365),
		},
	})

	str, err := token.SignedString([]byte(secret))

	if err != nil {
		return "", err
	}

	return str, nil
}

func ValidateJWTToken(token string) (int64, error) {
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(secret), nil
	})

	if err != nil {
		return -1, err
	}

	if claims, ok := t.Claims.(AppClaim); ok {
		if claims.UserId > 0 {
			return claims.UserId, nil
		}
	}

	return -1, fmt.Errorf("failed to get user from token: %s", err)
}
