package code

import "im/pkg/codex"

var (
	// param
	EmailEmpty    = codex.New(10001, "邮箱为空")
	PasswordEmpty = codex.New(10002, "密码为空")
	NicknameEmpty = codex.New(10003, "昵称为空")
	IdEmpty       = codex.New(10004, "ID为空")
	MobileEmpty   = codex.New(10005, "mobile为空")

	// err
	GenerateSaltFailed         = codex.New(10021, "密码盐值生成失败")
	GeneratePasswordHashFailed = codex.New(10022, "密码hash生成失败")
	RegisterFailed             = codex.New(10023, "注册失败")

	NotFoundUser         = codex.New(10024, "用户不存在")
	UpdateUserInfoFailed = codex.New(10025, "更新用户信息失败")
	EmailExist           = codex.New(10026, "邮箱已存在")
	MobileExist          = codex.New(10027, "Mobile已存在")
)
