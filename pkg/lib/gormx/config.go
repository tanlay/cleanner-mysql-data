package gormx

type DatabaseConf struct {
	JudgeryDSN       string `json:"judgery_dsn" toml:"judgery_dsn"`
	DataReportEnable bool   `json:"data_report_enable" toml:"data_report_enable"`
	MaxIdleConn      int    `json:"max_idle_conn" toml:"max_idle_conn"`
	MaxOpenConn      int    `json:"max_open_conn" toml:"max_open_conn"`
	ConnMaxLiftTime  int    `json:"conn_max_lift_time" toml:"conn_max_lift_time"`
}
