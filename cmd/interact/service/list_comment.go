package service

import (
	"context"
	"doushengV4/cmd/interact/dal/db"
	"doushengV4/cmd/interact/pack"
	"doushengV4/cmd/interact/rpc"
	"doushengV4/kitex_gen/interact"
	"doushengV4/kitex_gen/user"
)

type ListCommentService struct {
	ctx context.Context
}

func NewListCommentService(ctx context.Context) *ListCommentService {
	return &ListCommentService{ctx: ctx}
}

func (s *ListCommentService) ListComment(req *interact.CommentListRequest) ([]*interact.Comment, error) {
	cs, err := db.GetCommonsByVideoID(req.VideoId)
	if err != nil {
		return nil, err
	}
	comments := pack.Comments(cs)
	for i, c := range cs {
		infoUser, err := rpc.InfoUser(s.ctx, &user.InfoUserRequest{UserId: c.UserId})
		if err != nil {
			continue
		}
		comments[i].User = pack.User(infoUser.User)
	}
	return comments, nil
}
