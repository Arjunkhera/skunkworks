package root

// MongoConfig stores config data for connecting to mongo
type MongoConfig struct {
	Ip     string `json:"ip"`
	DbName string `json:"dbName"`
}

// ServerConfig stores config data for starting server
type ServerConfig struct {
	Port     string `json:"port"`
	SkynetIP string `json:"SkynetIP"`
}

// BootConfig stores config data to determing startup procedures
type BootConfig struct {
	BootConfigExists bool   `json:"bootconfig"`
	BootConfigPath   string `json:"bootconfigpath"`
	DeviceConfigPath string `json:"deviceconfigpath"`
}

// Config stores pointers to all other config structs
type Config struct {
	Mongo  *MongoConfig  `json:"mongo"`
	Server *ServerConfig `json:"server"`
	Boot   *BootConfig   `json:"boot"`
}
