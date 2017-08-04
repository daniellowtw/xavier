package cmd

import (
	"github.com/daniellowtw/xavier/api"
	"github.com/spf13/cobra"
)

func NewLearnCmd() *cobra.Command {
	return &cobra.Command{
		Use: "learn",
		RunE: func(cmd *cobra.Command, _ []string) error {
			e, err := newDBClientFromCmd(cmd)
			if err != nil {
				return err
			}
			if err := e.PopulateProcessQueue(); err != nil {
				return err
			}
			s := api.NewService(e)
			return s.LearnFromNewNews()
		},
	}
}
