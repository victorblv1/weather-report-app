package domain

type User struct {
    ID          string
    Name        string
    Preferences map[string]string
}

type UserRepository interface {
    FetchUser(id string) (*User, error)
    SaveUser(user *User) error
}