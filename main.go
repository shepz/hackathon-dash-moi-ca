// File: main.go
package main

import (
	"dash-moi-ca/cmd"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "dashmoica",
	Short: "Dashmoica is a CLI tool to create Grafana dashboards",
}

func main() {
	cmd.AddCreateCmd(rootCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
