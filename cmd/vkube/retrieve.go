package main

import (
	"github.com/sanjid133/vault-kube/config"
	"github.com/sanjid133/vault-kube/pkg/k8s"
	"github.com/sanjid133/vault-kube/pkg/vault"
	"github.com/sanjid133/vault-kube/pkg/vault/secret"
	"github.com/spf13/cobra"
	"log"
)

var retrieveCmd = &cobra.Command{
	Use:   "retrieve",
	Short: "retrieve data",
	RunE:  retrieve,
}

func retrieve(cmd *cobra.Command, args []string) error {
	cfg, err := config.LoadFromEnvironment()
	if err != nil {
		log.Fatal(err)
	}
	c, err := vault.New(cfg)
	if err != nil {
		log.Fatal(err)
	}
	k, err := k8s.New(c.Client, cfg)
	if err != nil {
		log.Fatal(err)
	}
	token, err := k.Login()
	if err != nil {
		// TODO: allow fail
		log.Fatal(err)
	}
	c.Client.SetToken(token)

	opts := map[string]string{
		"secret": cfg.SecretName,
		"path":   cfg.SecretPath,
	}
	sec, err := secret.RetrieveSecret(c.Client, cfg.SecretEngine, opts)
	if err != nil {
		log.Fatal(err)
	}

	if cfg.SecretKey == "" {
		for k := range sec.Data {
			if err := secret.StoreSecretIntoFile(sec, cfg.DataPath, k); err != nil {
				log.Fatal(err)
			}
		}
	} else {
		if err := secret.StoreSecretIntoFile(sec, cfg.DataPath, cfg.SecretKey); err != nil {
			log.Fatal(err)
		}
	}
	return nil
}
