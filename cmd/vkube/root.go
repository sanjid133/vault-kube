package main

import (
	"github.com/sanjid133/vault-kube/version"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "vkube",
	Short:   "vkube retrieve data from vault into kubernetes",
	Version: version.Version,
}

func main() {
	_ = rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(
		authCmd,
		retrieveCmd,
	)
}
