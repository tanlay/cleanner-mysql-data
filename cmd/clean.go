package cmd

import (
	"context"
	"github.com/spf13/cobra"
	"github.com/tanlay/cleanner-mysql-data/conf"
	"github.com/tanlay/cleanner-mysql-data/pkg/controller"
)

var date string

func init() {
	RootCmd.AddCommand(CleanData())
}

func CleanData() *cobra.Command {
	cleanCmd := &cobra.Command{
		Use:   "hand-clean",
		Short: "手动清理mysql数据",
		Long:  "手动清理mysql数据",
		RunE: func(cmd *cobra.Command, args []string) error {
			tasks := []func(conf conf.Conf) error{
				SetupJudgeryDB,
			}
			for _, task := range tasks {
				if err := task(*conf.C); err != nil {
					return err
				}
			}
			svc := controller.NewCleanTask(context.Background(), *conf.C)
			return svc.CleanDataTask(date)
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return conf.LoadConfigFromToml(cfgFile)
		},
	}
	cleanCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file path")
	cleanCmd.PersistentFlags().StringVarP(&date, "date", "d", "0000-01-01 00:00:00", "eg: 2006-01-02 15:04:05")
	return cleanCmd
}
