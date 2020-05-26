package model

import "time"

type LogConfig struct {
	LogLevel     string `json:"log_level"`
	MaxAge       int    `json:"max_age"`
	RotationTime int    `json:"rotation_time"`
}

type Group struct {
	Id           int       `json:"id" form:"id"`
	CompanyId    int       `json:"companyId" form:"companyId"`
	CompanyName  string    `json:"companyName" form:"companyName"`
	Name         string    `json:"name" form:"name"`
	Status       int       `json:"status" form:"status"`
	DefaultGroup int       `json:"defaultGroup" form:"defaultGroup"`
	Remark       string    `json:"remark" form:"remark"`
	CreatorId    int       `json:"creatorId" form:"creatorId"`
	CreatorName  string    `json:"creatorName" form:"creatorName"`
	Created      time.Time `json:"created" form:"created"`
	CorpId       int       `json:"corpId"`
	CorpName     string    `json:"corpName"`
}
