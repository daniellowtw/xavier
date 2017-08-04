package main

import (
	"fmt"
	"os"

	"github.com/daniellowtw/xavier/cmd"
	"github.com/spf13/cobra"
)

func main() {
	verCmd := &cobra.Command{
		Use:   "version",
		Short: "Print the current version of the program",
		Run: func(*cobra.Command, []string) {
			fmt.Println(version)
		},
	}
	cmd.RootCmd.AddCommand(verCmd)
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
