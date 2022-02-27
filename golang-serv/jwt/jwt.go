package jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"io/ioutil"
	"os"
	response "service/response/auth"
)

var hmacSecret []byte

type Claims struct {
	*response.Auth
	jwt.StandardClaims
}

func getSigningKey() ([]byte, error) {
	dir, err := os.Getwd()

	if err != nil {
		return nil, err
	}

	dir = fmt.Sprintf("%s/../.jwt_pass", dir)

	file, err := os.Open(dir)

	if err != nil {
		return nil, err
	}

	buf, err := ioutil.ReadAll(file)

	if err != nil {
		return nil, err
	}

	return buf, nil
}

func Generate(auth *response.Auth) (string, error) {
	resource := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		Claims{auth, jwt.StandardClaims{}},
	)

	key, err := getSigningKey()

	if err != nil {
		return "", err
	}

	token, err := resource.SignedString(key)

	if err != nil {
		return "", err
	}

	return token, nil
}

func Parse(token string) (*response.Auth, error) {
	claims, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
		}

		key, err := getSigningKey()

		if err != nil {
			return nil, err
		}

		return []byte(key), nil
	})

	if err != nil {
		return nil, err
	}

	auth := response.NewAuth()

	if data, ok := claims.Claims.(jwt.MapClaims); ok && claims.Valid {
		auth.SetName(data["name"].(string))
		auth.SetPhone(data["phone"].(string))
		auth.SetRole(data["role"].(string))
		auth.SetTimestamp(int64(data["timestamp"].(float64)))
		return auth, nil
	}

	return nil, err
}
