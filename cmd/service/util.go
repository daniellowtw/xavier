/*
This file provides convenient function to produce structs that are configured by the flags
*/
package service

import (
	"github.com/daniellowtw/xavier/api"
	"github.com/daniellowtw/xavier/db"
	"github.com/spf13/cobra"
	"xorm.io/xorm/log"
)

func NewDBClientFromCmd(cmd *cobra.Command) (*db.Client, error) {
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
	ll := log.LogLevel(logLevel)
	return db.NewSqlite3Client(dbName, showSql, ll)
}

func NewServiceFromCmd(cmd *cobra.Command) (*api.Service, error) {
	c, err := NewDBClientFromCmd(cmd)
	if err != nil {
		return nil, err
	}
	return api.NewService(c), nil
}
