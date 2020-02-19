package config

import (
	"github.com/sanjid133/vault-kube/errors"
	"os"
)

const (
	KeyVaultRole               = "VAULT_ROLE"
	KeyDataPath                = "VAULT_DATA_PATH"
	KeyServiceAccountTokenPath = "SERVICE_ACCOUNT_TOKEN_PATH"
	KeySecretName              = "VAULT_SECRET_NAME"
	KeySecretPath              = "VAULT_SECRET_PATH"
	KeySecretKey               = "VAULT_SECRET_KEY"
	KeySecretEngine            = "VAULT_SECRET_ENGINE"
)

const (
	defaultSecretEngine            = "KV"
	defaultServiceAccountTokenPath = "/var/run/secrets/kubernetes.io/serviceaccount/token"
)

type Config struct {
	Role                    string
	DataPath                string
	ServiceAccountTokenPath string
	SecretEngine            string
	SecretName              string
	SecretKey               string
	SecretPath              string
}

func LoadFromEnvironment() (*Config, error) {
	c := &Config{}
	c.Role = os.Getenv(KeyVaultRole)
	c.DataPath = os.Getenv(KeyDataPath)
	if c.DataPath == "" {
		return nil, errors.ErrMissingValue(KeyDataPath)
	}
	c.SecretEngine = os.Getenv(KeySecretEngine)
	if c.SecretEngine == "" {
		c.SecretEngine = defaultSecretEngine
	}
	c.SecretName = os.Getenv(KeySecretName)
	if c.SecretName == "" {
		return nil, errors.ErrMissingValue(KeySecretName)
	}
	c.SecretKey = os.Getenv(KeySecretKey)
	c.SecretPath = os.Getenv(KeySecretPath)
	if c.SecretPath == "" {
		return nil, errors.ErrMissingValue(KeySecretPath)
	}
	c.ServiceAccountTokenPath = os.Getenv(KeyServiceAccountTokenPath)
	if c.ServiceAccountTokenPath == "" {
		c.ServiceAccountTokenPath = defaultServiceAccountTokenPath
	}
	return c, nil
}
