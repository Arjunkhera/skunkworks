package root

type User struct {
	Identifier string `json:"ID"`

	UserName string `json:"UserName"`
	Password string `json:"PSW"`
}

type UserService interface {
	CreateUser(u *User) error
	GetUserByUsername(username string) (User, error)
	Login(cred Credentials) (User, error, bool)
}

type Record struct {
	Identifier string `json:"ID"`

	PublicKey  string `json:"PubKey"`
	CommonName string `json:"CommonName"`
}

type RecordService interface {
	CreateRecord(rec *Record) error
}

/*
type MiscellaneousData struct {
	Identifier string `json:"ID"`

	RandomData  string `json:"Key"`
	RandomValue string `json:"Value"`
}

type MiscellaneousDataService interface {
}


type Claim struct {
	RecordIdentifier string `json:"ID"`
}

type ClaimService interface {
}
*/
