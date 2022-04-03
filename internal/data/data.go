package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"kratos-realworld/internal/conf"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewDB, NewData, NewUserRepo, NewProfileRepo)

// Data .
type Data struct {
	db *gorm.DB
}

// NewData .
func NewData(c *conf.Data, logger log.Logger, db *gorm.DB) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{db: db}, cleanup, nil
}

func NewDB(c *conf.Data) *gorm.DB {
	db, err := gorm.Open(mysql.Open(c.Database.Dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true, // 不加外键
	})
	if err != nil{
		panic("failed to connect database")
	}
	InitDB(db)
	return db
}

func InitDB(db *gorm.DB){
	// 把数据库表结构创建好，不建议生产环境线上用AutoMigrate
	if err := db.AutoMigrate(&User{}); err != nil {
		panic(err)
	}
}