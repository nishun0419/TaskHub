package customer

import (
	"backend/models/customer"
	"testing"

	"backend/pkg/db/mocks"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/stretchr/testify/assert"
)

func TestAuthenticate(t *testing.T) {
	// モックデータベースを設定
	db, mock, err := mocks.GetDBMock()
	if err != nil {
		t.Fatalf("Failed to create SQL mock: %v", err)
	}
	service := NewCustomerService(db)

	// テスト用のユーザーを作成
	testUser := customer.Customer{
		Username: "Test User1",
		Email:    "test@example.com",
		Password: "password123",
	}
	hashedPassword, err := hashPassword(testUser.Password)
	assert.NoError(t, err)
	testUser.Password = hashedPassword

	// トランザクションのモックを設定
	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `customers`").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// ユーザーをデータベースに保存
	err = db.Create(&testUser).Error
	assert.NoError(t, err)

	tests := []struct {
		username     string
		email        string
		password     string
		expectError  bool
		errorMessage string
	}{
		{
			username:    "Test User1",
			email:       "test@example.com",
			password:    "password123",
			expectError: false,
		},
		{
			username:     "Test User2",
			email:        "nonexistent@example.com",
			password:     "password123",
			expectError:  true,
			errorMessage: "failed to authenticate customer: record not found",
		},
		{
			username:     "Test User3",
			email:        "test@example.com",
			password:     "wrongpassword",
			expectError:  true,
			errorMessage: "failed to authenticate customer: crypto/bcrypt: hashedPassword is not the hash of the given password",
		},
	}

	for _, tt := range tests {
		t.Run(tt.username, func(t *testing.T) {
			// トランザクションのモックを設定
			mock.ExpectBegin()
			mock.ExpectCommit()
			user, err := service.Authenticate(tt.email, tt.password)

			if tt.expectError {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.errorMessage)
				assert.Empty(t, user)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, testUser.Email, user.Email)
				assert.Equal(t, testUser.Username, user.Username)
			}
		})
	}
}

func TestGenerateToken(t *testing.T) {
	db, _, err := mocks.GetDBMock()
	if err != nil {
		t.Fatalf("Failed to create SQL mock: %v", err)
	}
	service := NewCustomerService(db)

	// テスト用のユーザーを作成
	testUser := customer.Customer{
		CustomerID: 1,
		Username:   "Test User",
		Email:      "test@example.com",
		Password:   "hashedpassword",
	}

	// トークン生成をテスト
	token, err := service.GenerateToken(testUser)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}
