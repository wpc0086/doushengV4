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
	"doushengV4/kitex_gen/interact"
	"doushengV4/pkg/errno"
	"errors"
)

func BuildFavoriteActionResp(err error) *interact.FavoriteActionResponse {
	if err == nil {
		return FavoriteActionResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return FavoriteActionResp(e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return FavoriteActionResp(s)
}

func FavoriteActionResp(err errno.ErrNo) *interact.FavoriteActionResponse {
	return &interact.FavoriteActionResponse{StatusCode: int32(err.ErrCode), StatusMsg: err.ErrMsg}
}

func BuildFavoriteListResp(err error) *interact.FavoriteListResponse {
	if err == nil {
		return FavoriteListResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return FavoriteListResp(e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return FavoriteListResp(s)
}

func FavoriteListResp(err errno.ErrNo) *interact.FavoriteListResponse {
	return &interact.FavoriteListResponse{StatusCode: int32(err.ErrCode), StatusMsg: err.ErrMsg}
}

func BuildCommentActionResp(err error) *interact.CommentActionResponse {
	if err == nil {
		return CommentActionResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return CommentActionResp(e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return CommentActionResp(s)
}

func CommentActionResp(err errno.ErrNo) *interact.CommentActionResponse {
	return &interact.CommentActionResponse{StatusCode: int32(err.ErrCode), StatusMsg: err.ErrMsg}
}

func BuildCommentListResp(err error) *interact.CommentListResponse {
	if err == nil {
		return CommentListResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return CommentListResp(e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return CommentListResp(s)
}

func CommentListResp(err errno.ErrNo) *interact.CommentListResponse {
	return &interact.CommentListResponse{StatusCode: int32(err.ErrCode), StatusMsg: err.ErrMsg}
}
