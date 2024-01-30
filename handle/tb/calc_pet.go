package tb

import (
	"fmt"
	"qqbot/handle/tb/emoji"
	"qqbot/handle/util/calc"
)

func (p *Pet) EmojiEvo() string {
	if p.EvoEx {
		return emoji.EvoEx
	}
	return emoji.Evo
}

func (p *Pet) View() string {
	view := fmt.Sprintf("%v %v\n", p.Feature.Emoji, p.Name)
	view += fmt.Sprintf("%v 精力：%v\n", emoji.HP, p.HP)
	view += fmt.Sprintf("%v 攻击：%v\n", emoji.Atk, p.Atk)
	view += fmt.Sprintf("%v 防御：%v\n", emoji.Def, p.Def)
	view += fmt.Sprintf("%v 魔攻：%v\n", emoji.Mak, p.Mak)
	view += fmt.Sprintf("%v 魔抗：%v\n", emoji.Mdf, p.Mdf)
	view += fmt.Sprintf("%v 速度：%v\n", emoji.Spd, p.Spd)
	view += fmt.Sprintf("%v 暴击：%v%%\n", emoji.Crt, calc.Mul(p.Crt, 100))
	view += fmt.Sprintf("%v 爆伤：%v%%\n", emoji.CriDmg, calc.NewNum(p.CriDmg).Sub(1).Mul(100))
	if p.EvoPetId != 0 {
		view += fmt.Sprintf("%v 下一进化：%v\n", p.EmojiEvo(), p.EvoPet.Name)
		view += fmt.Sprintf("%v 进化等级：%v\n", EmojiLv(p.EvoLv), p.EvoLv)
	}
	if p.Catch {
		view += fmt.Sprintf("🏵️ 可捕捉\n")
	}
	return view[:len(view)-1]
}
