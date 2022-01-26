package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
)

type Sweet struct {
	Hello string
}

type SweetRepo interface {
	Create(context.Context, *Sweet) error
	Update(context.Context, *Sweet) error
	GetRedis() map[string]*redis.Client
	GetRedisCtx() context.Context
}

type SweetUsecase struct {
	repo SweetRepo
	log  *log.Helper
}

func NewSweetUsecase(repo SweetRepo, logger log.Logger) *SweetUsecase {
	return &SweetUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *SweetUsecase) Create(ctx context.Context, g *Sweet) error {
	return uc.repo.Create(ctx, g)
}

func (uc *SweetUsecase) Update(ctx context.Context, g *Sweet) error {
	return uc.repo.Update(ctx, g)
}
