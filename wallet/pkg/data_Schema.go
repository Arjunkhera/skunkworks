package root

// User struct is the base struct for User
type User struct {
	UserName string `json:"UserName"`
	Password string `json:"PSW"`
}

// UserService defines the valid operations on User struct
type UserService interface {
	CreateUser(u *User) (string, error)
	Login(u User) (bool, error)
}

// Device struct is the base struct for Wallet Id
type Device struct {
	Identifier string `json:"Identifier"`

	PublicKey string `json:"PubKey"`
}

// DeviceService defines the valid operations on Device struct
type DeviceService interface {
	CreateDevice(d *Device) error
}

// PairIdentity is used to store the pairwise dids for the user
type PairIdentity struct {
	Identifier string `json:"Identifier"`

	UserName  string `json:"UserName"`
	PublicKey string `json:"PubKey"`
}

// PairIdentityService defines the valid operations on PairwiseIdentity struct
type PairIdentityService interface {
	CreatePairIdentity(username string) error
	GetPairIdentityByUsername(username string) (PairIdentity, error)
	GetAllPairIdentities() ([]PairIdentity, error)
}
