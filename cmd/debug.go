package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var (
	DebugCmd = &cobra.Command{
		Use: "debug <id>",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return fmt.Errorf("Need one argument: <urL>")
			}
			u := args[0]
			i, err := strconv.ParseInt(u, 10, 64)
			if err != nil {
				return err
			}
			data, err := s.DebugFeed(i)
			fmt.Printf("%s\n", data)
			return nil
		},
	}
)
