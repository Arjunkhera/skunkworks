package mongo

import (
	"encoding/json"
	"io/ioutil"
	root "wallet/pkg"
)

// UserService handles the interface of root.User
type UserService struct {
	filePath string
}

// NewUserService creates an instance of  UserService
func NewUserService(path string) *UserService {
	return &UserService{filePath: path}
}

// CreateUser stores credentials in the form of json
func (userServ *UserService) CreateUser(u *root.User) error {
	user, err := newUserModel(u)
	if err != nil {
		return err
	}

	file, err := json.MarshalIndent(user, "", " ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(userServ.filePath, file, 0644)
	return err
}

// Login verifies the authorisation details
func (userServ *UserService) Login(user root.User) (bool, error) {
	model := userModel{}

	// read the boot file
	file, err := ioutil.ReadFile(userServ.filePath)
	if err != nil {
		return false, err
	}

	// decode the json data
	err = json.Unmarshal([]byte(file), &model)
	if err != nil {
		return false, err
	}

	if model.UserName != user.UserName {
		return false, nil
	}

	err = model.comparePassword(user.Password)
	if err != nil {
		return false, err
	}

	return true, nil
}
