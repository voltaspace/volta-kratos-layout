package driver

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"sweet/internal/conf"
)

func NewMysql(c *conf.Data_Database) (db *gorm.DB){
	db, err := gorm.Open("mysql",c.Source)
	if err != nil {
		panic(err)
	}
	return
}
