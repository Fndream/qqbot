package service

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"math"
	"qqbot/handle/dto"
	"qqbot/handle/reward"
	"qqbot/handle/tb"
	"qqbot/pkg/db"
)

// Shop 商城
func Shop(pageNo int) (res *dto.Shop, err error) {
	// 查询商品数量
	var count int64
	if e := db.DB.Model(&tb.Good{}).Where("buy_gold > 0 or buy_diamond > 0").Count(&count).Error; e != nil {
		return nil, e
	}

	// 计算总页数
	pageSize := 6
	pageTotal := int(math.Ceil(float64(count) / float64(pageSize)))

	// 分页查询商品
	var goods []*tb.Good
	if e := db.DB.Model(&tb.Good{}).Preload("UserItem").Where("buy_gold > 0 or buy_diamond > 0").Offset((pageNo - 1) * pageSize).Limit(pageSize).Order("item_id asc").Find(&goods).Error; e != nil && !errors.Is(e, gorm.ErrRecordNotFound) {
		return nil, e
	}
	if len(goods) == 0 {
		return nil, errors.New("🛍 商城空空如也")
	}

	res = &dto.Shop{
		PageNo:    pageNo,
		PageTotal: pageTotal,
		Goods:     goods,
	}

	return
}

// Trading 交易所
func Trading(pageNo int) (res *dto.Shop, err error) {
	// 查询商品数量
	var count int64
	if e := db.DB.Model(&tb.Good{}).Where("sell_gold > 0 or sell_diamond > 0").Count(&count).Error; e != nil {
		return nil, e
	}

	// 计算总页数
	pageSize := 6
	pageTotal := int(math.Ceil(float64(count) / float64(pageSize)))

	// 分页查询商品
	var goods []*tb.Good
	if e := db.DB.Model(&tb.Good{}).Preload("UserItem").Where("sell_gold > 0 or sell_diamond > 0").Offset((pageNo - 1) * pageSize).Limit(pageSize).Order("item_id asc").Find(&goods).Error; e != nil && !errors.Is(e, gorm.ErrRecordNotFound) {
		return nil, e
	}
	if len(goods) == 0 {
		return nil, errors.New("🛍 商城空空如也")
	}

	res = &dto.Shop{
		PageNo:    pageNo,
		PageTotal: pageTotal,
		Goods:     goods,
	}

	return
}

func BuyGood(uid string, buyName string, buyCount int) (res *dto.BuyGood, err error) {
	// 查询要购买的商品信息
	var good tb.Good
	if e := db.DB.Model(&tb.Good{}).Joins("UserItem").Where("item.name = ?", buyName).First(&good).Error; e != nil {
		if errors.Is(e, gorm.ErrRecordNotFound) {
			return nil, errors.New(fmt.Sprintf("🛍 不存在名为%v的商品", buyName))
		}
		return nil, e
	}
	// 判断商品是否可购买
	if good.BuyGold == 0 && good.BuyDiamond == 0 {
		return nil, errors.New(fmt.Sprintf("%v %v暂时不可购买", good.Item.Emoji, buyName))
	}

	// 判断用户是否拥有足够的金币或钻石
	var user tb.UserInfo
	if e := db.DB.Model(&tb.UserInfo{}).First(&user, uid).Error; e != nil && !errors.Is(e, gorm.ErrRecordNotFound) {
		return nil, e
	}

	subGold := int(good.BuyGold) * buyCount
	subDiamond := int(good.BuyDiamond) * buyCount

	var emsg string
	if good.BuyGold > 0 && user.Gold < subGold {
		emsg += fmt.Sprintf("💰 金币不足\n💰 一共需要%v枚金币", int(good.BuyGold)*buyCount)
	}
	if good.BuyDiamond > 0 && user.Diamond < subDiamond {
		emsg += fmt.Sprintf("💎 钻石不足\n💎 一共需要%v颗钻石", int(good.BuyDiamond)*buyCount)
	}
	if emsg != "" {
		return nil, errors.New(emsg)
	}

	// 判断用户的背包是否达到购买上限
	var userItem tb.UserItem
	if e := db.DB.Model(&tb.UserItem{}).Preload("UserItem").Where("id = ? AND item_id = ?", uid, good.ItemID).First(&userItem).Error; e != nil && !errors.Is(e, gorm.ErrRecordNotFound) {
		return nil, e
	}
	if userItem.Count+buyCount > 9999 {
		return nil, errors.New(fmt.Sprintf("%v 背包无法容纳这么多的%v", good.Item.Emoji, good.Item.Name))
	}

	// 扣除金币钻石，发放物品
	reward := reward.NewUserReward(uid)
	reward.AddInfo(&tb.UserInfo{
		Gold:    -subGold,
		Diamond: -subDiamond,
	})
	reward.AddItem(good.ItemID, buyCount)
	if err = reward.Save(); err != nil {
		return
	}

	res = &dto.BuyGood{
		Good:       &good,
		BuyCount:   buyCount,
		SubGold:    subGold,
		SubDiamond: subDiamond,
	}
	return
}
