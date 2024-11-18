package code

import "im/pkg/codex"

var (
	PasswordEmpty        = codex.New(20001, "请输入密码")
	NewPasswordEmpty     = codex.New(20002, "请输入新密码")
	ConfirmPasswordEmpty = codex.New(20003, "请输入确认密码")
	NotMatchasswordEmpty = codex.New(20004, "密码不一致")
	EmailEmpty           = codex.New(20005, "请输入邮箱")
	CodeEmpty            = codex.New(20006, "请输入验证码")
	MobileEmpty          = codex.New(20007, "请输入手机号")
	NicknameEmpty        = codex.New(20007, "请输入昵称")

	ChangePasswordFailed = codex.New(20101, "修改密码失败")
	CodeIncorrect        = codex.New(20102, "验证码错误")
	ChangeEmailFailed    = codex.New(20103, "修改邮箱失败")
	ChangeMobileFailed   = codex.New(20104, "修改手机号失败")
	ModifyUserInfoFailed = codex.New(20105, "修改信息失败")
)
