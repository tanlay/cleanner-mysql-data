package model

type DataReport struct {
	Id             int64  `json:"id"`
	Cid            string `json:"cid"`
	CidType        int    `json:"cid_type"`
	Name           string `json:"name"`
	Phone          string `json:"phone"`
	DataSourceCode string `json:"data_source_code"`
	Data           string `json:"data"`
	Timestamp      int64  `json:"timestamp"`
	DataId         string `json:"data_id"`
	CreatedAt      int64  `json:"created_at"`
	DataGenTime    int64  `json:"data_gen_time"`
}

func (d *DataReport) TableName() string {
	return "data_report"
}
