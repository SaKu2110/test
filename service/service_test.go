package service

import(
	"testing"
	"github.com/SaKu2110/test/model"
)

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

//　構造体userの中身が存在することはcontroller.goの41 ~ 58行目で保証されているので
//  tokenが正常に発行されているのかを確かめるテストケースのみを作成した.
func TestCreateUserTokenSuccess(t *testing.T) {
	user := model.User{
		ID: "000000",
		PASSWORD: "password",
		ADMIN: true,
	}

    token, err := CreateUserToken(user)
    if err != nil {
        t.Fatalf("failed test %#v", err)
	}
	if token == "" {
		t.Fatalf("failed test %#v", err)
	}
}