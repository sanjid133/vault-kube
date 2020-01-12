package secret

import "github.com/hashicorp/vault/api"

const (
	KeySecret = "secret"
	KeyPath = "path"
)

type SecretManager interface {
	Initialize(*api.Client, map[string]string) error
	GetSecret() (*api.Secret, error)
}
