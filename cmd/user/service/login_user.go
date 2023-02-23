package service

import (
	"context"
	"doushengV4/cmd/user/dal/db"
	"doushengV4/kitex_gen/user"
	"doushengV4/pkg/errno"
	"doushengV4/pkg/mw"
)

type LoginUserService struct {
	ctx context.Context
}

func NewLoginUserService(ctx context.Context) *LoginUserService {
	return &LoginUserService{
		ctx: ctx,
	}
}

func (s *LoginUserService) LoginUser(req *user.LoginUserRequest) (user_id int64, token string, err error) {
	u, err := db.CheckUser(req.Username, req.Password)
	if err != nil {
		return -1, "", errno.ServiceErr
	}
	if u.Id == 0 {
		return -1, "", errno.LoginErr
	}
	//生成token
	token, err = mw.CreateToken(u.Id, u.Name)
	return u.Id, token, nil
}
