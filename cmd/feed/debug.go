package feed

import (
	"fmt"

	"github.com/daniellowtw/xavier/cmd/service"
	"github.com/spf13/cobra"
)

var (
	debugCmd = &cobra.Command{
		Use: "debug <id>",
		RunE: func(cmd *cobra.Command, args []string) error {
			i, err := extractInt1(args, "feed ID")
			if err != nil {
				return err
			}
			s, err := service.NewServiceFromCmd(cmd)
			if err != nil {
				return err
			}
			data, err := s.DebugFeed(i)
			fmt.Printf("%s\n", data)
			return nil
		},
	}
)

func init() {
	RootCmd.AddCommand(debugCmd)
}
