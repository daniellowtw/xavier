package feed

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var cfgFile string

var RootCmd = &cobra.Command{
	Use:   "feed",
	Short: "Commands related to feed",
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
