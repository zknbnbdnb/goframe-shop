package consts

const (
	ProjectName              = "goFrame-shop"
	ProjectUsage             = "zk"
	ProjectBrief             = "start http server"
	Version                  = "v0.2.0"             // 当前服务版本(用于模板展示)
	CaptchaDefaultName       = "CaptchaDefaultName" // 验证码默认存储空间名称
	ContextKey               = "ContextKey"         // 上下文变量存储键名，前后端系统共享
	FileMaxUploadCountMinute = 10                   // 同一用户1分钟之内最大上传数量
	GTokenBackendPrefix      = "Admin:"             // gtoken登陆 管理后台 前缀区分
	GTokenFrontendPrefix     = "User:"              // gtoken登陆 前台 前缀区分
	CtxAdminId               = "CtxAdminId"         // for admin
	CtxAdminName             = "CtxAdminName"
	CtxAdminIsAdmin          = "CtxAdminIsAdmin"
	CtxAdminRoleIds          = "CtxAdminRoleIds"
	CtxUserId                = "CtxUserId"
	CtxUserName              = "CtxUserName" // for user
	CtxUserAvatar            = "CtxUserAvatar"
	CtxUserSex               = "CtxUserSex"
	CtxUserSign              = "CtxUserSign"
	CtxUserStatus            = "CtxUserStatus"
	CodeMissingParameterMsg  = "缺少参数"
	CacheModeRedis           = 2
	BackendServerName        = "goFrame-shop"
	BackendMultiLogin        = true
	FrontendMultiLogin       = false
	GTokenExpireIn           = 86400 //10 * 24 * 60 * 60单位秒
	TokenType                = "Bearer"
	ErrorSecretAnswerMsg     = "密保答案错误"
	ErrLoginFailMsg          = "登录失败, 账号或密码错误."
)
