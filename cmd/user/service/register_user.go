package service

import (
	"context"
	"doushengV4/cmd/user/dal/db"
	"doushengV4/kitex_gen/user"
	"doushengV4/pkg/errno"
	"doushengV4/pkg/mw"
)

type RegisterUserService struct {
	ctx context.Context
}

// NewCheckUserService new CheckUserService
func NewRegisterUserService(ctx context.Context) *RegisterUserService {
	return &RegisterUserService{
		ctx: ctx,
	}
}

func (s *RegisterUserService) RegisterUser(req *user.RegisterUserRequest) (int64, string, error) {
	// 查询用户是否已存在
	users, err := db.GetUserByName(req.Username)

	if err != nil {
		return -1, "", err
	}
	if len(users) != 0 {
		return -1, "", errno.UserAlreadyExistErr
	}

	// 创建用户
	newUser := db.User{
		Name:     req.Username,
		Password: mw.Md5Encrypt(req.Password),
	}

	if err := db.CreateUser(&newUser); err != nil {
		return -1, "", err
	}

	// 生成访问令牌
	token, err := mw.CreateToken(newUser.Id, req.Username)
	if err != nil {
		return -1, "", nil
	}

	return newUser.Id, token, nil
}
