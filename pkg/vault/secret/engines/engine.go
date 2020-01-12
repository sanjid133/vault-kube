package engines

import (
	"github.com/sanjid133/vault-kube/errors"
	"github.com/sanjid133/vault-kube/pkg/vault/secret"
	"github.com/sanjid133/vault-kube/pkg/vault/secret/engines/kv"
)

func NewEngineManager(name string) (secret.SecretManager, error) {
	switch name {
		case kv.Name:
			return kv.NewManager()
	default:
		return nil, errors.ErrInvalidSecretEngine(name)
	}

}
