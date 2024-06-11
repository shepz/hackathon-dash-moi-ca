package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func AddCreateCmd(root *cobra.Command) {
	root.AddCommand(createCmd)
}

var createCmd = &cobra.Command{
	Use:   "create -f {filename}",
	Short: "Create dashboard from a YAML file",
	Run: func(cmd *cobra.Command, args []string) {
		filename, _ := cmd.Flags().GetString("filename")
		if filename == "" {
			fmt.Println("Please specify a filename using -f")
			return
		}
	},
}

func init() {
	createCmd.Flags().StringP("filename", "f", "", "Specify the YAML file")
}
