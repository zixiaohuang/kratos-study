package service

import (
	"context"
	v1 "kratos-realworld/api/realworld/v1"
)

func (s *RealWorldService) Login(ctx context.Context, req *v1.LoginRequest) (reply *v1.UserReply, err error){

	userlogin, err := s.uc.Login(ctx, req.User.Email, req.User.Password)
	if err != nil {
		return nil, err
	}
	return &v1.UserReply{
		User: &v1.UserReply_User{
			Email: userlogin.Email,
			Username: userlogin.Username,
			Token: userlogin.Token,
		},
	}, nil
}

func (s *RealWorldService)Register(ctx context.Context, req *v1.RegisterRequest) (reply *v1.UserReply, err error){
	u, err := s.uc.Register(ctx, req.User.Username, req.User.Email, req.User.Password)
	if err != nil {
		return nil, err
	}
	return &v1.UserReply{
		User: &v1.UserReply_User{
			Email: u.Email,
			Username: u.Username,
			Token: u.Token,
		},
	}, nil
}

func (s *RealWorldService)GetCurrentUser(ctx context.Context, req *v1.GetCurrentUserRequest) (reply *v1.UserReply, err error) {
	return &v1.UserReply{
	}, nil
}

func (s *RealWorldService)UpdateUser(ctx context.Context, req *v1.UpdateUserRequest) (reply *v1.UserReply, err error) {
	return &v1.UserReply{
	}, nil
}

func (s *RealWorldService)GetProfile(ctx context.Context, req *v1.GetProfileRequest) (reply *v1.ProfileReply, err error) {
	return &v1.ProfileReply{
	}, nil
}

