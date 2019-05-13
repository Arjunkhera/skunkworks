package root

type Record struct {
	Identifier string `json:"ID"`
	PublicKey  string `json:"PubKey"`
	Password   string `json:"Psw"`
	CommonName string `json:"CommonName"`
}

type RecordService interface {
	CreateRecord(rec *Record) error
	GetRecordByIdentifier(identifier string) (error, Record)
}
