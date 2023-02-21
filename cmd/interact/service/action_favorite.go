package service

import (
	"context"
	"doushengV4/cmd/interact/dal/db"
	"doushengV4/kitex_gen/interact"
	"doushengV4/pkg/errno"
	"doushengV4/pkg/mw"
)

type ActionFavoriteService struct {
	ctx context.Context
}

func NewActionFavoriteService(ctx context.Context) *ActionFavoriteService {
	return &ActionFavoriteService{ctx: ctx}
}

func (s *ActionFavoriteService) ActionFavorite(req *interact.FavoriteActionRequest) error {
	claims, ok := mw.CheckToken(req.Token)
	if ok != true {
		return errno.AuthorizationFailedErr
	}
	user_id := claims.UserId
	if req.ActionType == 1 {
		err := db.AddFavorite(req.VideoId, user_id, req.ActionType, 1)
		if err != nil {
			return err
		}
	} else if req.ActionType == 2 {
		err := db.AddFavorite(req.VideoId, user_id, req.ActionType, -1)
		if err != nil {
			return err
		}
	} else {
		return errno.ParamErr
	}

	return nil
}
