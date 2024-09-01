package customer

import (
	"backend/models/customer"
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// CustomerService defines a service with a database connection
type CustomerService struct {
	DB *gorm.DB
}

// NewCustomerService creates a new instance of CustomerService
func NewCustomerService(db *gorm.DB) *CustomerService {
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
	savedCustomer := customer.Customer{Name: input.Username, Password: input.Password}
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
