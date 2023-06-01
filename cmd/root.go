package cmd

import (
	"github.com/spf13/cobra"
)

var cfgFile string

var RootCmd = &cobra.Command{
	Use:   "cleanner-mysql-data",
	Short: "清理mysql数据",
	Long:  "清理mysql数据",
}
