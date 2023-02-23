package service

import (
	"context"
	"doushengV4/cmd/interact/dal/db"
	"doushengV4/cmd/interact/dal/redis"
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
	userId := claims.UserId
	//延时双删
	redis.RemoveFavoriteList(s.ctx, userId)
	if req.ActionType == 1 || req.ActionType == 2 {
		err := db.ModifyFavorite(req.VideoId, userId, 1)
		if err != nil {
			return err
		}
		redis.ModifyFavorite(s.ctx, userId, 1)
		redis.RemoveFavoriteList(s.ctx, userId)
		return nil
	} else {
		return errno.ParamErr
	}

}
