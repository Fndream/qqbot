package service

import (
	"errors"
	"gorm.io/gorm"
	"math"
	"qqbot/handle/dto"
	"qqbot/handle/reward"
	"qqbot/handle/tb"
	"qqbot/handle/util"
	"qqbot/pkg/db"
	"time"
)

// Sign 签到
func Sign(uid string) (res *dto.Sign, err error) {
	var userSign tb.UserSign
	if e := db.DB.Model(&tb.UserSign{}).First(&userSign, uid).Error; e != nil && !errors.Is(e, gorm.ErrRecordNotFound) {
		return nil, e
	}
	var userInfo tb.UserInfo
	if e := db.DB.Model(&tb.UserInfo{}).First(&userInfo, uid).Error; e != nil && !errors.Is(e, gorm.ErrRecordNotFound) {
		return nil, e
	}

	// 检查是否已签到
	ny, nm, nd := time.Now().Date()
	sy, sm, sd := userSign.UpdatedAt.Date()
	if ny == sy && nm == sm && nd == sd {
		err = errors.New("📅 今天已经签到过了")
		return
	}

	rewardGold := util.RandInt(1000, 2000)
	var rewardDiamond int
	switch c := userSign.Continuous; {
	case c > 30:
		rewardDiamond = 3
	case c > 7:
		rewardDiamond = 2
	default:
		rewardDiamond = 1
	}

	var newUserSign tb.UserSign
	if err = db.DB.Transaction(func(tx *gorm.DB) error {
		if userSign.ID == "" {
			// insert userSign
			newUserSign = tb.UserSign{ID: uid, Count: 1, Continuous: 1}
			if e := tx.Create(&newUserSign).Error; e != nil {
				return e
			}
		} else {
			// update userSign
			newUserSign = tb.UserSign{Count: userSign.Count + 1, Continuous: userSign.Continuous + 1}
			yy, ym, yd := time.Now().AddDate(0, 0, -1).Date()
			uy, um, ud := userSign.UpdatedAt.Date()
			if yy != uy || ym != um || yd != ud {
				newUserSign.Continuous = 1
			}
			if e := tx.Model(&userSign).Updates(&newUserSign).Error; e != nil {
				return e
			}
		}

		if e := reward.NewUserRewardTx(uid, tx).AddInfo(&tb.UserInfo{
			Gold:    rewardGold,
			Diamond: rewardDiamond,
		}).Save(); e != nil {
			return e
		}
		return nil
	}); err != nil {
		return nil, err
	}

	res = &dto.Sign{
		Count:         int(newUserSign.Count),
		Continuous:    int(newUserSign.Continuous),
		RewardGold:    rewardGold,
		RewardDiamond: rewardDiamond,
	}
	return
}

// Query 查询
func Query(uid string) (res *dto.Query, err error) {
	var userSign tb.UserSign
	if e := db.DB.Model(&tb.UserSign{}).First(&userSign, uid).Error; e != nil && !errors.Is(e, gorm.ErrRecordNotFound) {
		return nil, e
	}
	var userInfo tb.UserInfo
	if e := db.DB.Model(&tb.UserInfo{}).First(&userInfo, uid).Error; e != nil && !errors.Is(e, gorm.ErrRecordNotFound) {
		return nil, e
	}
	res = &dto.Query{
		UserInfo: &userInfo,
		UserSign: &userSign,
	}
	return
}

// Bag 背包
func Bag(uid string, pageNo int) (res *dto.Bag, err error) {
	// 查询拥有物品种类数量
	var count int64
	if e := db.DB.Model(&tb.UserItem{}).Where(&tb.UserItem{ID: uid}).Count(&count).Error; e != nil {
		return nil, e
	}

	// 计算总页数
	pageSize := 10
	pageTotal := int(math.Ceil(float64(count) / float64(pageSize)))

	// 分页查询背包物品
	var userItems []*tb.UserItem
	if e := db.DB.Model(&tb.UserItem{}).Preload("Item").Where(&tb.UserItem{ID: uid}).Offset((pageNo - 1) * pageSize).Limit(pageSize).Order("item_id asc").Find(&userItems).Error; e != nil && !errors.Is(e, gorm.ErrRecordNotFound) {
		return nil, e
	}
	if len(userItems) == 0 {
		return nil, errors.New("🎒 背包空空如也")
	}

	res = &dto.Bag{
		PageNo:    pageNo,
		PageTotal: pageTotal,
		Items:     userItems,
	}

	return
}

func GetUserInfo(uid string) (res *tb.UserInfo, err error) {
	var userInfo tb.UserInfo
	if e := db.DB.Model(&tb.UserInfo{}).First(&userInfo, uid).Error; e != nil && !errors.Is(e, gorm.ErrRecordNotFound) {
		return nil, e
	}
	return &userInfo, nil
}

func GetUserItem(uid string, itemName string) (res *tb.UserItem, err error) {
	var userItem tb.UserItem
	if e := db.DB.Model(&tb.UserItem{}).Joins("Item").Where("user_items.id = ? AND item.name = ?", uid, itemName).First(&userItem).Error; e != nil && errors.Is(e, gorm.ErrRecordNotFound) {
		return nil, e
	}
	return &userItem, nil
}
