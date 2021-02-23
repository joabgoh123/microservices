package repo

import (
	"context"
	"errors"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type UserRepo interface{
	CreateUser(ctx context.Context, user *user) error
	RetrieveUser(ctx context.Context, user *user) (user, error)
}

type Repo struct{
}

func (r *Repo) NewRepo() (*Repo, error) {
	dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	_, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return &Repo{}, nil
	}
	return nil, errors.New("Cannot connect to database")
}