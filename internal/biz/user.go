package biz

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"golang.org/x/crypto/bcrypt"
)

// 领域对象
type User struct {
	Email string
	Username string
	Bio string
	Image string
	PasswordHash string
}

type UserLogin struct {
	Email string
	Token string
	Username string
	Bio string
	Image string
}

func hashPassword(pwd string) string {
	b, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v", b)
	return string(b)
}

func verifyPassword(hashed, input string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(input)); err != nil{
		return false
	}
	return true
}

// 接口，定义一些方法
type UserRepo interface {
	CreateUser(ctx context.Context, user *User) error
	GetUserByEmail(ctx context.Context, email string) (*User, error)
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

func (uc *UserUsecase) Register(ctx context.Context, username, email, password string) (*UserLogin, error) {
	u := &User{
		Email: email,
		Username: username,
		PasswordHash: hashPassword(password),
	}
	if err := uc.ur.CreateUser(ctx, u); err != nil {
		return nil, err
	}
	return &UserLogin{
		Email: email,
		Username: username,
		Token: "abc",
	}, nil
}

func (uc *UserUsecase) Login(ctx context.Context, email, password string) (*UserLogin, error) {
	u, err := uc.ur.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if !verifyPassword(u.PasswordHash, password) {
		return nil, errors.New("login failed")
	}
	return &UserLogin{
		Email: u.Email,
		Username: u.Username,
		Bio: u.Bio,
		Image: u.Image,
		Token: "abc",
	}, nil
}
