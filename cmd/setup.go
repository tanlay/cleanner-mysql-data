package cmd

import (
	"github.com/tanlay/cleanner-mysql-data/conf"
	"github.com/tanlay/cleanner-mysql-data/pkg/lib/gormx"
)

func SetupJudgeryDB(conf conf.Conf) error {
	return gormx.ConnectJudgeryDB(conf.Database)
}
