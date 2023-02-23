// Copyright 2022 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package errno

import (
	"doushengV4/kitex_gen/user"
	"errors"
	"fmt"
)

type ErrNo struct {
	ErrCode int64
	ErrMsg  string
}

func (e ErrNo) Error() string {
	return fmt.Sprintf("err_code=%d, err_msg=%s", e.ErrCode, e.ErrMsg)
}

func NewErrNo(code int64, msg string) ErrNo {
	return ErrNo{
		ErrCode: code,
		ErrMsg:  msg,
	}
}

func (e ErrNo) WithMessage(msg string) ErrNo {
	e.ErrMsg = msg
	return e
}

var (
	Success                = NewErrNo(int64(user.ErrCode_SuccessCode), "Success")
	ServiceErr             = NewErrNo(int64(user.ErrCode_ServiceErrCode), "Service is unable to start successfully")
	ParamErr               = NewErrNo(int64(user.ErrCode_ParamErrCode), "Wrong Parameter has been given")
	UserAlreadyExistErr    = NewErrNo(int64(user.ErrCode_UserAlreadyExistErrCode), "User already exists")
	AuthorizationFailedErr = NewErrNo(int64(user.ErrCode_AuthorizationFailedErrCode), "Authorization failed")
	ReapteOperationErr     = NewErrNo(int64(10005), "请勿重复操作")
	LoginErr               = NewErrNo(int64(10006), "用户名或密码不正确")
	EmailErr               = NewErrNo(int64(10007), "邮箱格式错误")
	DataErr                = NewErrNo(int64(10008), "数据找不到")
)

// ConvertErr convert error to Errno
func ConvertErr(err error) ErrNo {
	Err := ErrNo{}
	if errors.As(err, &Err) {
		return Err
	}
	s := ServiceErr
	s.ErrMsg = err.Error()
	return s
}
