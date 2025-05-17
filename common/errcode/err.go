package errcode

import (
	"encoding/json"
	"fmt"
	"net/http"
	"runtime"
)

type AppError struct {
	// 唯一标识
	code int
	// 直观描述
	msg      string
	cause    error
	occurred string
}

func newError(code int, msg string) *AppError {
	if _, duplicated := codes[code]; duplicated {
		panic(fmt.Sprintf("the code %s already existed", code))
	}
	return &AppError{
		code: code,
		msg:  msg,
	}
}

// Error 实现了 error 接口
// 直接序列化，更快
func (a *AppError) Error() string {
	// 避免访问空指针成员
	if a == nil {
		return ""
	}
	formatterErr := struct {
		Code     int    `json:"code"`
		Msg      string `json:"msg"`
		Cause    string `json:"cause"`
		Occurred string `json:"occurred"`
	}{
		Code:     a.code,
		Msg:      a.msg,
		Occurred: a.occurred,
	}
	if a.cause != nil {
		formatterErr.Cause = a.cause.Error()
	}
	bytes, _ := json.Marshal(formatterErr)
	return string(bytes)
}

func (a *AppError) String() string {
	return a.Error()
}

func (a *AppError) GetCode() int {
	return a.code
}

func (a *AppError) GetMsg() string {
	return a.msg
}

func (a *AppError) HttpStatusCode() int {
	switch a.code {
	case Success.GetCode():
		return http.StatusOK
	case ErrServer.GetCode():
		return http.StatusInternalServerError
	case ErrParams.GetCode():
		return http.StatusBadRequest
	case ErrNotFound.GetCode():
		return http.StatusNotFound
	case ErrTooManyRequests.GetCode():
		return http.StatusTooManyRequests
	case ErrToken.GetCode():
		return http.StatusUnauthorized
	case ErrForbidden.GetCode():
		return http.StatusForbidden
	default:
		return http.StatusInternalServerError
	}
}

// Wrap 函数主要用户包装错误信息，用于日志记录
// 包装底层任务和 WithCause 一样是为了记录错误链条
func Wrap(msg string, err error) *AppError {
	return &AppError{
		msg:      msg,
		cause:    err,
		code:     -1,
		occurred: getErrorInfo(),
	}
}

// WithCause 复用预定好的错误信息
// 使用于错误码定义地比较详细的项目
func (a *AppError) WithCause(err error) *AppError {
	a.cause = err
	a.occurred = getErrorInfo()
	return a
}

func getErrorInfo() string {
	pc, file, line, _ := runtime.Caller(2)
	funcName := runtime.FuncForPC(pc).Name()
	return fmt.Sprintf("funName: %s file: %s line: %d", funcName, file, line)
}
