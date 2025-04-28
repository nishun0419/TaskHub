package customer

import (
	"backend/domain/customer"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockCustomerRepository struct {
	mock.Mock
}

func (m *MockCustomerRepository) RegisterCustomer(customer *customer.Customer) error {
	args := m.Called(customer)
	return args.Error(0)
}

func (m *MockCustomerRepository) FindByEmail(email string) (customer.Customer, error) {
	args := m.Called(email)
	return args.Get(0).(customer.Customer), args.Error(1)
}

func TestRegisterCustomer(t *testing.T) {
	// Setup
	mockRepo := new(MockCustomerRepository)
	usecase := NewCustomerUsecase(mockRepo)

	input := customer.RegisterInput{
		Username: "testuser",
		Email:    "test@example.com",
		Password: "password123",
	}

	mockRepo.On("RegisterCustomer", mock.AnythingOfType("*customer.Customer")).Return(nil)

	err := usecase.RegisterCustomer(input)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestAuthenticate(t *testing.T) {
	// Setup
	mockRepo := new(MockCustomerRepository)
	usecase := NewCustomerUsecase(mockRepo)

	email := "test@example.com"
	password := "password123"
	hashedPassword, _ := hashPassword(password)

	expectedCustomer := customer.Customer{
		CustomerID: 1,
		Username:   "testuser",
		Email:      email,
		Password:   hashedPassword,
	}

	// Test successful authentication
	t.Run("successful authentication", func(t *testing.T) {
		mockRepo.On("FindByEmail", email).Return(expectedCustomer, nil)

		cust, err := usecase.Authenticate(email, password)

		assert.NoError(t, err)
		assert.Equal(t, expectedCustomer.CustomerID, cust.CustomerID)
		assert.Equal(t, expectedCustomer.Username, cust.Username)
		assert.Equal(t, expectedCustomer.Email, cust.Email)
	})

	// Test failed authentication with wrong password
	t.Run("failed authentication with wrong password", func(t *testing.T) {
		mockRepo.On("FindByEmail", email).Return(expectedCustomer, nil)

		_, err := usecase.Authenticate(email, "wrongpassword")

		assert.Error(t, err)
	})

	// Test failed authentication with non-existent email
	t.Run("failed authentication with non-existent email", func(t *testing.T) {
		mockRepo.On("FindByEmail", "nonexistent@example.com").Return(customer.Customer{}, assert.AnError)

		_, err := usecase.Authenticate("nonexistent@example.com", password)

		assert.Error(t, err)
	})
}

func TestGenerateToken(t *testing.T) {
	// Setup
	mockRepo := new(MockCustomerRepository)
	usecase := NewCustomerUsecase(mockRepo)

	// Set JWT secret for testing
	os.Setenv("JWT_SECRET", "test-secret")
	defer os.Unsetenv("JWT_SECRET")

	cust := customer.Customer{
		CustomerID: 1,
		Username:   "testuser",
		Email:      "test@example.com",
	}

	// Execute
	token, err := usecase.GenerateToken(cust)

	// Assert
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}
