package users

// Repository is an interface that defines the methods that repository should implement
type Repository interface {
	GetByID(id int) (*User, error)
}
