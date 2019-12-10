package datasource

import (
	"database/sql"
	"fmt"
	conf2 "github.com/birjemin/gin-structure/conf"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/cast"
	"time"
)

var db *gorm.DB

func GetDB() *gorm.DB {
	return db
}

// redis状态
func StatsDB() sql.DBStats {
	return db.DB().Stats()
}

// 关闭db
func CloseDb() error {
	return db.DB().Close()
}

func init() {
	path := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=true",
		conf2.String("mysql.user"),
		conf2.String("mysql.pass"),
		conf2.String("mysql.host"),
		conf2.Int("mysql.port"),
		conf2.String("mysql.db"),
	)

	var err error
	db, err = gorm.Open("mysql", path)
	if err != nil {
		panic(err)
	}
	db.SingularTable(true)
	db.DB().SetConnMaxLifetime(cast.ToDuration(conf2.Int("mysql.idletime"))  * time.Second)
	db.DB().SetMaxIdleConns(conf2.Int("mysql.maxidle"))   // 设置最大闲置个数
	db.DB().SetMaxOpenConns(conf2.Int("mysql.maxactive")) // 最大打开的连接数
	db.SingularTable(true)                                // 表生成结尾不带s
	// 是否启用Logger，显示详细日志
	db.LogMode(conf2.Bool("mysql.debug"))
}
