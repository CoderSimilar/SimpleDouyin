package mysql

import (
	"SimpleDouyin/module"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Init mysql数据库初始化
func Init() (err error) {

	DB, err = connect2mysql()
	if err != nil {
		panic("failed to connect database")
	}

	// 自动迁移
	err = DB.AutoMigrate(&module.User{}, &module.Video{})
	if err != nil {
		panic("failed to migrate data table")
	}

	return err
}

func connect2mysql() (*gorm.DB, error) {

	dsn := "root:aa995231030@tcp(47.102.144.228:3306)/simpledouyin?charset=utf8&parseTime=True&loc=Local"
	// defer disConnect()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

// Close 关闭mysql数据库
func Close() {
	DB, _ := DB.DB()
	_ = DB.Close()
}
