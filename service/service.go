package service

import(
	"fmt"
	"time"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/SaKu2110/test/model"
)

var secretKey = "75c92a074c341e9964329c0550c2673730ed8479c885c43122c90a2843177d5ef21cb50cfadcccb20aeb730487c11e09ee4dbbb02387242ef264e74cbee97213"

func CheckRequestValue(request model.LoginRequest) error {
	if request.ID == "" {
		return fmt.Errorf("client id value did not exist")
	}
	if request.PASSWORD == "" {
		return fmt.Errorf("client password value did not exist")
	}

	return nil
}

func CreateUserToken(user model.User) (string, error) {
	token := jwt.New(jwt.GetSigningMethod("HS256"))

	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	claims["role"] = user.ROLE
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	stringToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return stringToken, nil
}