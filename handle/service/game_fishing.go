package service

import (
	"errors"
	"qqbot/cmd/cache"
	"qqbot/handle/dto"
	"qqbot/handle/reward"
	"qqbot/handle/tb"
	"qqbot/handle/util"
	"qqbot/pkg/db"
	"time"
)

var fishCache = cache.New(time.Minute, time.Minute)

// Fishing 钓鱼
func Fishing(uid string) (res *dto.Fishing, err error) {
	// 获取玩家钓鱼数据
	v, ok := fishCache.Get(uid)

	// 如果正在钓鱼(数据存在且没过咬钩时间)
	if ok && time.Now().Before(v.(time.Time)) {
		err = errors.New("🎣 正在钓鱼中，在合适的时机进行【拉竿】吧")
		return
	}

	// 刷新钓鱼
	biteSecond := util.RandInt(5, 17)
	escapeTime := time.Now().Add(time.Duration(biteSecond+2) * time.Second)
	fishCache.Set(uid, escapeTime)
	res = &dto.Fishing{
		BiteSecond: biteSecond,
		MaxSecond:  biteSecond + 2,
	}
	return
}

// Draw 拉竿
func Draw(uid string) (res *dto.Draw, err error) {
	// 获取玩家钓鱼数据
	v, ok := fishCache.Get(uid)

	// 如果没在钓鱼(数据不存在)
	if !ok {
		err = errors.New(util.RandString([]string{
			"🎣 还没开始【钓鱼】呢，你拉啥子竿",
			"🎣 没在钓鱼中，先去【钓鱼】吧",
			"🎣 还没开始钓鱼呢，先去【钓鱼】吧",
			"🎣 还没开始【钓鱼】呢",
		}))
		return
	}

	escapeTime := v.(time.Time)
	biteTime := escapeTime.Add(-(time.Second * 2))

	// 如果未到钓鱼时间
	if time.Now().Before(biteTime) {
		fishCache.Delete(uid)
		err = errors.New(util.RandString([]string{
			"🎣 拉早了，鱼还没上钩呢",
			"🎣 笨！拉早了，鱼还没上钩呢！",
			"🎣 拉早了，鱼还没上钩呢，再来一次吧",
		}))
		return
	}

	// 如果已过脱钩时间
	if time.Now().After(escapeTime) {
		fishCache.Delete(uid)
		err = errors.New(util.RandString([]string{
			"🎣 拉晚了，鱼跑了啦！",
			"🎣 笨！拉晚了，鱼跑了啦！",
			"🎣 笨！鱼逃竿了！空竿好耶！",
			"🎣 哎呀，拉晚了，鱼已经跑掉惹qwq",
		}))
		return
	}

	// 计算奖励
	typ := util.Lottery([]int{7380, 2000, 500, 60, 40, 10, 10})
	var goldType int
	var gold int
	var diamond int
	var itemId uint
	var itemCount int
	switch typ {
	case DrawGold:
		goldType = util.RandInt(0, 2)
		switch goldType {
		case 0:
			gold = util.RandInt(100, 300)
		case 1:
			gold = util.RandInt(300, 500)
		case 2:
			gold = util.RandInt(1, 100)
		}
	case DrawDiamond:
		diamond = 1
	case DrawGoldTreasure:
		gold = 6666
	case DrawDiamondTreasure:
		diamond = 6
	case DrawGoldGod:
		gold = 66666
	case DrawDiamondGod:
		diamond = 66
	case DrawExpFruit:
		r := util.RandInt(1, 10)
		if r <= 8 {
			itemId = uint(util.RandInt(10001, 10008))
			itemCount = 1
		} else {
			itemId = uint(util.RandInt(10009, 10010))
			itemCount = 1
		}
	}

	// 发放奖励
	userReward := reward.NewUserReward(uid)
	userReward.AddInfo(&tb.UserInfo{
		Gold:    gold,
		Diamond: diamond,
	})
	if itemId != 0 {
		userReward.AddItem(itemId, itemCount)
	}
	if err = userReward.Save(); err != nil {
		return
	}

	// 清除缓存
	fishCache.Delete(uid)

	var item tb.Item
	if itemId != 0 {
		db.DB.Model(&tb.Item{}).First(&item, itemId)
	}

	res = &dto.Draw{
		Type:      typ,
		GoldType:  goldType,
		Gold:      gold,
		Diamond:   diamond,
		Item:      &item,
		ItemCount: itemCount,
	}
	return
}

const (
	DrawGold = iota
	DrawExpFruit
	DrawDiamond
	DrawGoldTreasure
	DrawDiamondTreasure
	DrawGoldGod
	DrawDiamondGod
)
const (
	DrawGoldFish = iota
	DrawGoldBigFish
	DrawGoldMiniFish
)
