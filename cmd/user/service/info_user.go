package service

import (
	"context"
	"doushengV4/cmd/user/dal/db"
	"doushengV4/cmd/user/pack"
	"doushengV4/kitex_gen/user"
)

type InfoUserService struct {
	ctx context.Context
}

func NewInfoUserService(ctx context.Context) *InfoUserService {
	return &InfoUserService{
		ctx: ctx,
	}
}

func (s *InfoUserService) InfoUser(req *user.InfoUserRequest) (user *user.User, err error) {
	u, err := db.GetUserById(req.UserId)
	if err != nil {
		return nil, err
	}
	u.Password = ""
	infoUser := pack.User(u)
	//查询视频表获取作品数
	videos, err := db.GetVideoListByUserId(infoUser.Id)
	if err != nil {
		return infoUser, err
	}
	infoUser.WorkCount = int64(len(videos))
	//获赞数量
	infoUser.TotalFavorited = 0
	for _, video := range videos {
		infoUser.TotalFavorited += int64(video.FavoriteCount)
	}
	//查询喜欢表获取喜欢数
	favoiteCount, _ := db.GetFavoritesByUserId(infoUser.Id)
	infoUser.FavoriteCount = favoiteCount
	return infoUser, nil
}
