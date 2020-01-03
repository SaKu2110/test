package controller

import(
    "fmt"
    "regexp"
    "testing"
    "strings"
    "net/http/httptest"
    "github.com/stretchr/testify/assert"

    "net/http"
    "github.com/gin-gonic/gin"

    sqlmock "github.com/DATA-DOG/go-sqlmock"
    "github.com/jinzhu/gorm"
)

func setupRouter(db *gorm.DB) *gin.Engine {
    ctrl := IsController{DB: db}

	router := gin.Default()
    router.POST("/signin", ctrl.SignInHandler)
	return router
}

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

// テストケース 1 成功
func TestSignInHandlerSuccess(t *testing.T) {
    // create db mock
    db, mock, err := InitializeDBMoc()
    if err != nil {
        t.Fatal(err)
    }
    defer db.Close()
    db.LogMode(true)

    id := "saku2110"
    password := "hogehoge"
    role := "admin"

    // setup db mock
    mock.ExpectQuery(regexp.QuoteMeta(
        "SELECT * FROM `users` WHERE (id=?)")).
        WithArgs(id).
        WillReturnRows(
            sqlmock.NewRows([]string{"id", "password", "admin"}).
            AddRow(id, password, role))

    // setup router
    router := setupRouter(db)

    // create request
    body := strings.NewReader(`{"id": "saku2110", "password": "hogehoge"}`)

    request := httptest.NewRequest("POST", "/signin", body)
    request.Header.Add("Content-Type", "application/json")
    request.Header.Add("Accept", "application/json")

    recorder := httptest.NewRecorder()
    router.ServeHTTP(recorder, request)

    assert.Equal(t, http.StatusOK, recorder.Code)
    fmt.Println(recorder.Body.String())
}

/** 失敗ケース 2
 * user idが入力されていないためエラー
 */
func TestFaildIDValue(t *testing.T) {
    // create db mock
    db, mock, err := InitializeDBMoc()
    if err != nil {
        t.Fatal(err)
    }
    defer db.Close()
    db.LogMode(true)

    id := "saku2110"
    password := "hogehoge"
    role := "admin"

    // setup db mock
    mock.ExpectQuery(regexp.QuoteMeta(
        "SELECT * FROM `users` WHERE (id=?)")).
        WithArgs(id).
        WillReturnRows(
            sqlmock.NewRows([]string{"id", "password", "admin"}).
            AddRow(id, password, role))

    // setup router
    router := setupRouter(db)

    // create request
    body := strings.NewReader(`{"id": "", "password": "hogehoge"}`)

    request := httptest.NewRequest("POST", "/signin", body)
    request.Header.Add("Content-Type", "application/json")
    request.Header.Add("Accept", "application/json")

    recorder := httptest.NewRecorder()
    router.ServeHTTP(recorder, request)

    assert.Equal(t, http.StatusInternalServerError, recorder.Code)
    fmt.Println(recorder.Body.String())
}

/** 失敗ケース 3
 * user passwordが入力されていないためエラー
 */
func TestFaildPasswordValue(t *testing.T) {
    // create db mock
    db, mock, err := InitializeDBMoc()
    if err != nil {
        t.Fatal(err)
    }
    defer db.Close()
    db.LogMode(true)

    id := "saku2110"
    password := "hogehoge"
    role := "admin"

    // setup db mock
    mock.ExpectQuery(regexp.QuoteMeta(
        "SELECT * FROM `users` WHERE (id=?)")).
        WithArgs(id).
        WillReturnRows(
            sqlmock.NewRows([]string{"id", "password", "admin"}).
            AddRow(id, password, role))

    // setup router
    router := setupRouter(db)

    // create request
    body := strings.NewReader(`{"id": "saku2110", "password": ""}`)

    request := httptest.NewRequest("POST", "/signin", body)
    request.Header.Add("Content-Type", "application/json")
    request.Header.Add("Accept", "application/json")

    recorder := httptest.NewRecorder()
    router.ServeHTTP(recorder, request)

    assert.Equal(t, http.StatusInternalServerError, recorder.Code)
    fmt.Println(recorder.Body.String())
}

/** 失敗ケース 4
 * user idがdb内に存在しなかった
 */
func TestFaildNonexistentId(t *testing.T) {
    // create db mock
    db, mock, err := InitializeDBMoc()
    if err != nil {
        t.Fatal(err)
    }
    defer db.Close()
    db.LogMode(true)

    id := "saku2110"
    password := "hogehoge"
    role := "admin"

    // setup db mock
    mock.ExpectQuery(regexp.QuoteMeta(
        "SELECT * FROM `users` WHERE (id=?)")).
        WithArgs(id).
        WillReturnRows(
            sqlmock.NewRows([]string{"id", "password", "admin"}).
            AddRow(id, password, role))

    // setup router
    router := setupRouter(db)

    // create request
    body := strings.NewReader(`{"id": "saku", "password": "hogehoge"}`)

    request := httptest.NewRequest("POST", "/signin", body)
    request.Header.Add("Content-Type", "application/json")
    request.Header.Add("Accept", "application/json")

    recorder := httptest.NewRecorder()
    router.ServeHTTP(recorder, request)

    assert.Equal(t, http.StatusInternalServerError, recorder.Code)
    fmt.Println(recorder.Body.String())
}

/** 失敗ケース 5
 * user passwordがdb内に存在するものと異なる
 */
 func TestFaildWrongPassword(t *testing.T) {
    // create db mock
    db, mock, err := InitializeDBMoc()
    if err != nil {
        t.Fatal(err)
    }
    defer db.Close()
    db.LogMode(true)

    id := "saku2110"
    password := "hogehoge"
    role := "admin"

    // setup db mock
    mock.ExpectQuery(regexp.QuoteMeta(
        "SELECT * FROM `users` WHERE (id=?)")).
        WithArgs(id).
        WillReturnRows(
            sqlmock.NewRows([]string{"id", "password", "admin"}).
            AddRow(id, password, role))

    // setup router
    router := setupRouter(db)

    // create request
    body := strings.NewReader(`{"id": "saku2110", "password": "hugahuga"}`)

    request := httptest.NewRequest("POST", "/signin", body)
    request.Header.Add("Content-Type", "application/json")
    request.Header.Add("Accept", "application/json")

    recorder := httptest.NewRecorder()
    router.ServeHTTP(recorder, request)

    assert.Equal(t, http.StatusInternalServerError, recorder.Code)
    fmt.Println(recorder.Body.String())
}