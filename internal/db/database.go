package db

import (
	"database/sql"
	"fmt"
	"sync"
	"time"

	"github.com/gnanasuriyan/go-message-server/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type IDatabaseConfig interface {
	GetDatabaseConfig() config.DBConfig
}

type IAppDB interface {
}

type AppDb struct {
	*sql.DB
}

var once sync.Once
var Db *AppDb

func InitDatabase() *AppDb {
	once.Do(func() {
		dbConfig := config.GetConfig().GetDatabaseConfig()
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Dbname)
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err)
		}
		sqlDB, err := db.DB()
		if err != nil {
			panic(err)
		}
		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(100)
		sqlDB.SetConnMaxLifetime(time.Hour)
		Db = &AppDb{sqlDB}
	})
	return Db
}
