package main

type Database interface {
	GetUser(id int) (*User, error)
	SaveUser(user *User) error
}

type User struct {
	ID   int
	Name string
}

type UserService struct {
	db Database
}

func (s *UserService) GetUser(id int) (*User, error) {
	// TODO: Get user from database: return s.db.GetUser(id)
	return nil, nil
}

func (s *UserService) CreateUser(name string) (*User, error) {
	// TODO: Create user:
	// user := &User{Name: name}
	// err := s.db.SaveUser(user)
	// if err != nil {
	//     return nil, err
	// }
	// return user, nil
	return nil, nil
}
