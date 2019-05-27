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
	GetAllRecords() ([]Record, error)
}

type ClaimDefn struct {
	UserIdentifier string `json:"UID"`
	CommonName     string `json:"CNAME"`

	ClaimDefnIdentifier string            `json:"CDID"`
	AttributesToType    map[string]string `json:"ATTR"`
}

type ClaimService interface {
	CreateClaimDefn(map[string]string) (string, error)
	CreateClaim(string, string, string) error

	GetClaimByUserID(string) ([]Claim, error)
	GetClaimByCommonName(string, string) (Claim, error)
	GetClaimDefnByClaimDefnID(string) ([]ClaimDefn, error)

	GetAllClaims() ([]Claim, error)
	GetAllClaimDefns() ([]ClaimDefn, error)
}

type PairIdentity struct {
	Identifier          string `json:"Identifier"`
	ClaimDefnIdentifier string `json:"CDID"`

	UserName       string `json:"UserName"`
	OtherPartyName string `json:"OtherPartyName`
	PublicKey      string `json:"PubKey"`

	Endpoint   string            `json:"Endpoint"`
	HashedData map[string]string `json:HashedData`
}

type PairIdentityService interface {
	CreatePairIdentity(new PairIdentity) error
}
