package customer

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"backend/pkg/db/mocks"
	service "backend/service/customer"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func TestLogin(t *testing.T) {
	router := gin.Default()

	// モックデータベースを設定
	db, mock, err := mocks.GetDBMock()
	if err != nil {
		t.Fatalf("Failed to create SQL mock: %v", err)
	}

	// テスト用のパスワードハッシュを生成
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
	if err != nil {
		t.Fatalf("Failed to hash password: %v", err)
	}

	// モックの期待動作を設定
	expectedQuery := "SELECT \\* FROM `customers` WHERE email = \\? ORDER BY `customers`.`customer_id` LIMIT \\?"
	fmt.Printf("Expected query pattern: %s\n", expectedQuery)

	mock.ExpectQuery(expectedQuery).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnRows(sqlmock.NewRows([]string{"customer_id", "name", "email", "password"}).
			AddRow(1, "Test User", "test@example.com", string(hashedPassword)))

	// モックDBを設定
	customerService := service.NewCustomerService(db)
	router.POST("/login", LoginHandler(customerService))

	t.Run("正常なログイン", func(t *testing.T) {
		requestBody := `{"email":"test@example.com","password":"password123"}`
		req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer([]byte(requestBody)))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), `"success":true`)
		assert.Contains(t, w.Body.String(), `"token"`)
	})

	t.Run("メールアドレスが欠けている", func(t *testing.T) {
		requestBody := `{"password":"password123"}`
		req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer([]byte(requestBody)))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.JSONEq(t, `{"success":false,"message":["メールアドレスは必須項目です。"]}`, w.Body.String())
	})

	t.Run("パスワードが欠けている", func(t *testing.T) {
		requestBody := `{"email":"test@example.com"}`
		req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer([]byte(requestBody)))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.JSONEq(t, `{"success":false,"message":["パスワードは必須項目です。"]}`, w.Body.String())
	})

	t.Run("メールアドレスが有効でない", func(t *testing.T) {
		requestBody := `{"email":"invalid-email","password":"password123"}`
		req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer([]byte(requestBody)))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.JSONEq(t, `{"success":false,"message":["メールアドレスは有効なメールアドレスではありません。"]}`, w.Body.String())
	})

	t.Run("存在しないメールアドレス", func(t *testing.T) {
		// モックの期待動作を設定
		mock.ExpectQuery(expectedQuery).
			WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).
			WillReturnError(gorm.ErrRecordNotFound)

		requestBody := `{"email":"nonexistent@example.com","password":"password123"}`
		req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer([]byte(requestBody)))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
		assert.JSONEq(t, `{"success":false,"message":["メールアドレスまたはパスワードが正しくありません。"]}`, w.Body.String())
	})

	t.Run("間違ったパスワード", func(t *testing.T) {
		// モックの期待動作を設定
		mock.ExpectQuery(expectedQuery).
			WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).
			WillReturnRows(sqlmock.NewRows([]string{"customer_id", "name", "email", "password"}).
				AddRow(1, "Test User", "test@example.com", string(hashedPassword)))

		requestBody := `{"email":"test@example.com","password":"wrongpassword"}`
		req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer([]byte(requestBody)))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
		assert.JSONEq(t, `{"success":false,"message":["メールアドレスまたはパスワードが正しくありません。"]}`, w.Body.String())
	})
}
