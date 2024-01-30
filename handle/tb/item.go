package tb

type Item struct {
	ID          uint   `gorm:"primarykey;autoIncrement:true"`
	Emoji       string `gorm:"not null;type:varchar(1);comment:物品emoji"`
	Name        string `gorm:"not null;type:varchar(6);uniqueIndex:idx_items_name;comment:物品名称"`
	Description string `gorm:"not null;type:varchar(32);comment:物品描述"`
}

type Good struct {
	ItemID      uint `gorm:"primarykey;autoIncrement:false;comment:物品ID"`
	Item        Item `gorm:"->"`
	BuyType     uint `gorm:"not null;default:0;comment:购买类型(0:&&，1:||)"`
	BuyGold     uint `gorm:"not null;default:0;comment:购买所需金币"`
	BuyDiamond  uint `gorm:"not null;default:0;comment:购买所需钻石"`
	SellGold    uint `gorm:"not null;default:0;comment:出售可得金币"`
	SellDiamond uint `gorm:"not null;default:0;comment:购买可得钻石"`
}
