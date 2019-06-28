package main

import (
	"github.com/gomsa/tools/config"
)

// Conf 配置
var Conf config.Config = config.Config{
	Service: "nats-api",
	Version: "latest",
	Permissions: []config.Permission{
		{Service: "nats-api", Method: "Templates.Create", Auth: true, Policy: true, Name: "创建事件模板", Description: "创建新事件模板权限。"},
		{Service: "nats-api", Method: "Templates.Delete", Auth: true, Policy: true, Name: "删除事件模板", Description: "删除事件模板。"},
		{Service: "nats-api", Method: "Templates.Update", Auth: true, Policy: true, Name: "更新事件模板", Description: "更新事件模板信息。"},
		{Service: "nats-api", Method: "Templates.Get", Auth: true, Policy: true, Name: "查询事件模板", Description: "查询事件模板信息权限。"},
		{Service: "nats-api", Method: "Templates.List", Auth: true, Policy: true, Name: "事件模板列表", Description: "查询事件模板列表"},
		{Service: "nats-api", Method: "Templates.All", Auth: true, Policy: true, Name: "全部事件模板", Description: "查询全部事件模板。"},

		{Service: "nats-api", Method: "Nats.MobileVerify", Auth: true, Policy: false, Name: "手机验证码发送", Description: "发送手机验证码"},
	},
}
