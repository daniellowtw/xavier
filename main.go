package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/daniellowtw/xavier/cmd"
)

func main() {
	rootCmd := &cobra.Command{
		Use: os.Args[0],
	}
	rootCmd.AddCommand(cmd.AddCmd, cmd.UpdateAllCmd, cmd.ListAllCmd, cmd.DeleteCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
