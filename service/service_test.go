package service

import(
	"testing"
	"github.com/SaKu2110/test/model"
)

/** service test 1-1 success
 * リクエストが正しい
 */
func TestCheckRequestValueSuccess(t *testing.T) {
	request := model.LoginRequest{
		ID: "000000",
		PASSWORD: "password",
	}

    err := CheckRequestValue(request)
    if err != nil {
        t.Fatalf("failed test %#v", err)
    }
}

/** service test 1-2 faild
 * user id が存在しない
 */
func TestCheckRequestValueFaildIDValue(t *testing.T) {
	request := model.LoginRequest{
		ID: "",
		PASSWORD: "password",
	}

    err := CheckRequestValue(request)
    if err == nil {
        t.Fatalf("failed test %#v", err)
    }
}

/** service test 1-3 success
 * user password が存在しない
 */
func TestCheckRequestValueFaildPASSWORDVale(t *testing.T) {
	request := model.LoginRequest{
		ID: "000000",
		PASSWORD: "",
	}

    err := CheckRequestValue(request)
    if err == nil {
        t.Fatalf("failed test %#v", err)
    }
}

/** service test 2-1 success
 * jwtを生成する
 * 構造体userの中身が存在することはcontroller.goの41 ~ 58行目で保証されているので
 * tokenが正常に発行されているのかを確かめるテストケースのみを作成した.
 */
func TestCreateUserTokenSuccess(t *testing.T) {
	user := model.User{
		ID: "000000",
		PASSWORD: "password",
		ROLE: "admin",
	}

    token, err := CreateUserToken(user)
    if err != nil {
        t.Fatalf("failed test %#v", err)
	}
	if token == "" {
		t.Fatalf("failed test %#v", err)
	}
}