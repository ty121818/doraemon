package common

const (
	AlertMethodSms      = "SMS"
	AlertMethodLanxin   = "LANXIN"
	AlertMethodCall     = "CALL"
	AlertMethodHook     = "HOOK"
	AlertMethodDingTalk = "DINGTALK"
)

const (
	AlertStatusInactive			= 0
	AlertStatusPending	        = 1
	AlertStatusFiring			= 2
)