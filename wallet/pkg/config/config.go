package config

import (
	"os"

	root "wallet/pkg"
)

// GetConfig obtains the environment variables for root.configuration
func GetConfig() *root.Config {
	return &root.Config{
		Mongo: &root.MongoConfig{
			Ip:     envOrDefaultString("wallet:mongo:ip", "127.0.0.1:27017"),
			DbName: envOrDefaultString("wallet:mongo:dbName", "wallet")},
		Server: &root.ServerConfig{Port: envOrDefaultString("wallet:server:port", ":8000")},
		Boot: &root.BootConfig{
			BootConfigExists: false,
			BootConfigPath:   os.Getenv("HOME") + "/wallet.txt"}}
}

func CheckBootConfigFile(rt *root.BootConfig) error {
	// check if the boot file exists or not
	_, err := os.Stat(rt.BootConfigPath)
	// already exists
	if err == nil {
		rt.BootConfigExists = true
		return nil
	}

	// does not exist
	file, err := os.Create(rt.BootConfigPath)
	if err != nil {
		return err
	}

	file.Close()
	return nil
}

func envOrDefaultString(envVar string, defaultValue string) string {
	value := os.Getenv(envVar)
	if value == "" {
		return defaultValue
	}

	return value
}
