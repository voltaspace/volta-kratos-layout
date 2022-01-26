package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"sweet/internal/biz"
)

type sweetRepo struct {
	data *Data
	log  *log.Helper
}

// NewSweetRepo .
func NewSweetRepo(data *Data, logger log.Logger) biz.SweetRepo {
	return &sweetRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *sweetRepo) Create(ctx context.Context, g *biz.Sweet) error {
	return nil
}

func (r *sweetRepo) Update(ctx context.Context, g *biz.Sweet) error {
	return nil
}

func (r *sweetRepo) GetRedis() map[string]*redis.Client{
	return r.data.Redis
}

func (r *sweetRepo) GetRedisCtx() context.Context{
	return r.data.RedisCtx
}


