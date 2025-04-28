package customer

type CustomerRepository interface {
	RegisterCustomer(customer *Customer) error
	FindByEmail(email string) error
}
