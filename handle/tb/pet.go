package tb

import "gorm.io/gorm"

type Pet struct {
	ID        uint    `gorm:"primarykey;autoIncrement:true;comment:宠物id"`
	Name      string  `gorm:"not null;type:varchar(8);comment:宠物名称"`
	Image     string  `gorm:"not null;type:varchar(128);comment:图片路径"`
	FeatureId uint    `gorm:"not null;default:18;comment:宠物属性id"`
	Feature   Feature `gorm:"->"`
	HP        uint    `gorm:"not null;comment:基础生命值"`
	Atk       uint    `gorm:"not null;comment:基础物理攻击力"`
	Def       uint    `gorm:"not null;comment:基础物理防御力"`
	Mak       uint    `gorm:"not null;comment:基础魔法攻击力"`
	Mdf       uint    `gorm:"not null;comment:基础魔法防御力"`
	Spd       uint    `gorm:"not null;comment:基础速度"`
	Crt       float64 `gorm:"not null;default:0.2;comment:基础暴击率"`
	CriDmg    float64 `gorm:"not null;default:2.0;comment:基础暴击伤害"`
	ExpLevel  uint    `gorm:"not null;comment:经验级别"`
	EvoPetId  uint    `gorm:"not null;comment:进化目标宠物ID"`
	EvoPet    *Pet    `gorm:"->"`
	EvoLv     uint    `gorm:"not null;comment:进化所需等级"`
	EvoEx     bool    `gorm:"not null;comment:是否超进化(进化后等级归1)"`
	Catch     bool    `gorm:"not null;comment:是否可捕捉"`
}

type Feature struct {
	ID        uint        `gorm:"primarykey;autoIncrement:true;comment:属性id"`
	Emoji     string      `gorm:"not null;type:varchar(1);comment:属性emoji"`
	Name      string      `gorm:"not null;type:varchar(2);comment:属性名称"`
	HpAdd     float64     `gorm:"not null;comment:生命加成比例"`
	AtkAdd    float64     `gorm:"not null;comment:物攻加成比例"`
	DefAdd    float64     `gorm:"not null;comment:物防加成比例"`
	MakAdd    float64     `gorm:"not null;comment:魔攻加成比例"`
	MdfAdd    float64     `gorm:"not null;comment:魔抗加成比例"`
	SpdAdd    float64     `gorm:"not null;comment:速度加成比例"`
	CrtAdd    float64     `gorm:"not null;comment:暴击率加成"`
	CriDmgAdd float64     `gorm:"not null;comment:暴击伤害加成"`
	Resist    StringSlice `gorm:"not null;type:varchar(255);comment:克制属性"`
	Rebound   StringSlice `gorm:"not null;type:varchar(255);comment:抵抗属性"`
}

type Character struct {
	ID        uint    `gorm:"primarykey;autoIncrement:true;comment:性格id"`
	Emoji     string  `gorm:"not null;type:varchar(3);comment:性格emoji"`
	Name      string  `gorm:"not null;type:varchar(2);comment:性格名称"`
	HpAdd     float64 `gorm:"not null;comment:生命加成比例"`
	AtkAdd    float64 `gorm:"not null;comment:物攻加成比例"`
	DefAdd    float64 `gorm:"not null;comment:物防加成比例"`
	MakAdd    float64 `gorm:"not null;comment:魔攻加成比例"`
	MdfAdd    float64 `gorm:"not null;comment:魔抗加成比例"`
	SpdAdd    float64 `gorm:"not null;comment:速度加成比例"`
	CrtAdd    float64 `gorm:"not null;comment:暴击率加成"`
	CriDmgAdd float64 `gorm:"not null;comment:暴击伤害加成"`
	Effect    string  `gorm:"not null;type:varchar(255);comment:性格影响"`
}

type AdventureMap struct {
	ID        uint           `gorm:"primarykey;autoIncrement:true;comment:地图id"`
	Emoji     string         `gorm:"not null;type:varchar(1);comment:地图emoji"`
	Name      string         `gorm:"not null;type:varchar(255);comment:地图名称"`
	Order     uint           `gorm:"not null;default:0;comment:排序"`
	Type      uint           `gorm:"not null;default:0;comment:对手类型(0: 普通敌人，1: boss)"`
	MinLv     uint           `gorm:"not null;comment:对手最低等级"`
	MaxLv     uint           `gorm:"not null;comment:对手最高等级"`
	PetList   StringSlice    `gorm:"not null;comment:宠物列表(ID)"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
