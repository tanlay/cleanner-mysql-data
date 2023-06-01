package model

type HealthReport struct {
	Id                  int64  `json:"id"`
	Cid                 string `json:"cid"`
	CidType             int    `json:"cid_type"`
	Name                string `json:"name"`
	Phone               string `json:"phone"`
	LivingProvinceCode  string `json:"living_province_code"`
	LivingCityCode      string `json:"living_city_code"`
	LivingCountryCode   string `json:"living_country_code"`
	LivingDetailAddress string `json:"living_detail_address"`
	LivingRiskLevel     int    `json:"living_risk_level"`
	HousingType         string `json:"housing_type"`
	IsVisitedOverseas   int8   `json:"is_visited_overseas"`
	HighRiskAreaIds     string `json:"high_risk_area_ids"`
	HighestRiskLevel    int    `json:"highest_risk_level"`
	IsContacted         int8   `json:"is_contacted"`
	IsIsolation         int8   `json:"is_isolation"`
	CurrentStatus       string `json:"current_status"`
	CurrentSymptom      string `json:"current_symptom"`
	OtherSymptom        string `json:"other_symptom"`
	CurrentTemperature  string `json:"current_temperature"`
	ApplyCid            string `json:"apply_cid"`
	ApplyCidType        int    `json:"apply_cid_type"`
	CreatedAt           int64  `json:"created_at"`
	ReportTime          int64  `json:"report_time"`
	DataId              string `json:"data_id"`
	Ext                 string `json:"ext"`
}

func (h *HealthReport) TableName() string {
	return "health_report"
}
