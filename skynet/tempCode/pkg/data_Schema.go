package root

type User struct {
	Identifier string `json:"ID"`

	UserName string `json:"UserName"`
	Password string `json:"PSW"`
}

type UserService interface {
	CreateUser(u *User) error
	GetUserByUsername(username string) (error, User)
}

type Record struct {
	Identifier string `json:"ID"`
	PublicKey  string `json:"PubKey"`
	CommonName string `json:"CommonName"`
}

type RecordService interface {
	CreateRecord(rec *Record) error
	GetRecordByIdentifier(identifier string) (*Record, error)
}

type Claim struct {
	RecordIdentifier string `json:"ID"`
}

type ClaimService interface {
}
