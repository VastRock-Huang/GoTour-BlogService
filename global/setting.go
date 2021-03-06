package global

import (
	"github.com/vastrock-huang/gotour-blogservice/pkg/logger"
	"github.com/vastrock-huang/gotour-blogservice/pkg/setting"
)

var (
	//全局设置
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	//全局日志
	Logger          *logger.Logger
	//JWT认证
	JWTSetting *setting.JWTSettingS
	//邮件配置
	EmailSetting *setting.EmailSettingS
)
