// Code generated by goctl. DO NOT EDIT.
package types

type AddCronjobReq struct {
	Title    string `json:"title,option"      validate:"min=4,max=50"            label:"任务名称"`
	Uuid     string `json:"uuid,option"       validate:"min=4,max=50"            label:"UUID"`
	Ident    string `json:"ident"      validate:"min=2,max=50"            label:"任务标识"`
	Crontab  string `json:"crontab"    label:"定时参数"`
	Status   int64  `json:"status"     validate:"number,gte=0,lte=1"      label:"状态"`
	OrderNum int64  `json:"orderNum,option"   validate:"number,gte=0,lte=9999"   label:"排序"`
	Remark   string `json:"remark,option"     validate:"max=200"                 label:"备注"`
}

type AutoCurd struct {
	Id    int64  `json:"id"`
	Uuid  string `json:"uuid"`
	Title string `json:"title"`
}

type Cronjob struct {
	Id       int64  `json:"id"`
	Uuid     string `json:"uuid"`
	Title    string `json:"title"`
	Ident    string `json:"ident"`
	Crontab  string `json:"crontab"`
	Status   int64  `json:"status"`
	OrderNum int64  `json:"orderNum"`
	Remark   string `json:"remark"`
}

type CronjobListReq struct {
	Title string `from:"title,option"`
}

type CronjobListResp struct {
	List []Cronjob `json:"list"`
}

type CronjobPageReq struct {
	Title string `from:"title,option"`
	PageReq
}

type CronjobPageResp struct {
	List       []Cronjob  `json:"list"`
	Pagination Pagination `json:"pagination"`
}

type DeleteCronjobReq struct {
	Id int64 `json:"id"  validate:"number,gte=1" label:"定时任务id"`
}

type PageReq struct {
	Page  int64 `form:"page"  validate:"number,gte=1"  label:"页数"`
	Limit int64 `form:"limit" validate:"number,gte=1"  label:"条数"`
}

type Pagination struct {
	Page  int64 `json:"page"`
	Limit int64 `json:"limit"`
	Total int64 `json:"total"`
}

type TdFirm struct {
	FirmId    int64  `json:"firm_id"`
	FirmName  string `json:"firm_name"`
	FirmAlias string `json:"firm_alias"`
	FirmCode  string `json:"firm_code"`
	FirmDesc  string `json:"firm_desc"`
}

type TdFirmCreateReq struct {
	FirmName  string `json:"firm_name"`
	FirmAlias string `json:"firm_alias"`
	FirmCode  string `json:"firm_code"`
	FirmDesc  string `json:"firm_desc"`
}

type TdFirmCreateResp struct {
	FirmId    int64  `json:"firm_id"`
	FirmName  string `json:"firm_name"`
	FirmAlias string `json:"firm_alias"`
	FirmCode  string `json:"firm_code"`
	FirmDesc  string `json:"firm_desc"`
}

type TdFirmDeleteReq struct {
	FirmId int64 `form:"firm_id"`
}

type TdFirmDeleteResp struct {
	FirmId int64 `json:"firm_id"`
}

type TdFirmDeletesReq struct {
	FirmId []int64 `form:"[]firm_id"`
}

type TdFirmDeletesResp struct {
	FirmId int64 `json:"[]firm_id"`
}

type TdFirmDetailReq struct {
	FirmId int64 `form:"firm_id"`
}

type TdFirmDetailResp struct {
	FirmId    int64  `json:"firm_id"`
	FirmName  string `json:"firm_name"`
	FirmAlias string `json:"firm_alias"`
	FirmCode  string `json:"firm_code"`
	FirmDesc  string `json:"firm_desc"`
}

type TdFirmListReq struct {
	PageReq
	FirmName  *string `form:"firm_name,option"`
	FirmAlias *string `form:"firm_alias,option"`
	FirmCode  *string `form:"firm_code,option"`
	FirmDesc  *string `form:"firm_desc,option"`
}

type TdFirmListResp struct {
	List  []*TdFirm `json:"list"`
	Total int64     `json:"total"`
}

type TdFirmPageReq struct {
	PageReq
	FirmName  *string `form:"firm_name,option"`
	FirmAlias *string `form:"firm_alias,option"`
	FirmCode  *string `form:"firm_code,option"`
	FirmDesc  *string `form:"firm_desc,option"`
}

type TdFirmPageResp struct {
	List       []*TdFirm  `json:"list"`
	Pagination Pagination `json:"pagination"`
}

type TdFirmUpdateReq struct {
	FirmId    int64  `json:"firm_id"`
	FirmName  string `json:"firm_name"`
	FirmAlias string `json:"firm_alias"`
	FirmCode  string `json:"firm_code"`
	FirmDesc  string `json:"firm_desc"`
}

type TdFirmUpdateResp struct {
	FirmId    int64  `json:"firm_id"`
	FirmName  string `json:"firm_name"`
	FirmAlias string `json:"firm_alias"`
	FirmCode  string `json:"firm_code"`
	FirmDesc  string `json:"firm_desc"`
}

type UpdateCronjobReq struct {
	Id       int64  `json:"id"  label:"任务名称ID"`
	Title    string `json:"title,option"  validate:"min=4,max=50"            label:"任务名称"`
	Uuid     string `json:"uuid,option"  validate:"min=4,max=50"             label:"UUID"`
	Ident    string `json:"ident"  validate:"min=2,max=50"               label:"任务标识"`
	Crontab  string `json:"crontab"                                             label:"定时参数"`
	Status   int64  `json:"status" validate:"number,gte=0,lte=1"         label:"状态"`
	OrderNum int64  `json:"orderNum,option" validate:"number,gte=0,lte=9999"   label:"排序"`
	Remark   string `json:"remark,option" validate:"max=200"                 label:"备注"`
}

type ViewCronjobReq struct {
	Id int64 `json:"id"  validate:"number,gte=1" label:"定时任务id"`
}

type ViewCronjobResp struct {
	Cronjob
}
