package domain

type Tutor struct {
	User
}

func NewTutor(id string, name string, password string) *Tutor {
	return &Tutor{
		User: User{
			ID:   id,
			Name: name,
			Password: password,
			LoginStatus: false,
		},
	}
}