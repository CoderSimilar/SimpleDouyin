package LarryTest

// // 链接数据库
// import (
// 	"fmt"

// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// )

// // 定义数据模型结构体
// type TestTable struct {
// 	ID             uint   `gorm:"primaryKey"`
// 	Name           string `gorm:"column:name"`
// 	UpdateDatetime string `gorm:"column:update_datetime"`
// }

// var db *gorm.DB // 将 db 声明为全局变量

// func init() {
// 	var err error
// 	db, err = connect2mysql()
// 	if err != nil {
// 		panic("failed to connect database")
// 	}

// 	err = db.AutoMigrate(&TestTable{})
// 	if err != nil {
// 		panic("failed to migrate data table")
// 	}
// }

// func connect2mysql() (*gorm.DB, error) {
// 	dsn := "root:aa995231030@tcp(47.102.144.228:3306)/simpledouyin?charset=utf8&parseTime=True&loc=Local"
// 	// defer disConnect()
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		return nil, err
// 	}
// 	return db, nil
// }

// func disConnect() {
// 	sqlDB, _ := db.DB()
// 	sqlDB.Close()
// }

// 	// func GetRowsByTable(tableName string, limits ...int) ([]*interface{}, error) {
// 	// 	var rows []*interface{}
// 	// 	table := db.Table(tableName)
// 	// 	if len(limits) == 0 {
// 	// 		limits = append(limits, 1000)
// 	// 	}
// 	// 	fmt.Println("Querying table:", tableName, "with limit:", limits[0])
// 	// 	if err := table.Limit(limits[0]).Find(&rows).Error; err != nil {
// 	// 		fmt.Println("Error querying table:", err)
// 	// 		return nil, err
// 	// 	}
// 	// 	return rows, nil
// 	// }

// // 从指定的表中获取行
// func GetRowsByTable(tableName string, modelSlice interface{}, limits ...int) error {
// 	table := db.Table(tableName)
// 	if len(limits) == 0 {
// 		limits = append(limits, 1000)
// 	}
// 	fmt.Println("Querying table:", tableName, "with limit:", limits[0])
// 	if err := table.Limit(limits[0]).Find(modelSlice).Error; err != nil {
// 		fmt.Println("Error querying table:", err)
// 		return err
// 	}
// 	return nil
// }
