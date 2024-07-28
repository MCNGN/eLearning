package domain

type User struct {
	ID   string
	Name string
	Email string
	Username string
	Password string
	LoginStatus bool
}

func NewUser(id string, name string, password string) *User {
	return &User{
		ID:   id,
		Name: name,
		Password: password,
		LoginStatus: false,
	}
}