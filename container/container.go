package container

import (
	"fmt"
	"log"

	//Mysql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/hysem/charts/config"
	"github.com/jmoiron/sqlx"
)

var (
	dbWriter *sqlx.DB
	dbReader *sqlx.DB
)

//Init loads container services
func Init() {
	conf := config.Current
	dbReader = initDB(&conf.Mysql)
	dbWriter = initDB(&conf.Mysql)
}

func initDB(mysqlConfig *config.MysqlConfig) *sqlx.DB {
	db, err := sqlx.Connect("mysql", fmt.Sprintf("%s:%s@(%s:%d)/%s",
		mysqlConfig.Username,
		mysqlConfig.Password,
		mysqlConfig.Host,
		mysqlConfig.Port,
		mysqlConfig.Database,
	))
	if err != nil {
		log.Fatalln("Error while connecting to database", err)
	}
	return db
}

//DBWriter returns db connection for writing to database
func DBWriter() *sqlx.DB {
	return dbWriter
}

//DBReader returns db connection for reading from database
func DBReader() *sqlx.DB {
	return dbReader
}
