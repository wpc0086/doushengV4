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

package pack

import (
	"doushengV4/kitex_gen/user"
	"doushengV4/pkg/errno"
	"errors"
)

// BuildRegisterResp build RegisterResp from error
func BuildRegisterResp(err error) *user.RegisterResp {
	if err == nil {
		return RegisterResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return RegisterResp(e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return RegisterResp(s)
}

func RegisterResp(err errno.ErrNo) *user.RegisterResp {
	return &user.RegisterResp{StatusCode: int32(err.ErrCode), StatusMsg: err.ErrMsg}
}

func BuildLoginResp(err error) *user.LoginResp {
	if err == nil {
		return LoginResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return LoginResp(e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return LoginResp(s)
}

func LoginResp(err errno.ErrNo) *user.LoginResp {
	return &user.LoginResp{StatusCode: int32(err.ErrCode), StatusMsg: err.ErrMsg}
}

func BuildInfoResp(err error) *user.InfoUserResponse {
	if err == nil {
		return InfoResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return InfoResp(e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return InfoResp(s)
}

func InfoResp(err errno.ErrNo) *user.InfoUserResponse {
	return &user.InfoUserResponse{StatusCode: int32(err.ErrCode), StatusMsg: err.ErrMsg}
}
