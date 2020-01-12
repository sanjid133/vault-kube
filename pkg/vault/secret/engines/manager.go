package engines

import "github.com/hashicorp/vault/api"

type SecretManager interface {
	Initialize(*api.Client, map[string]string) error
	GetSecret() (*api.Secret, error)
}
