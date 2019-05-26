package mongo

import (
	root "skynet/pkg"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type recordModel struct {
	Identifier string `json:"ID"`

	PublicKey  string `json:"PubKey"`
	CommonName string `json:"CommonName"`
}

func newRecordModel(rec *root.Record) {*recordModel, error) {
	user := userModel{UserName: u.UserName}
	err := user.setSaltedPassword(u.Password)
	if err != nil {
		return nil, err
	}

	user.Identifier, err = generateUniqueIdentifier(user.UserName, user.PasswordHash, user.Salt)

	return &user, err
}

/*
type miscellaneousDataModel struct {
	Identifier string `json:"ID"`

	RandomData  string `json:"Key"`
	RandomValue string `json:"Value"`
}
*/
