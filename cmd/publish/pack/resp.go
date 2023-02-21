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
	"doushengV4/kitex_gen/publish"
	"doushengV4/pkg/errno"
	"errors"
)

// BuildActionResp build ActionResp from error
func BuildActionResp(err error) *publish.ActionResp {
	if err == nil {
		return ActionResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return ActionResp(e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return ActionResp(s)
}

func ActionResp(err errno.ErrNo) *publish.ActionResp {
	return &publish.ActionResp{StatusCode: int32(err.ErrCode), StatusMsg: err.ErrMsg}
}

func BuildListResp(err error) *publish.ListResp {
	if err == nil {
		return ListResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return ListResp(e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return ListResp(s)
}

func ListResp(err errno.ErrNo) *publish.ListResp {
	return &publish.ListResp{StatusCode: int32(err.ErrCode), StatusMsg: err.ErrMsg}
}

func BuildFeedResp(err error) *publish.FeedResponse {
	if err == nil {
		return FeedResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return FeedResp(e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return FeedResp(s)
}

func FeedResp(err errno.ErrNo) *publish.FeedResponse {
	return &publish.FeedResponse{StatusCode: int32(err.ErrCode), StatusMsg: err.ErrMsg}
}
