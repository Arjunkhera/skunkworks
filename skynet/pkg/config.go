package root

type MongoConfig struct {
	Ip     string `json:"ip"`
	DbName string `json:"dbname"`
}

type Config struct {
	Mongo *MongoConfig `json:"mongo"`
}
