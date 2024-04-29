package databases

import (
	"be22/clean-arch/app/configs"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDBMysql(cfg *configs.AppConfig) *gorm.DB {
	// var err error
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	// dsn := "root:qwerty123@tcp(127.0.0.1:3306)/db_be22_unit2?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DB_USERNAME, cfg.DB_PASSWORD, cfg.DB_HOSTNAME, cfg.DB_PORT, cfg.DB_NAME)
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return DB
}
