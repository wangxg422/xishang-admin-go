package system

import (
	"backend/model/common/request"
	"time"
)

type SysConfigQuery struct {
	request.PageInfo
	ConfigName string    `json:"configName"`
	ConfigKey  string    `json:"configKey"`
	ConfigType string    `json:"configType"`
	BeginTime  time.Time `json:"beginTime"`
	EndTime    time.Time `json:"endTime"`
}
