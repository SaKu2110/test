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