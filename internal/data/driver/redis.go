package driver

import (
	"context"
	"github.com/go-redis/redis/v8"
	"sweet/internal/conf"
)

func NewRedis(c *conf.Data,ctx context.Context) (conn map[string]*redis.Client){
	conn = make(map[string]*redis.Client)
	for _,cf := range c.Redis {
		if cf.Power == false {
			continue
		}
		rdb := redis.NewClient(&redis.Options{
			Addr:     cf.Addr,
			Password: cf.Password,
			WriteTimeout: cf.WriteTimeout.AsDuration(),
			ReadTimeout: cf.ReadTimeout.AsDuration(),
			DB:       int(cf.Db),
		})
		_, err := rdb.Ping(ctx).Result()
		if err != nil {
			panic("连接redis出错，错误信息:"+err.Error())
		}
		conn[cf.Name] = rdb
	}
	return conn
}

func NewRedisCtx() context.Context{
	return context.Background()
}