package customer

import (
	"backend/domain/customer"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type CustomerUsecase struct {
	CustomerRepository customer.CustomerRepository
}

// NewCustomerService creates a new instance of CustomerService
func NewCustomerService(repo customer.CustomerRepository) *CustomerUsecase {
	return &CustomerUsecase{repo}
}

func (u *CustomerUsecase) RegisterCustomer(input customer.RegisterInput) error {
	// パスワードのハッシュ化
	hashedPassword, err := hashPassword(input.Password)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	// ハッシュ化されたパスワードを設定
	input.Password = hashedPassword
	savedCustomer := customer.Customer{Username: input.Username, Email: input.Email, Password: input.Password}
	// 顧客をデータベースに保存
	if err := u.CustomerRepository.RegisterCustomer(&savedCustomer); err != nil {
		return fmt.Errorf("failed to register customer: %w", err)
	}

	return nil
}

// hashPassword パスワードをハッシュ化するヘルパー関数
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// ユーザー認証を行う
func (u *CustomerUsecase) Authenticate(email, password string) (customer.Customer, error) {
	var cust customer.Customer
	if err := u.CustomerRepository.FindByEmail(email); err != nil {
		return customer.Customer{}, fmt.Errorf("failed to authenticate customer: %w", err)
	}

	err := bcrypt.CompareHashAndPassword([]byte(cust.Password), []byte(password))
	if err != nil {
		return customer.Customer{}, fmt.Errorf("failed to authenticate customer: %w", err)
	}

	return cust, nil
}

// JWTトークンを生成する
func (u *CustomerUsecase) GenerateToken(cust customer.Customer) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"customer_id": cust.CustomerID,
		"exp":         time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}
