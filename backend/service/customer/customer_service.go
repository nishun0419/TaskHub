package customer

import (
	"backend/models/customer"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// CustomerService interface
type CustomerServiceInterface interface {
	RegisterCustomer(input customer.RegisterInput) (customer.Customer, error)
	Authenticate(email, password string) (customer.Customer, error)
	GenerateToken(cust customer.Customer) (string, error)
}

// CustomerService defines a service with a database connection
type CustomerService struct {
	DB *gorm.DB
}

// NewCustomerService creates a new instance of CustomerService
func NewCustomerService(db *gorm.DB) CustomerServiceInterface {
	return &CustomerService{DB: db}
}

// RegisterCustomer handles the logic of registering a new customer
func (s *CustomerService) RegisterCustomer(input customer.RegisterInput) (customer.Customer, error) {
	// パスワードのハッシュ化
	hashedPassword, err := hashPassword(input.Password)
	if err != nil {
		return customer.Customer{}, fmt.Errorf("failed to hash password: %w", err)
	}

	// ハッシュ化されたパスワードを設定
	input.Password = hashedPassword
	savedCustomer := customer.Customer{Username: input.Username, Email: input.Email, Password: input.Password}
	// 顧客をデータベースに保存
	if result := s.DB.Create(&savedCustomer); result.Error != nil {
		return customer.Customer{}, fmt.Errorf("failed to register customer: %w", err)
	}

	return customer.Customer{}, nil
}

// hashPassword パスワードをハッシュ化するヘルパー関数
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// ユーザー認証を行う
func (s *CustomerService) Authenticate(email, password string) (customer.Customer, error) {
	var cust customer.Customer
	if err := s.DB.Where("email = ?", email).First(&cust).Error; err != nil {
		return customer.Customer{}, fmt.Errorf("failed to authenticate customer: %w", err)
	}

	err := bcrypt.CompareHashAndPassword([]byte(cust.Password), []byte(password))
	if err != nil {
		return customer.Customer{}, fmt.Errorf("failed to authenticate customer: %w", err)
	}

	return cust, nil
}

// JWTトークンを生成する
func (s *CustomerService) GenerateToken(cust customer.Customer) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"customer_id": cust.CustomerID,
		"exp":         time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}
