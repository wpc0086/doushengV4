package main

import (
	"context"
	"doushengV4/cmd/user/dal/redis"
	"doushengV4/cmd/user/pack"
	"doushengV4/cmd/user/service"
	user "doushengV4/kitex_gen/user"
	"doushengV4/pkg/errno"
	"doushengV4/pkg/mw"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// RegisterUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) RegisterUser(ctx context.Context, req *user.RegisterUserRequest) (resp *user.RegisterResp, err error) {
	password := req.Password
	username := req.Username
	// 校验参数
	if len(username) > 32 {
		resp = pack.BuildRegisterResp(errno.ParamErr)
		return resp, nil
	}
	if len(password) > 32 {
		resp = pack.BuildRegisterResp(errno.ParamErr)
		return resp, nil
	}
	//注册用户
	uid, token, err := service.NewRegisterUserService(ctx).RegisterUser(req)
	if err != nil {
		resp = pack.BuildRegisterResp(err)
		return resp, nil
	}
	resp = &user.RegisterResp{StatusCode: int32(user.ErrCode_SuccessCode), UserId: uid, Token: token}
	return resp, nil
}

// LoginUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) LoginUser(ctx context.Context, req *user.LoginUserRequest) (resp *user.LoginResp, err error) {
	if len(req.Password) < 5 {
		resp = pack.BuildLoginResp(errno.ParamErr)
		return resp, nil
	}
	pwdMD5 := mw.Md5Encrypt(req.Password)
	req.Password = pwdMD5
	user_id, token, err := service.NewLoginUserService(ctx).LoginUser(req)
	if err != nil {
		resp = pack.BuildLoginResp(err)
		return resp, nil
	}
	resp = &user.LoginResp{StatusCode: int32(user.ErrCode_SuccessCode), UserId: user_id, Token: token}
	return resp, nil
}

// InforUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) InforUser(ctx context.Context, req *user.InfoUserRequest) (resp *user.InfoUserResponse, err error) {
	if redis.GetErr(ctx, req.Token) == true { //防止缓存穿透
		return nil, errno.DataErr
	}
	infoUser, err := service.NewInfoUserService(ctx).InfoUser(req)
	if err != nil {
		resp = pack.BuildInfoResp(err)
		redis.SaveErr(ctx, req.Token) //防止缓存穿透
		return resp, nil
	}
	resp = &user.InfoUserResponse{StatusCode: int32(user.ErrCode_SuccessCode), User: infoUser}
	return resp, nil
}
