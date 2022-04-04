package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"kratos-realworld/internal/biz"
)

// 小写，具体实现
type userRepo struct {
	data *Data
	log  *log.Helper
}

type User struct {
	gorm.Model
	Email string `gorm:"size:500"`
	Username string `gorm:"size:500"`
	Bio string `gorm:"size:1000"`
	Image string `gorm:"size:1000"`
	PasswordHash string `gorm:"size:500"`
}

// NewGreeterRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *userRepo) CreateUser(ctx context.Context, u *biz.User) error {
	user := User{
		Email: u.Email,
		Username: u.Username,
		Bio: u.Bio,
		Image: u.Image,
		PasswordHash: u.PasswordHash,
	}
	//spew.Dump(user)
	rv := r.data.db.Create(&user)
	return rv.Error
}

func (uc *userRepo) GetUserByEmail(ctx context.Context, email string) (rv *biz.User,err error) {
	u := new(User)
	result := uc.data.db.Where("email = ?", email).First(u)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.NotFound("user", "not found by email")
	}

	if result.Error != nil {
		return nil, err
	}
	return &biz.User{
		Email: u.Email,
		Username: u.Username,
		Bio:	u.Bio,
		Image: u.Image,
		PasswordHash: u.PasswordHash,
	}, nil
}


type profileRepo struct {
	data *Data
	log  *log.Helper
}

func NewProfileRepo(data *Data, logger log.Logger) biz.ProfileRepo {
	return &profileRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}