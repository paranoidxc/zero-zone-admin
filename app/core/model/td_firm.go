package model

type TmpTdFirm struct {
	FirmId    int64  `json:"firmId" gorm:"primaryKey;column:firm_id;type:varchar(36);comment:唯一id;"`
	FirmName  string `json:"firmName" gorm:"column:firm_name;not null;uniqueIndex;type:varchar(255);comment:厂商名称;"`
	FirmAlias string `json:"firmAlias" gorm:"column:firm_alias;type:varchar(255);comment:厂商别名;"`
	FirmCode  string `json:"firmCode" gorm:"column:firm_code;type:varchar(255);comment:厂商编码;"`
	FirmDesc  string `json:"firmDesc" gorm:"column:firm_desc;type:varchar(255);comment:厂商描述;"`
}
