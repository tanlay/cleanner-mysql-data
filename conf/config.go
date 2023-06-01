package conf

import (
	"errors"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/tanlay/cleanner-mysql-data/pkg/lib/gormx"
	"go.uber.org/zap"
)

var C = &Conf{}

type Conf struct {
	PerTotalCount int                `json:"per_total_count" toml:"per_total_count"`
	BatchCount    int                `json:"batch_count" toml:"batch_count"`
	GoCount       int                `json:"go_count" toml:"go_count"`
	IntervalTime  int                `json:"interval_time" toml:"interval_time"`
	Logger        LoggerConf         `json:"logger" toml:"logger"`
	Database      gormx.DatabaseConf `json:"database" toml:"database"`
}

type LoggerConf struct {
	Env    string `json:"env" toml:"env"`
	Level  string `json:"level" toml:"level"`
	OutPut string `json:"output" toml:"output"`
}

func LoadConfigFromToml(cfgFile string) error {
	if cfgFile == "" {
		return errors.New("未指定配置文件")
	} else {
		if _, err := toml.DecodeFile(cfgFile, C); err != nil {
			return err
		}
	}
	//输出日志到指定文件
	zapLevel := zap.NewAtomicLevel()
	if err := zapLevel.UnmarshalText([]byte(C.Logger.Level)); err != nil {
		panic(err.Error())
	}
	var zapConf zap.Config
	if env := C.Logger.Env; env == "dev" {
		zapConf = zap.NewDevelopmentConfig()
	} else {
		zapConf = zap.NewProductionConfig()
	}
	if C.Logger.OutPut != "" {
		zapConf.OutputPaths = []string{C.Logger.OutPut}
		zapConf.ErrorOutputPaths = []string{C.Logger.OutPut}
	}

	if logger, err := zapConf.Build(); err != nil {
		panic(err.Error())
	} else {
		zap.RedirectStdLog(logger)
		zap.ReplaceGlobals(logger)
	}
	zap.L().Info(fmt.Sprintf("load config：%s", cfgFile))
	return nil
}
