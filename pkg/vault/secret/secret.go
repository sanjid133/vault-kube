package secret

import (
	"github.com/hashicorp/vault/api"
	"github.com/sanjid133/vault-kube/pkg/vault/secret/engines"
	"github.com/sanjid133/vault-kube/util"
	"path/filepath"
)

func RetrieveSecret(client *api.Client, engineName string, opts map[string]string) (*api.Secret, error) {
	engine, err := engines.NewEngineManager(engineName)
	if err != nil {
		return nil, err
	}
	if err := engine.Initialize(client, opts); err != nil {
		return nil, err
	}
	sec, err := engine.GetSecret()
	if err != nil {
		return nil, err
	}
	return sec, nil
}

func StoreSecretIntoFile(sec *api.Secret, filePath, keyName string) error {
	dir := filepath.Dir(filePath)
	if err := util.EnsurePath(dir); err != nil {
		return err
	}
	if val, found := sec.Data[keyName]; found {
		file := filepath.Join(filePath, keyName)
		if err := util.WriteData(file, val); err != nil {
			return err
		}
	}
	return nil
}
