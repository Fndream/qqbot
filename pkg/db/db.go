package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"qqbot/handle/tb"
)

var DB *gorm.DB

func SetUp() {
	var err error

	dsn := "root:root@tcp(127.0.0.1:3306)/guild_bot?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction:                   false,
		DisableForeignKeyConstraintWhenMigrating: true,
		//Logger:                                   logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}

	autoMigrate()

	sqlDB, err := DB.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.SetMaxIdleConns(16)
	sqlDB.SetMaxOpenConns(64)
}

func autoMigrate() {
	if err := DB.AutoMigrate(
		&tb.UserInfo{}, // 用户数据
		&tb.UserSign{}, // 用户签到
		&tb.UserItem{}, // 用户背包
		&tb.UserPet{},  // 用户宠物

		&tb.Good{}, // 商店
		&tb.Item{}, // 物品

		&tb.Pet{},       //宠物
		&tb.Feature{},   // 属性
		&tb.Character{}, // 性格
	); err != nil {
		panic(err)
	}
}
