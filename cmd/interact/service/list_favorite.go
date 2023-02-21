package service

import (
	"context"
	"doushengV4/cmd/interact/dal/db"
	"doushengV4/cmd/interact/pack"
	"doushengV4/cmd/interact/rpc"
	"doushengV4/kitex_gen/interact"
	"doushengV4/kitex_gen/user"
)

type ListFavoriteService struct {
	ctx context.Context
}

func NewListFavoriteService(ctx context.Context) *ListFavoriteService {
	return &ListFavoriteService{ctx: ctx}
}

func (s *ListFavoriteService) ListFavorite(req *interact.FavoriteListRequest) ([]*interact.Video, error) {
	videos, err := db.GetVideosByUserId(req.UserId)
	if err != nil {
		return nil, err
	}
	publishes := pack.Publishs(videos)
	for i, video := range videos {
		infoUser, err := rpc.InfoUser(s.ctx, &user.InfoUserRequest{UserId: video.AuthorId})
		if err != nil {
			continue
		}
		publishes[i].Author = pack.User(infoUser.User)
	}
	return publishes, nil
}
