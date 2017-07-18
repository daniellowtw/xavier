package main

import (
	"fmt"
	"os"

	"github.com/daniellowtw/xavier/cmd"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use: os.Args[0],
	}
	verCmd := &cobra.Command{
		Use:   "version",
		Short: "Print the current version of the program",
		Run: func(*cobra.Command, []string) {
			fmt.Println(version)
		},
	}
	rootCmd.AddCommand(verCmd)
	rootCmd.AddCommand(cmd.AddCmd,
		cmd.UpdateAllCmd,
		cmd.UpdateFeedCmd,
		cmd.ListAllCmd,
		cmd.DeleteCmd,
		cmd.WebCmd,
		cmd.ImportFeedSourceFromFileCmd,
		cmd.AddNewToQueueCmd,
		cmd.LearnCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
