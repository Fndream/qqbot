package reward

import (
	"gorm.io/gorm"
	"qqbot/handle/tb"
)

type UserPetReward struct {
	db     *gorm.DB
	oldPet *tb.UserPet
	newPet *tb.UserPet
}

func NewUserPetReward(pet *tb.UserPet) *UserPetReward {
	return NewUserPetRewardTx(pet, nil)
}

func NewUserPetRewardTx(pet *tb.UserPet, tx *gorm.DB) *UserPetReward {
	newPet := *pet
	return &UserPetReward{
		db:     DB(tx),
		oldPet: pet,
		newPet: &newPet,
	}
}

func (u *UserPetReward) GetOldPet() *tb.UserPet {
	return u.oldPet
}

func (u *UserPetReward) GetNewPet() *tb.UserPet {
	return u.newPet
}

func (u *UserPetReward) RandCharacter() *UserPetReward {
	u.newPet.RandCharacter()
	return u
}

func (u *UserPetReward) RandTalent() *UserPetReward {
	u.newPet.RandTalent()
	return u
}

func (u *UserPetReward) AddExp(exp uint) *UserPetReward {
	u.newPet.Exp += exp
	u.levelUp()
	return u
}

func (u *UserPetReward) levelUp() {
	expLevel := tb.ExpLevelOf(u.newPet.Pet.ExpLevel)
	for u.newPet.Lv < 100 && int(u.newPet.Exp) >= expLevel[u.newPet.Lv+1]-expLevel[u.newPet.Lv] {
		u.newPet.Lv++
		u.newPet.Exp -= uint(expLevel[u.newPet.Lv+1] - expLevel[u.newPet.Lv])
	}
}

func (u *UserPetReward) AddStv(attr tb.PetAttr, add int) *UserPetReward {
	switch attr {
	case tb.HP:
		u.newPet.StvHp += uint(add)
	case tb.Atk:
		u.newPet.StvAtk += uint(add)
	case tb.Def:
		u.newPet.StvDef += uint(add)
	case tb.Mak:
		u.newPet.StvMak += uint(add)
	case tb.Mdf:
		u.newPet.StvMdf += uint(add)
	case tb.Spd:
		u.newPet.StvSpd += uint(add)
	}
	return u
}

func (u *UserPetReward) ClearStv() *UserPetReward {
	u.newPet.ClearStv()
	return u
}

func (u *UserPetReward) Evo() *UserPetReward {
	u.newPet.Evo()
	return u
}

func (u *UserPetReward) Save() error {
	updPet := tb.UserPet{}

	if u.newPet.Lv != u.oldPet.Lv {
		updPet.Lv = u.newPet.Lv
	}
	if u.newPet.Exp != u.oldPet.Exp {
		updPet.Exp = u.newPet.Exp
	}

	if u.newPet.TaleHp != u.oldPet.TaleHp {
		updPet.TaleHp = u.newPet.TaleHp
	}
	if u.newPet.TaleAtk != u.oldPet.TaleAtk {
		updPet.TaleAtk = u.newPet.TaleAtk
	}
	if u.newPet.TaleDef != u.oldPet.TaleDef {
		updPet.TaleDef = u.newPet.TaleDef
	}
	if u.newPet.TaleMak != u.oldPet.TaleMak {
		updPet.TaleMak = u.newPet.TaleMak
	}
	if u.newPet.TaleMdf != u.oldPet.TaleMdf {
		updPet.TaleMdf = u.newPet.TaleMdf
	}
	if u.newPet.TaleSpd != u.oldPet.TaleSpd {
		updPet.TaleSpd = u.newPet.TaleSpd
	}

	if u.newPet.StvHp != u.oldPet.StvHp {
		updPet.StvHp = u.newPet.StvHp
	}
	if u.newPet.StvAtk != u.oldPet.StvAtk {
		updPet.StvAtk = u.newPet.StvAtk
	}
	if u.newPet.StvDef != u.oldPet.StvDef {
		updPet.StvDef = u.newPet.StvDef
	}
	if u.newPet.StvMak != u.oldPet.StvMak {
		updPet.StvMak = u.newPet.StvMak
	}
	if u.newPet.StvMdf != u.oldPet.StvMdf {
		updPet.StvMdf = u.newPet.StvMdf
	}
	if u.newPet.StvSpd != u.oldPet.StvSpd {
		updPet.StvSpd = u.newPet.StvSpd
	}

	if u.newPet.PetId != u.oldPet.PetId {
		updPet.PetId = u.newPet.PetId
	}
	if u.newPet.ChaId != u.oldPet.ChaId {
		updPet.ChaId = u.newPet.ChaId
	}

	if u.newPet.Status != u.oldPet.Status {
		updPet.Status = u.newPet.Status
	}
	if u.newPet.TreatEndTime != u.oldPet.TreatEndTime {
		updPet.TreatEndTime = u.newPet.TreatEndTime
	}

	if e := u.db.Model(u.newPet).Updates(&updPet).Error; e != nil {
		return e
	}
	return nil
}
