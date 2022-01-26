package driver

import (
	"context"
	"github.com/olivere/elastic/v7"
	"strconv"
	"sweet/internal/conf"
)

func NewElasticSearch(c *conf.Data) (conn map[string]*elastic.Client){
	conn = make(map[string]*elastic.Client)
	for  _,cf := range c.Elasticsearch {
		if cf.Power == false {
			continue
		}
		client, err := elastic.NewClient(
			elastic.SetSniff(false),
			elastic.SetURL(cf.Host),
			elastic.SetBasicAuth(cf.Username, cf.Password),
		)
		if err != nil {
			panic(err.Error())
		}
		_, code, err := client.Ping(cf.Host).Do(context.Background())
		if err != nil {
			panic("elastic ping err " + strconv.FormatInt(int64(code),10))
		}
		_, err = client.ElasticsearchVersion(cf.Host)
		if err != nil {
			panic("elastic search version err")
		}
		conn[cf.Name] = client
	}
	return
}
