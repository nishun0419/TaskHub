package customer

import (
	"backend/domain/customer"

	"gorm.io/gorm"
)

type CustomerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) *CustomerRepository {
	return &CustomerRepository{db}
}

func (r *CustomerRepository) RegisterCustomer(customer *customer.Customer) error {
	return r.db.Create(customer).Error
}

func (r *CustomerRepository) FindByEmail(email string) error {
	var customer customer.Customer
	if err := r.db.Where("email = ?", email).First(&customer).Error; err != nil {
		return err
	}
	return nil
}
