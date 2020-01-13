package config

import (
	"github.com/sanjid133/vault-kube/errors"
	"os"
)

const (
	KeyVaultAddr	= "VAULT_ADDR"
	KeyVaultRole               = "VAULT_ROLE"
	KeyDataPath                = "VAULT_DATA_PATH"
	KeyServiceAccountTokenPath = "SERVICE_ACCOUNT_TOKEN_PATH"
)

const (
	defaultServiceAccountTokenPath = "/var/run/secrets/kubernetes.io/serviceaccount/token"
	defaultVaultAddr = "http://127.0.0.1:8200"
)

type Config struct {
	Role                    string
	DataPath                string
	ServiceAccountTokenPath string
}

func LoadFromEnvironment() (*Config, error) {
	c := &Config{}
	c.Role = os.Getenv(KeyVaultRole)
	c.DataPath = os.Getenv(KeyDataPath)
	if c.DataPath == "" {
		return nil, errors.ErrMissingValue(KeyDataPath)
	}
	c.ServiceAccountTokenPath = os.Getenv(KeyServiceAccountTokenPath)
	if c.ServiceAccountTokenPath == "" {
		c.ServiceAccountTokenPath = defaultServiceAccountTokenPath
	}
	return c, nil
}
