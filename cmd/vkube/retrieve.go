package main

import (
	"fmt"
	"github.com/sanjid133/vault-kube/config"
	"github.com/sanjid133/vault-kube/pkg/k8s"
	"github.com/sanjid133/vault-kube/pkg/vault"
	"github.com/spf13/cobra"
	"log"
)

var retrieveCmd = &cobra.Command{
	Use:   "authenticate",
	Short: "authenticate kubernetes login",
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
	fmt.Println(token)
	return nil
}
