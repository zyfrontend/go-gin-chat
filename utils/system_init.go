package utils

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var DB *gorm.DB

func InitConfig() {
	viper.SetConfigName("app")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
}
func InitMysql() {
	// 自定义日志模板 打印SQL语句
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, // 慢阈值
			LogLevel:      logger.Info, // 级别
			Colorful:      true,        // 彩色
		},
	)
	db, err := gorm.Open(mysql.Open(viper.GetString("mysql.dns")), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		fmt.Println("err: ", err)
	}
	DB = db
	fmt.Println("config mysql inited ....")
}
