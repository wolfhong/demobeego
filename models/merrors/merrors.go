package merrors

import (
	"errors"
)

var (
	NoPermission  = errors.New("no permission")
	EmailExist    = errors.New("邮箱已经存在")
	AccountExist  = errors.New("账号已经存在")
	TokenNotExist = errors.New("token not exist")
	TokenExpired  = errors.New("token already expired")
)
