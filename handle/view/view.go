package view

import (
	"fmt"
	"qqbot/handle/tb"
	"qqbot/handle/tb/emoji"
)

// 封装复用View

func LevelUpView(oldPet *tb.UserPet, newPet *tb.UserPet) string {
	oldPet.Calc()
	newPet.Calc()
	msg := fmt.Sprintf("🔥 宠物升级！\n%v Lv %v -> %v\n", tb.EmojiLv(newPet.Lv), oldPet.Lv, newPet.Lv)
	msg += PetBattleAttrCompareView(oldPet, newPet)
	return msg
}

func PetBattleAttrCompareView(oldPet *tb.UserPet, newPet *tb.UserPet) string {
	oldPet.Calc()
	newPet.Calc()
	msg := fmt.Sprintf("%v 精力：%v -> %v\n", emoji.HP, oldPet.Hp, newPet.Hp)
	msg += fmt.Sprintf("%v 攻击：%v -> %v\n", emoji.Atk, oldPet.Atk, newPet.Atk)
	msg += fmt.Sprintf("%v 防御：%v -> %v\n", emoji.Def, oldPet.Def, newPet.Def)
	msg += fmt.Sprintf("%v 魔攻：%v -> %v\n", emoji.Mak, oldPet.Mak, newPet.Mak)
	msg += fmt.Sprintf("%v 魔抗：%v -> %v\n", emoji.Mdf, oldPet.Mdf, newPet.Mdf)
	msg += fmt.Sprintf("%v 速度：%v -> %v\n", emoji.Spd, oldPet.Spd, newPet.Spd)
	return msg
}
