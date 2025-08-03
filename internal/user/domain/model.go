package domain

type User struct {
	ID          string
	Name        string
	Preferences map[string]string
}