package errno

import (
	"errors"
	"fmt"
)

// 业务状态码
const (
	SuccessCode                = 0     // 请求成功
	ServerErrCode              = 10001 // 服务器错误
	ParamErrorCode             = 10002 // 参数错误
	AlreadyExistErrCode        = 10003 // 数据已存在
	AuthorizationFailedErrCode = 10004 // 身份验证失败
	NotFoundErrCode            = 10005 // 数据未找到
)

type ErrNo struct {
	ErrCode int
	ErrMsg  string
}

func (e ErrNo) Error() string {
	return fmt.Sprintf("err_code=%d, err_msg=%s", e.ErrCode, e.ErrMsg)
}

func (e ErrNo) WithMessage(msg string) ErrNo {
	e.ErrMsg = msg
	return e
}

func NewErrNo(code int, msg string) ErrNo {
	return ErrNo{
		ErrCode: code,
		ErrMsg:  msg,
	}
}

var (
	Success                = NewErrNo(SuccessCode, "success")
	ServerErr              = NewErrNo(ServerErrCode, "internal server error")
	ParamErr               = NewErrNo(ParamErrorCode, "params error")
	AlreadyExistErr        = NewErrNo(AlreadyExistErrCode, "data already exists")
	AuthorizationFailedErr = NewErrNo(AuthorizationFailedErrCode, "authorization failed")
	NotFoundErr            = NewErrNo(NotFoundErrCode, "not found")
)

// ConvertErr convert error to ErrNo
func ConvertErr(err error) ErrNo {
	Err := ErrNo{}

	// 判断 err 是否为 ErrNo，如果是，则会把错误信息传递给 Err 变量
	if errors.As(err, &Err) {
		return Err
	}

	s := ServerErr
	s.ErrMsg = err.Error()
	return s
}
