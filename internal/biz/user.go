package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

// 领域对象
type User struct {
	Username string
}

// 接口，定义一些方法
type UserRepo interface {
	CreateUser(ctx context.Context, user *User) error
}

type ProfileRepo interface {

}

type UserUsecase struct {
	ur UserRepo
	pr ProfileRepo
	log  *log.Helper
}

func NewUserUsecase(ur UserRepo, pr ProfileRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{
		ur:ur, pr: pr, log: log.NewHelper(logger),
	}
}

func (uc *UserUsecase) Register(ctx context.Context, u * User) error {
	if err := uc.ur.CreateUser(ctx, u); err != nil {

	}
	return nil
}

