package root

// User struct is the base struct for User
type User struct {
	UserName string `json:"UserName"`
	Password string `json:"PSW"`
}

// UserService defines the valid operations on User struct
type UserService interface {
	CreateUser(u *User) error
	Login(u User) (bool, error)
}
