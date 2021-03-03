package models

import "gorm.io/gorm"

// Config is based off config.yml.
type Config struct {
	OwnerID          string   `yaml:"ownerID"`
	APIBibleKey      string   `yaml:"apiBibleKey"`
	IsDryRun         bool     `yaml:"dry"`
	LetsEncryptEmail string   `yaml:"letsEncryptEmail"`
	AccessKey        string   `yaml:"accessKey"`
	DBHost           string   `yaml:"databaseHost"`
	DBPort           int      `yaml:"databasePort"`
	DBUser           string   `yaml:"databaseUser"`
	DBPass           string   `yaml:"databasePass"`
	DecryptionKey    string   `yaml:"decryptionKey"`
	EncryptedFiles   []string `yaml:"encryptedFiles"`
	DB               gorm.DB
}
