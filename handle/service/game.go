package service

import (
	"errors"
	"gorm.io/gorm"
	"qqbot/handle/dto"
	"qqbot/handle/reward"
	"qqbot/handle/tb"
	"qqbot/handle/util"
	"qqbot/pkg/db"
)

// Gambling 赌博
func Gambling(uid string, count int) (res *dto.Gambling, err error) {
	// 获取玩家金币数据
	var userInfo tb.UserInfo
	if e := db.DB.Model(&tb.UserInfo{}).First(&userInfo, uid).Error; e != nil && !errors.Is(e, gorm.ErrRecordNotFound) {
		return nil, e
	}

	// 判断玩家金币是否足够
	if userInfo.Gold < count {
		return nil, errors.New("🍱 金币不足")
	}

	// 计算奖励
	success := util.RandBool()
	var rewardGold int
	var rewardDiamond int

	if success {
		rewardGold = count
	} else {
		rewardGold = -count
	}

	if success && util.RandIntn(100) < 10 {
		rewardDiamond = 1
	}

	// 发放奖励
	if err = reward.NewUserReward(uid).AddInfo(&tb.UserInfo{
		Gold:    rewardGold,
		Diamond: rewardDiamond,
	}).Save(); err != nil {
		return
	}

	res = &dto.Gambling{
		Success: success,
		Gold:    count,
		Diamond: rewardDiamond,
	}
	return
}
