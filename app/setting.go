package app

type SettingTable struct {
	Key   string `gorm:"primaryKey"`
	Value string `gorm:"column:json_value"`
}

func (SettingTable) TableName() string {
	return "settings"
}

func (SettingTable) KeyField() string {
	return "key"
}

func (SettingTable) ValueField() string {
	return "json_value"
}

func (SettingTable) MigrationKey() string {
	return "hrm_table_versions"
}

func (SettingTable) SeedKey() string {
	return "hrm_executed_seeds"
}

type CompanySettingTable struct {
	SettingTable
}

func (CompanySettingTable) TableName() string {
	return "inifiles"
}

type GeneralResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
