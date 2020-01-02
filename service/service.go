package service

import(
	"fmt"
	"time"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/SaKu2110/test/model"
)

func CheckRequestValue(request model.LoginRequest) error {
	if request.ID == "" {
		return fmt.Errorf("client id value did not exist")
	}
	if request.PASSWORD == "" {
		return fmt.Errorf("client password value did not exist")
	}

	return nil
}

func CreateUserToken(user model.User) string, error {
	token := jwt.New(jwt.SigningMethodRS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	claims["admin"] = user.ADMIN
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	stringToken, err := token.SignedString(signKey)
	if err != nil {
		return "", err
	}
	return stringToken, nil
}