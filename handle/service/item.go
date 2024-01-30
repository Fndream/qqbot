package service

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"qqbot/handle/dto"
	"qqbot/handle/reward"
	"qqbot/handle/tb"
	"qqbot/handle/util/calc"
	"qqbot/pkg/db"
)

// 经验果

func UseExpFruitByPetSerial(uid string, userItem *tb.UserItem, count int, serial uint) (res *dto.UseExpFruit, err error) {
	userPet, err := QueryUserPetBySerial(uid, serial)
	if err != nil {
		return nil, err
	}
	return useExpFruit(uid, userItem, count, userPet)
}

func UseExpFruitByPetName(uid string, userItem *tb.UserItem, count int, petName string) (res *dto.UseExpFruit, err error) {
	userPet, err := QueryUserPetByName(uid, petName)
	if err != nil {
		return nil, err
	}
	return useExpFruit(uid, userItem, count, userPet)
}

func useExpFruit(uid string, userItem *tb.UserItem, count int, userPet *tb.UserPet) (res *dto.UseExpFruit, err error) {
	if userPet.Lv >= 100 {
		return nil, fmt.Errorf("📕 %v已满级", userPet.Pet.Name)
	}

	expLevel := tb.ExpLevelOf(userPet.Pet.ExpLevel)
	if userPet.Exp > uint(expLevel[100]-expLevel[userPet.Lv]) {
		return nil, errors.New("🔵 经验溢出")
	}

	var rewardExp int
	switch userItem.Item.ID {
	case 10001:
		rewardExp = 1000
	case 10002:
		rewardExp = 2000
	case 10003:
		rewardExp = 3000
	case 10004:
		rewardExp = 5000
	case 10005:
		rewardExp = 8000
	case 10006:
		rewardExp = 10000
	case 10007:
		rewardExp = 20000
	case 10008:
		rewardExp = 50000
	case 10009:
		rewardExp = 100000
	case 10010:
		rewardExp = 200000
	case 10011:
		rewardExp = 500000
	default:
		return nil, errors.New("unknown fruit")
	}

	// 避免浪费
	if rewardExp*count > expLevel[100]-expLevel[userPet.Lv]-int(userPet.Exp) {
		count = (expLevel[100] - expLevel[userPet.Lv] - int(userPet.Exp)) / rewardExp
	}

	res = &dto.UseExpFruit{Use: &dto.Use{UserItem: userItem, UseCount: count, OldPet: userPet}, AddExp: rewardExp}
	if err = db.DB.Transaction(func(tx *gorm.DB) error {
		petReward := reward.NewUserPetRewardTx(userPet, tx).AddExp(uint(rewardExp * count))
		res.NewPet = petReward.GetNewPet()
		res.AddLv = int(res.NewPet.Lv - res.OldPet.Lv)
		if e := petReward.Save(); e != nil {
			return e
		}
		if e := reward.NewUserRewardTx(uid, tx).SubItem(userItem.Item.ID, count, userItem.Count).Save(); e != nil {
			return e
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return
}

// 努力果

func UseStvFruitByPetSerial(uid string, userItem *tb.UserItem, count int, serial uint) (res *dto.UseStvFruit, err error) {
	userPet, err := QueryUserPetBySerial(uid, serial)
	if err != nil {
		return nil, err
	}
	return useStvFruit(uid, userItem, count, userPet)
}

func UseStvFruitByPetName(uid string, userItem *tb.UserItem, count int, petName string) (res *dto.UseStvFruit, err error) {
	userPet, err := QueryUserPetByName(uid, petName)
	if err != nil {
		return nil, err
	}
	return useStvFruit(uid, userItem, count, userPet)
}

func useStvFruit(uid string, userItem *tb.UserItem, count int, userPet *tb.UserPet) (res *dto.UseStvFruit, err error) {
	if userPet.TotalStv()+uint(count) > 510 {
		return nil, fmt.Errorf("%v 无法为%v使用这么多%v，超出宠物努力极限", userItem.Item.Emoji, userPet.Pet.Name, userItem.Item.Name)
	}

	var curStv uint
	var attr tb.PetAttr
	switch userItem.Item.ID {
	case 10101:
		curStv = userPet.StvHp
		attr = tb.HP
	case 10102:
		curStv = userPet.StvAtk
		attr = tb.Atk
	case 10103:
		curStv = userPet.StvDef
		attr = tb.Def
	case 10104:
		curStv = userPet.StvMak
		attr = tb.Mak
	case 10105:
		curStv = userPet.StvMdf
		attr = tb.Mdf
	case 10106:
		curStv = userPet.StvSpd
		attr = tb.Spd
	default:
		return nil, errors.New("unknown strive fruit")
	}

	if curStv+uint(count) > 255 {
		return nil, fmt.Errorf("%v 无法为%v使用这么多%v，超出宠物单属性努力极限", userItem.Item.Emoji, userPet.Pet.Name, userItem.Item.Name)
	}

	res = &dto.UseStvFruit{Use: &dto.Use{UserItem: userItem, UseCount: count, OldPet: userPet}}
	if err = db.DB.Transaction(func(tx *gorm.DB) error {
		petReward := reward.NewUserPetRewardTx(userPet, tx).AddStv(attr, count)
		res.NewPet = petReward.GetNewPet()
		if e := petReward.Save(); e != nil {
			return e
		}
		if e := reward.NewUserRewardTx(uid, tx).SubItem(userItem.Item.ID, count, userItem.Count).Save(); e != nil {
			return e
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return
}

// 遗忘果

func UseForgetFruitByPetSerial(uid string, userItem *tb.UserItem, serial uint) (res *dto.UseForgetFruit, err error) {
	userPet, err := QueryUserPetBySerial(uid, serial)
	if err != nil {
		return nil, err
	}
	return useForgetFruit(uid, userItem, userPet)
}

func UseForgetFruitByPetName(uid string, userItem *tb.UserItem, petName string) (res *dto.UseForgetFruit, err error) {
	userPet, err := QueryUserPetByName(uid, petName)
	if err != nil {
		return nil, err
	}
	return useForgetFruit(uid, userItem, userPet)
}

func useForgetFruit(uid string, userItem *tb.UserItem, userPet *tb.UserPet) (res *dto.UseForgetFruit, err error) {
	if userPet.TotalStv() <= 0 {
		return nil, fmt.Errorf("%v 无法为%v使用%v，没有需要遗忘的努力值", userItem.Item.Emoji, userPet.Pet.Name, userItem.Item.Name)
	}
	res = &dto.UseForgetFruit{Use: &dto.Use{UserItem: userItem, UseCount: 1, OldPet: userPet}}
	if err = db.DB.Transaction(func(tx *gorm.DB) error {
		petReward := reward.NewUserPetRewardTx(userPet, tx).ClearStv()
		res.NewPet = petReward.GetNewPet()
		if e := petReward.Save(); e != nil {
			return e
		}
		returnStrive := int(calc.Mul(float64(res.OldPet.TotalStv()), 0.8))
		res.ReturnStrive = returnStrive
		if e := reward.NewUserRewardTx(uid, tx).SubItem(userItem.Item.ID, 1, userItem.Count).AddInfo(&tb.UserInfo{Strive: returnStrive}).Save(); e != nil {
			return e
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return
}

// 果儿糖

func UseTalentFruitByPetSerial(uid string, userItem *tb.UserItem, serial uint) (res *dto.UseTalentFruit, err error) {
	userPet, err := QueryUserPetBySerial(uid, serial)
	if err != nil {
		return nil, err
	}
	return useTalentFruit(uid, userItem, userPet)
}

func UseTalentFruitByPetName(uid string, userItem *tb.UserItem, petName string) (res *dto.UseTalentFruit, err error) {
	userPet, err := QueryUserPetByName(uid, petName)
	if err != nil {
		return nil, err
	}
	return useTalentFruit(uid, userItem, userPet)
}

func useTalentFruit(uid string, userItem *tb.UserItem, userPet *tb.UserPet) (res *dto.UseTalentFruit, err error) {
	res = &dto.UseTalentFruit{Use: &dto.Use{UserItem: userItem, UseCount: 1, OldPet: userPet}}
	if err = db.DB.Transaction(func(tx *gorm.DB) error {
		petReward := reward.NewUserPetRewardTx(userPet, tx).RandTalent()
		res.NewPet = petReward.GetNewPet()
		if e := petReward.Save(); e != nil {
			return e
		}
		if e := reward.NewUserRewardTx(uid, tx).SubItem(userItem.Item.ID, 1, userItem.Count).Save(); e != nil {
			return e
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return
}

// 圣灵果

func UseEvoFruitByPetSerial(uid string, userItem *tb.UserItem, serial uint) (res *dto.UseEvoFruit, err error) {
	userPet, err := QueryUserPetBySerial(uid, serial)
	if err != nil {
		return nil, err
	}
	return useEvoFruit(uid, userItem, userPet)
}

func UseEvoFruitByPetName(uid string, userItem *tb.UserItem, petName string) (res *dto.UseEvoFruit, err error) {
	userPet, err := QueryUserPetByName(uid, petName)
	if err != nil {
		return nil, err
	}
	return useEvoFruit(uid, userItem, userPet)
}

func useEvoFruit(uid string, userItem *tb.UserItem, userPet *tb.UserPet) (res *dto.UseEvoFruit, err error) {
	if userPet.Pet.EvoPetId == 0 {
		return nil, fmt.Errorf("%v %v无法使用%v再进化", userPet.Pet.Feature.Emoji, userPet.Pet.Name, userItem.Item.Name)
	}
	if userPet.Lv < userPet.Pet.EvoLv {
		return nil, fmt.Errorf("%v %v未达到进化等级", userPet.Pet.Feature.Emoji, userPet.Pet.Name)
	}
	res = &dto.UseEvoFruit{Use: &dto.Use{UserItem: userItem, UseCount: 1, OldPet: userPet}}
	if err = db.DB.Transaction(func(tx *gorm.DB) error {
		petReward := reward.NewUserPetRewardTx(userPet, tx).Evo()
		if e := petReward.Save(); e != nil {
			return e
		}

		newPet, e := GetUserPetById(userPet.ID)
		if e != nil {
			return e
		}
		if newPet == nil {
			return errors.New("not found update userPet for evolved")
		}
		res.NewPet = newPet

		if e := reward.NewUserRewardTx(uid, tx).SubItem(userItem.Item.ID, 1, userItem.Count).Save(); e != nil {
			return e
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return
}

// 归初果

func UseCharacterFruitByPetSerial(uid string, userItem *tb.UserItem, serial uint) (res *dto.UseCharacterFruit, err error) {
	userPet, err := QueryUserPetBySerial(uid, serial)
	if err != nil {
		return nil, err
	}
	return useCharacterFruit(uid, userItem, userPet)
}

func UseCharacterFruitByPetName(uid string, userItem *tb.UserItem, petName string) (res *dto.UseCharacterFruit, err error) {
	userPet, err := QueryUserPetByName(uid, petName)
	if err != nil {
		return nil, err
	}
	return useCharacterFruit(uid, userItem, userPet)
}

func useCharacterFruit(uid string, userItem *tb.UserItem, userPet *tb.UserPet) (res *dto.UseCharacterFruit, err error) {
	res = &dto.UseCharacterFruit{Use: &dto.Use{UserItem: userItem, UseCount: 1, OldPet: userPet}}
	if err = db.DB.Transaction(func(tx *gorm.DB) error {
		petReward := reward.NewUserPetRewardTx(userPet, tx).RandCharacter()
		if e := petReward.Save(); e != nil {
			return e
		}

		newPet, e := GetUserPetById(userPet.ID)
		if e != nil {
			return e
		}
		if newPet == nil {
			return errors.New("not found update userPet for character")
		}
		res.NewPet = newPet

		if e := reward.NewUserRewardTx(uid, tx).SubItem(userItem.Item.ID, 1, userItem.Count).Save(); e != nil {
			return e
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return
}
