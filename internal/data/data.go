package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"github.com/olivere/elastic/v7"
	"sweet/internal/conf"
	"sweet/internal/data/driver"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewSweetRepo)

// Data .
type Data struct {
	DB map[string]*gorm.DB
	Redis map[string]*redis.Client
	RedisCtx context.Context
	Es map[string]*elastic.Client
	// TODO wrapped database client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	var data Data
	dbs := make(map[string]*gorm.DB)
	for _,dbConf := range c.Database{
		if dbConf.Driver == "mysql" && dbConf.Power == true{
			db := driver.NewMysql(dbConf)
			dbs[dbConf.Name] = db
		}
	}
	data.DB = dbs
	data.RedisCtx = driver.NewRedisCtx();
	data.Redis = driver.NewRedis(c,data.RedisCtx)
	data.Es = driver.NewElasticSearch(c)
	return &data, cleanup, nil
}

