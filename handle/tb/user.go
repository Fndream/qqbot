package tb

import (
	"gorm.io/gorm"
	"qqbot/handle/util"
	"time"
)

type UserSign struct {
	ID         string    `gorm:"primarykey;type:varchar(20);comment:用户id"`
	Count      uint      `gorm:"not null;default:0;comment:签到次数"`
	Continuous uint      `gorm:"not null;default:0;comment:连签次数"`
	CreatedAt  time.Time `gorm:"not null;comment:创建时间"`
	UpdatedAt  time.Time `gorm:"not null;comment:更新时间"`
}

type UserInfo struct {
	ID      string `gorm:"primarykey;type:varchar(20);comment:用户id"`
	Gold    int    `gorm:"not null;default:0;comment:金币"`
	Diamond int    `gorm:"not null;default:0;comment:钻石"`
	Strive  int    `gorm:"not null;default:0;comment:努力点"`
}

type UserItem struct {
	ID     string `gorm:"primarykey;type:varchar(20);comment:用户id"`
	ItemId uint   `gorm:"primarykey;comment:物品id"`
	Item   Item   `gorm:"->"`
	Count  int    `gorm:"not null;default:0;comment:数量"`
}

type UserPet struct {
	ID               string         `gorm:"primarykey;type:varchar(20);comment:用户宠物id"`
	UID              string         `gorm:"not null;uniqueIndex:idx_UserPet_uid_serial;index:idx_UserPet_uid_petId;type:varchar(20);comment:用户id"`
	Serial           uint           `gorm:"not null;uniqueIndex:idx_UserPet_uid_serial;comment:序号"`
	PetId            uint           `gorm:"not null;index:idx_UserPet_uid_petId;comment:宠物id"`
	Pet              Pet            `gorm:"->"`
	ChaId            uint           `gorm:"not null;comment:性格id"`
	Cha              Character      `gorm:"->"`
	Lv               uint           `gorm:"not null;default:1;comment:等级"`
	Exp              uint           `gorm:"not null;default:0;comment:经验"`
	TaleHp           uint           `gorm:"not null;default:0;comment:生命天赋"`
	TaleAtk          uint           `gorm:"not null;default:0;comment:攻击天赋"`
	TaleDef          uint           `gorm:"not null;default:0;comment:防御天赋"`
	TaleMak          uint           `gorm:"not null;default:0;comment:魔攻天赋"`
	TaleMdf          uint           `gorm:"not null;default:0;comment:魔抗天赋"`
	TaleSpd          uint           `gorm:"not null;default:0;comment:速度天赋"`
	StvHp            uint           `gorm:"not null;default:0;comment:生命努力值"`
	StvAtk           uint           `gorm:"not null;default:0;comment:攻击努力值"`
	StvDef           uint           `gorm:"not null;default:0;comment:防御努力值"`
	StvMak           uint           `gorm:"not null;default:0;comment:魔攻努力值"`
	StvMdf           uint           `gorm:"not null;default:0;comment:魔抗努力值"`
	StvSpd           uint           `gorm:"not null;default:0;comment:速度努力值"`
	Status           uint           `gorm:"not null;default:0;comment:状态：0正常，1昏睡，2治疗"`
	TreatEndTime     time.Time      `gorm:"not null;comment:治疗结束时间"`
	CreatedAt        time.Time      `gorm:"not null;comment:创建时间"`
	UpdatedAt        time.Time      `gorm:"not null;comment:更新时间"`
	DeletedAt        gorm.DeletedAt `gorm:"index"`
	BattleAttributes `gorm:"-"`
}

func (p *UserPet) BeforeCreate(tx *gorm.DB) error {
	if p.ID == "" {
		p.ID = util.NewId()
	}
	return nil
}

type UserPetBag struct {
	ID      string  `gorm:"primaryKey;comment:用户id"`
	Slot1Id string  `gorm:"not null;default:;comment:1号槽位"`
	Slot2Id string  `gorm:"not null;default:;comment:2号槽位"`
	Slot3Id string  `gorm:"not null;default:;comment:3号槽位"`
	Slot4Id string  `gorm:"not null;default:;comment:4号槽位"`
	Slot5Id string  `gorm:"not null;default:;comment:5号槽位"`
	Slot6Id string  `gorm:"not null;default:;comment:6号槽位"`
	Slot1   UserPet `gorm:"->"`
	Slot2   UserPet `gorm:"->"`
	Slot3   UserPet `gorm:"->"`
	Slot4   UserPet `gorm:"->"`
	Slot5   UserPet `gorm:"->"`
	Slot6   UserPet `gorm:"->"`
}
