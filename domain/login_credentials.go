package domain

type LoginCredentials struct {
    Username string
    Password string
}

func NewLoginCredentials(username, password string) LoginCredentials {
    return LoginCredentials{
        Username: username,
        Password: password,
    }
}