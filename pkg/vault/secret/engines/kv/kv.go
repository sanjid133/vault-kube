package kv

import (
	"fmt"
	"github.com/hashicorp/vault/api"
	"github.com/sanjid133/vault-kube/errors"
)

const (
	Name      = "KV"
	KeySecret = "secret"
	KeyPath   = "path"
)

type Info struct {
	path   string
	secret string
	client *api.Client
}

func NewManager() (*Info, error) {
	return &Info{}, nil
}

func (i *Info) Initialize(c *api.Client, vals map[string]string) error {
	if c == nil {
		return errors.ErrMissingValue("vault api client")
	}
	i.client = c
	v, find := vals[KeySecret]
	if !find {
		return errors.ErrMissingValue(KeySecret)
	}
	i.secret = v
	p, find := vals[KeyPath]
	if !find {
		return errors.ErrMissingValue(KeyPath)
	}
	i.path = p
	return nil
}

func (i *Info) GetSecret() (*api.Secret, error) {
	v1Path := func() string {
		return fmt.Sprintf("/v1/%s/%s", i.path, i.secret)
	}
	fmt.Println(v1Path())

	req := i.client.NewRequest("GET", v1Path())
	resp, err := i.client.RawRequest(req)
	if err != nil {
		return nil, errors.PrepareError("failed to get secret", err)
	}
	defer resp.Body.Close()

	sec, err := api.ParseSecret(resp.Body)
	if err != nil {
		return nil, errors.PrepareError("failed to parse secret", err)
	}
	return sec, nil
}
