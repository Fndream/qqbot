package reward

import (
	"errors"
	"gorm.io/gorm"
	"qqbot/handle/tb"
)

type UserReward struct {
	db        *gorm.DB
	info      *tb.UserInfo
	items     map[uint]int
	haveItems map[uint]int
	pets      []*tb.UserPet
}

func NewUserReward(uid string) *UserReward {
	return NewUserRewardTx(uid, nil)
}

func NewUserRewardTx(uid string, tx *gorm.DB) *UserReward {
	return &UserReward{
		db:        DB(tx),
		info:      &tb.UserInfo{ID: uid},
		items:     map[uint]int{},
		haveItems: map[uint]int{},
		pets:      []*tb.UserPet{},
	}
}

func (u *UserReward) AddInfo(add *tb.UserInfo) *UserReward {
	u.info.Gold += add.Gold
	u.info.Diamond += add.Diamond
	u.info.Strive += add.Strive
	return u
}

func (u *UserReward) AddItem(itemId uint, count int) *UserReward {
	u.items[itemId] += count
	return u
}

func (u *UserReward) SubItem(itemId uint, count int, have int) *UserReward {
	u.items[itemId] += -count
	u.haveItems[itemId] = have
	return u
}

func (u *UserReward) AddItemMap(items map[uint]int) *UserReward {
	for k, v := range items {
		u.items[k] += v
	}
	return u
}

func (u *UserReward) AddPets(pets ...*tb.UserPet) *UserReward {
	u.pets = append(u.pets, pets...)
	return u
}

func (u *UserReward) Save() error {
	return u.db.Transaction(func(tx *gorm.DB) error {
		var r *gorm.DB
		sql := map[string]interface{}{}
		if u.info.Gold != 0 {
			sql["gold"] = gorm.Expr("`gold` + ?", u.info.Gold)
		}
		if u.info.Diamond != 0 {
			sql["diamond"] = gorm.Expr("`diamond` + ?", u.info.Diamond)
		}
		if u.info.Strive != 0 {
			sql["strive"] = gorm.Expr("`strive` + ?", u.info.Strive)
		}
		if len(sql) > 0 {
			if r = tx.Model(u.info).Updates(sql); r.Error != nil {
				return r.Error
			}
			if r.RowsAffected == 0 {
				if r = tx.Create(u.info); r.Error != nil {
					return r.Error
				}
			}
		}

		for itemId, add := range u.items {
			if add == 0 {
				continue
			}
			if u.haveItems[itemId]+add == 0 {
				if r = u.db.Delete(&tb.UserItem{ID: u.info.ID, ItemId: itemId}); r.Error != nil {
					return r.Error
				}
				continue
			}
			if r = u.db.Model(&tb.UserItem{ID: u.info.ID, ItemId: itemId}).Update("count", gorm.Expr("`count` + ?", add)); r.Error != nil {
				return r.Error
			}
			if r.RowsAffected == 0 && add > 0 {
				r = u.db.Create(&tb.UserItem{ID: u.info.ID, ItemId: itemId, Count: add})
				if r.Error != nil {
					return r.Error
				}
			}
		}

		for _, pet := range u.pets {
			var ser tb.UserPet
			if r = u.db.Model(&ser).Select("`serial`").Where(&tb.UserPet{UID: u.info.ID}).Order("`serial` desc").First(&ser); r.Error != nil && !errors.Is(r.Error, gorm.ErrRecordNotFound) {
				return r.Error
			}
			pet.Serial = ser.Serial + 1
			if r = u.db.Create(pet); r.Error != nil {
				return r.Error
			}
		}
		return nil
	})
}
