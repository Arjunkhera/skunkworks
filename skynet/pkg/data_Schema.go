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

type ClaimService interface {
	GetClaimByUserID(string) ([]Claim, error)
	GetClaimDefnByClaimDefnID(string) ([]ClaimDefn, error)
	CreateClaimDefn(map[string]string) (string, error)
	CreateClaim(string, string) error
}
type Record struct {
	Identifier string `json:"ID"`

	PublicKey  string `json:"PubKey"`
	CommonName string `json:"CommonName"`
}

type RecordService interface {
	CreateRecord(rec *Record) error
	GetAllRecords() ([]Record, error)
	//GetRecordByUsername(username string) (Record, error)
}

type Claim struct {
	UserIdentifier      string `json:"UID"`
	ClaimDefnIdentifier string `json:"CDID"`
}

type ClaimDefn struct {
	ClaimDefnIdentifier string            `json:"CDID"`
	AttributesToType    map[string]string `json:"ATTR"`
}

/*
type MiscellaneousData struct {
	Identifier string `json:"ID"`

	RandomData  string `json:"Key"`
	RandomValue string `json:"Value"`
}

type MiscellaneousDataService interface {
}



type ClaimService interface {
}
*/
