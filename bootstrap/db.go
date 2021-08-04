package bootstrap

import (
	"buhang/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

var DB *gorm.DB

func Init() {
	var _ error

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true&loc=Local",
		config.Viper.GetString("DB_USERNAME"),
		config.Viper.GetString("DB_PASSWORD"),
		config.Viper.GetString("DB_HOST"),
		config.Viper.GetString("DB_PORT"),
		config.Viper.GetString("DB_DATABASE"),
	)

	gormConfig := mysql.New(mysql.Config{
		DSN: dsn,
	})

	var level gormlogger.LogLevel
	DB, _ = gorm.Open(gormConfig, &gorm.Config{
		Logger: gormlogger.Default.LogMode(level),
	})
}
