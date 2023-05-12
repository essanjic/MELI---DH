package users

type MySQLRepository struct {
}

func (repository *MySQLRepository) GetByID(id int) (*User, error) {
	return &User{ID: id, Username: "test", Password: "test"}, nil
}
