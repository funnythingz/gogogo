package db

import (
	"fmt"
	"github.com/BurntSushi/toml"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/k0kubun/pp"
	"log"
)

var Dbmap gorm.DB

func DbOpen(
	adapter string,
	host string,
	username string,
	password string,
	database string,
	encoding string,
) {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=True&collation=utf8mb4_unicode_ci", username, password, host, database, encoding)
	Dbmap, _ = gorm.Open(adapter, dataSourceName)

	Dbmap.DB()

	Dbmap.DB().Ping()
	Dbmap.DB().SetMaxIdleConns(10)
	Dbmap.DB().SetMaxOpenConns(100)

	Dbmap.SingularTable(true)
}

type env struct {
	Host     string `toml:"host"`
	Username string `toml:"username"`
	Password string `toml:"password"`
	Database string `toml:"database"`
}

type connection struct {
	Adapter  string `toml:"adapter"`
	Encoding string `toml:"encoding"`
	Host     string `toml:"host"`
	Database string `toml:"database"`
}

type Config struct {
	Connection connection
	Databases  map[string]env
}

func DbConfig() Config {
	var config Config
	if _, err := toml.DecodeFile("db/database.toml", &config); err != nil {
		log.Println(err)
	}
	return config
}

func DbConnect(env string) {
	config := DbConfig()

	var (
		dbEnv = config.Databases[env]

		host = func() string {
			if dbEnv.Host != "" {
				return dbEnv.Host
			}
			return config.Connection.Host
		}()

		username = func() string {
			return dbEnv.Username
		}()

		password = func() string {
			return dbEnv.Password
		}()

		database = func() string {
			if dbEnv.Database != "" {
				return dbEnv.Database
			}
			return config.Connection.Database
		}()
	)

	DbOpen(
		config.Connection.Adapter,
		host,
		username,
		password,
		database,
		config.Connection.Encoding,
	)
}
