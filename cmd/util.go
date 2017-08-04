/*
This file provides convenient function to produce structs that are configured by the flags
*/
package cmd

import (
	"github.com/daniellowtw/xavier/api"
	"github.com/daniellowtw/xavier/db"
	"github.com/go-xorm/core"
	"github.com/spf13/cobra"
)

func newDBClientFromCmd(cmd *cobra.Command) (*db.Client, error) {
	dbName, err := cmd.Flags().GetString("db")
	if err != nil {
		return nil, err
	}
	showSql, err := cmd.Flags().GetBool("show-sql")
	if err != nil {
		return nil, err
	}
	logLevel, err := cmd.Flags().GetInt("log-level")
	if err != nil {
		return nil, err
	}
	ll := core.LogLevel(logLevel)
	return db.NewSqlite3Client(dbName, showSql, ll)
}

func newServiceFromCmd(cmd *cobra.Command) (*api.Service, error) {
	c, err := newDBClientFromCmd(cmd)
	if err != nil {
		return nil, err
	}
	return api.NewService(c), nil
}
