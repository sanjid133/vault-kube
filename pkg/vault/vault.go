package vault

import (
	"github.com/hashicorp/vault/api"
	"github.com/sanjid133/vault-kube/config"
	"github.com/sanjid133/vault-kube/errors"
)

type Vault struct {
	cfg    *config.Config
	Client *api.Client
}

type vaultLogicalWriter interface {
	Write(path string, data map[string]interface{}) (*api.Secret, error)
}

var vaultLogical = func(c *api.Client) vaultLogicalWriter {
	return c.Logical()
}

func New(cfg *config.Config) (*Vault, error) {
	v := &Vault{
		cfg: cfg,
	}

	vaultConfig := api.DefaultConfig()
	if err := vaultConfig.ReadEnvironment(); err != nil {
		return nil, errors.PrepareError("failed to read environment for vault", err)
	}
	var err error
	v.Client, err = api.NewClient(vaultConfig)
	if err != nil {
		return nil, errors.PrepareError("failed to create vault Client", err)
	}
	return v, nil
}
