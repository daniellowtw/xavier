package feed

import (
	"fmt"
	"strconv"

	"github.com/daniellowtw/xavier/cmd/service"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use: "delete <id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		i, err := extractInt1(args, "feed ID")
		if err != nil {
			return err
		}
		s, err := service.NewServiceFromCmd(cmd)
		if err != nil {
			return err
		}
		return s.DeleteFeed(i)
	},
}

func init() {
	RootCmd.AddCommand(deleteCmd)
}

// returns the int
func extractInt1(args []string, name string) (int64, error) {
	if len(args) != 1 {
		return 0, fmt.Errorf("need one argument: <%s>", name)
	}
	u := args[0]
	return strconv.ParseInt(u, 10, 64)
}
