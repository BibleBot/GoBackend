package models

// Config is based off config.yml.
type Config struct {
	OwnerID          string   `yaml:"ownerID"`
	APIBibleKey      string   `yaml:"apiBibleKey"`
	IsDryRun         bool     `yaml:"dry"`
	AccessKey        string   `yaml:"accessKey"`
	LetsEncryptEmail string   `yaml:"letsEncryptEmail"`
	DecryptionKey    string   `yaml:"decryptionKey"`
	EncryptedFiles   []string `yaml:"encryptedFiles"`
}
