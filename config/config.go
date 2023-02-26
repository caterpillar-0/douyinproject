package config

import (
	"MyProject/dao"
	"MyProject/entity"
	"fmt"
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	SERVER_HOST string
	SERVER_PORT string
)

func LoadDB() *gorm.DB {
	SetupViper()
	dns := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		viper.GetString("mysql.name"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.host"),
		viper.GetString("mysql.port"),
		viper.GetString("mysql.database"),
	)
	log.Println(dns)
	//username:password@protocol(address)/dbname?param=value
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "t_", // 表名前缀，`User` 的表名应该是 `t_users`
			SingularTable: true, // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `t_user`
		},
	})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}

func SetupDB() {
	db := LoadDB()
	db.AutoMigrate(&entity.User{})
	//combine db with global val Q, otherwise Q = dao.User(db)
	dao.SetDefault(db)

}

func SetupViper() {
	//viper setting
	viper.SetConfigFile("conf.yaml")
	err := viper.ReadInConfig() //read conf.yaml
	if err != nil {
		panic("faild to load config file")
	}
}

func SetUp() {
	SetupDB()

	//get global variables
	SERVER_HOST = viper.GetString("server.host")
	SERVER_PORT = viper.GetString("server.port")

}
