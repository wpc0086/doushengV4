package service

import (
	"context"
	"doushengV4/cmd/interact/dal/db"
	"doushengV4/kitex_gen/interact"
	"doushengV4/pkg/errno"
	"doushengV4/pkg/mw"
)

type ActionCommentService struct {
	ctx context.Context
}

func NewActionCommentService(ctx context.Context) *ActionCommentService {
	return &ActionCommentService{ctx: ctx}
}

func (s *ActionCommentService) ActionComment(req *interact.CommentActionRequest) (*db.Comment, error) {
	claims, ok := mw.CheckToken(req.Token)
	if ok != true {
		return nil, errno.AuthorizationFailedErr
	}
	user_id := claims.UserId
	if req.ActionType == 1 {
		comment, err := db.SaveCommon(req.VideoId, user_id, req.CommentText)
		if err != nil {
			return nil, err
		}
		return comment, nil
	} else if req.ActionType == 2 {
		err := db.DelCommonById(req.CommentId, req.VideoId)
		if err != nil {
			return nil, err
		}
		return nil, nil
	} else {
		return nil, errno.ParamErr
	}
}
