package customer

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"backend/pkg/db/mocks"
	service "backend/service/customer"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql" // MySQL ドライバ
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	// Ginのデフォルトのルーターを使用
	router := gin.Default()

	// モックデータベースを設定
	db, mock, err := mocks.GetDBMock() // mocks.GetDBMock()はモックDBを返す関数
	if err != nil {
		t.Fatalf("Failed to create SQL mock: %v", err)
	}

	// モックの期待動作を設定
	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `customers`").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// モックDBを設定
	customerService := service.NewCustomerService(db)
	// /register エンドポイントに対して Register ハンドラを設定
	router.POST("/register", RegisterHandler(customerService))
	t.Run("登録成功できている", func(t *testing.T) {
		requestBody := `{"username":"john_doe","password":"secret","email":"john_doe@example.com"}`
		req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer([]byte(requestBody)))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.JSONEq(t, `{"success":true,"message":"登録が完了しました。"}`, w.Body.String())
	})

	t.Run("ユーザー名が欠けている", func(t *testing.T) {
		requestBody := `{"password":"secret","email":"john_doe@example.com"}`
		req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer([]byte(requestBody)))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.JSONEq(t, `{"success":false,"message":["ユーザー名は必須項目です。"]}`, w.Body.String())
	})

	t.Run("パスワードが欠けている", func(t *testing.T) {
		requestBody := `{"username":"john_doe","email":"john_doe@example.com"}`
		req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer([]byte(requestBody)))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.JSONEq(t, `{"success":false,"message":["パスワードは必須項目です。"]}`, w.Body.String())
	})

	t.Run("メールアドレスが欠けている", func(t *testing.T) {
		requestBody := `{"username":"john_doe","password":"secret"}`
		req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer([]byte(requestBody)))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.JSONEq(t, `{"success":false,"message":["メールアドレスは必須項目です。"]}`, w.Body.String())
	})

	t.Run("メールアドレスが有効でない", func(t *testing.T) {
		requestBody := `{"username":"john_doe","password":"secret","email":"invalid-email"}`
		req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer([]byte(requestBody)))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.JSONEq(t, `{"success":false,"message":["メールアドレスは有効なメールアドレスではありません。"]}`, w.Body.String())
	})

	t.Run("ユーザー名の文字数が3文字未満", func(t *testing.T) {
		requestBody := `{"username":"jo","password":"secret","email":"john_doe@example.com"}`
		req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer([]byte(requestBody)))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.JSONEq(t, `{"success":false,"message":["ユーザー名は3文字以上で入力してください。"]}`, w.Body.String())
	})

	t.Run("ユーザー名の文字数が10文字より多い", func(t *testing.T) {
		requestBody := `{"username":"jo123456789","password":"secret","email":"john_doe@example.com"}`
		req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer([]byte(requestBody)))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.JSONEq(t, `{"success":false,"message":["ユーザー名は10文字以下で入力してください。"]}`, w.Body.String())
	})

	t.Run("パスワードの文字数が6文字未満", func(t *testing.T) {
		requestBody := `{"username":"john_doe","password":"secr","email":"john_doe@example.com"}`
		req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer([]byte(requestBody)))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.JSONEq(t, `{"success":false,"message":["パスワードは6文字以上で入力してください。"]}`, w.Body.String())
	})

	t.Run("パスワードの文字数が20文字より多い", func(t *testing.T) {
		requestBody := `{"username":"john_doe","email":"john_doe@example.com","password":"123456789012345678901"}`
		req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer([]byte(requestBody)))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.JSONEq(t, `{"success":false,"message":["パスワードは20文字以下で入力してください。"]}`, w.Body.String())
	})
}
