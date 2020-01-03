package controller

import(
    "regexp"
	"testing"

    sqlmock "github.com/DATA-DOG/go-sqlmock"
    "github.com/jinzhu/gorm"
)

func InitializeDBMoc() (*gorm.DB, sqlmock.Sqlmock, error){
	db, mock, err := sqlmock.New()
    if err != nil {
        return nil, nil, err
    }

    gdb, err := gorm.Open("mysql", db)
    if err != nil {
        return nil, nil, err
    }
    return gdb, mock, nil
}

func TestSignInHandlerSuccess(t *testing.T) {
    db, mock, err := InitializeDBMoc()
    if err != nil {
        t.Fatal(err)
    }
    defer db.Close()
    db.LogMode(true)

    ctrl := IsController{DB: db}

    id := "saku2110"
    password := "hogehoge"
    role := "admin"

    // Mockの設定
    mock.ExpectQuery(regexp.QuoteMeta(
        `SELECT * FROM "users" WHERE (id = $1)`)).
        WithArgs(id).
        WillReturnRows(
            sqlmock.NewRows([]string{"id", "password", "admin"}).
            AddRow(id, password, role))

    // 実行
    // ctrl.SignInHandler()
}