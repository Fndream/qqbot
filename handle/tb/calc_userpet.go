package tb

import (
	"fmt"
	"qqbot/handle/tb/emoji"
	"qqbot/handle/util"
	"qqbot/handle/util/calc"
	"strconv"
	"time"
)

type PetAttr int

var (
	ExpLevel1 = []int{0, 1, 14, 70, 156, 327, 543, 887, 1291, 1866, 2516, 3380, 4334, 5545, 6861, 8477, 10213, 12292, 15298, 18828, 22558, 26881, 32237, 38285, 45065, 52601, 60948, 70146, 80216, 91216, 103186, 116733, 131349, 147074, 166236, 186756, 208652, 231968, 256778, 283097, 311001, 341397, 373457, 407261, 442818, 480172, 519406, 560526, 603576, 648642, 695727, 746037, 798477, 853091, 909971, 969114, 1030564, 1094416, 1160664, 1229352, 1300578, 1375674, 1453365, 1533753, 1616825, 1702685, 1791318, 1882768, 1977142, 2074422, 2174652, 2279507, 2387423, 2498444, 2612684, 2730118, 2850862, 2974889, 3102243, 3233043, 3367259, 3506723, 3649792, 3796432, 3946687, 4100683, 4258383, 4419831, 4585156, 4754318, 4927448, 5106516, 5289576, 5476763, 5668031, 5863424, 6063080, 6266950, 6475078, 6687605, 6904479}
	ExpLevel2 = []int{0, 1, 14, 70, 156, 327, 543, 887, 1291, 1866, 2516, 3380, 4334, 5545, 6861, 8477, 10213, 12292, 15298, 18828, 22558, 26881, 32649, 39129, 46361, 54368, 63697, 73917, 85047, 97147, 110257, 124982, 140816, 157799, 178435, 200475, 224717, 250445, 277736, 306602, 337122, 370200, 405008, 441629, 480069, 520372, 563630, 608862, 656112, 705470, 756935, 811717, 868717, 927979, 989599, 1053570, 1121165, 1191277, 1263895, 1339063, 1416884, 1498685, 1583191, 1670509, 1760621, 1853636, 1950987, 2051287, 2154649, 2261049, 2370531, 2484776, 2602214, 2722889, 2846921, 2974279, 3106762, 3242682, 3382083, 3525091, 3671669, 3823649, 3979395, 4138866, 4302106, 4469248, 4642148, 4818972, 4999857, 5184755, 5373805, 5568969, 5768301, 5971944, 6179844, 6392045, 6610817, 6834001, 7061641, 7293887, 7530678}
	ExpLevel3 = []int{0, 1, 14, 70, 156, 327, 543, 887, 1291, 1866, 2516, 3380, 4334, 5545, 6861, 8477, 10213, 12292, 15298, 18828, 22558, 26881, 33061, 39973, 47657, 56606, 66426, 77157, 89347, 102547, 116797, 132700, 149752, 167993, 190103, 214423, 240229, 267565, 297337, 328750, 361886, 397646, 435202, 474640, 515963, 560198, 606474, 654790, 706240, 759890, 815735, 874989, 936549, 1000459, 1066819, 1136825, 1209336, 1284456, 1363444, 1445092, 1529508, 1618014, 1709335, 1803583, 1900735, 2002336, 2106952, 2214627, 2326977, 2442497, 2561231, 2684866, 2811826, 2942155, 3075979, 3214915, 3357460, 3503574, 3655022, 3810238, 3969178, 4133674, 4302097, 4474399, 4650624, 4832790, 5018990, 5209268, 5405713, 5606347, 5811317, 6022577, 6238181, 6458280, 6682812, 6913922, 7149686, 7390038, 7637190, 7889155, 8145863}
	ExpLevel4 = []int{0, 1, 14, 70, 156, 327, 543, 887, 1291, 1866, 2516, 3380, 4334, 5545, 6861, 8477, 10213, 12292, 15298, 18828, 22558, 26881, 33473, 40817, 49405, 58825, 69627, 81380, 94100, 108400, 123790, 140871, 159141, 178640, 202961, 228801, 256953, 286701, 318127, 352087, 387839, 426281, 466585, 508840, 554007, 601191, 651491, 703919, 758519, 816461, 876686, 940412, 1006532, 1075090, 1147375, 1222209, 1300865, 1382245, 1466329, 1554457, 1645468, 1740679, 1838815, 1939993, 2045593, 2154349, 2267683, 2384208, 2504048, 2628688, 2756674, 2889699, 3026181, 3166164, 3311412, 3460272, 3614556, 3772563, 3934337, 4101761, 4273063, 4450075, 4631175, 4816308, 5007373, 5202685, 5404085, 5609739, 5819799, 6036169, 6257059, 6484415, 6716291, 6952846, 7196089, 7444007, 7698887, 7958553, 8223049, 8494733, 8771358}
	ExpLevel5 = []int{0, 1, 14, 70, 156, 327, 543, 887, 1291, 1866, 2516, 3380, 4334, 5545, 6861, 8477, 10213, 12292, 15298, 18828, 22558, 26881, 33885, 42093, 51133, 61495, 72788, 85563, 99343, 114743, 131273, 149532, 169020, 190406, 216201, 244321, 274037, 306197, 340104, 376611, 414979, 456103, 499155, 545166, 593216, 644332, 697650, 754190, 812990, 875224, 939829, 1008027, 1078707, 1153075, 1230100, 1310969, 1394541, 1482181, 1572635, 1667243, 1764849, 1866765, 1971716, 2081210, 2193850, 2311192, 2431791, 2557166, 2685994, 2819754, 2956992, 3099407, 3245411, 3396657, 3551697, 3712135, 3876481, 4046381, 4220202, 4399834, 4583498, 4773026, 4966803, 5166600, 5370650, 5580986, 5795686, 6016716, 6242336, 6474442, 6711252, 6954704, 7202852, 7457920, 7717795, 7984622, 8256494, 8535474, 8819482, 9110885, 9407427}
	ExpLevel6 = []int{0, 1, 14, 70, 156, 327, 543, 887, 1291, 1866, 2516, 3380, 4334, 5545, 6861, 8477, 10213, 12292, 15298, 18828, 22558, 26881, 36357, 47157, 59361, 73020, 88241, 105104, 123654, 144004, 166234, 190383, 216570, 244875, 279514, 316754, 356636, 399248, 444733, 493126, 544574, 599108, 656816, 717851, 782238, 850065, 921491, 996535, 1075285, 1157906, 1244411, 1334969, 1429589, 1528359, 1631454, 1738877, 1850716, 1967152, 2088182, 2213894, 2344475, 2479916, 2620305, 2765835, 2916491, 3072470, 3233753, 3400428, 3572698, 3750538, 3934036, 4123401, 4318602, 4519727, 4726991, 4940357, 5160044, 5386011, 5618346, 5857274, 6102748, 6354856, 6613829, 6879614, 7152299, 7432121, 7719021, 8013087, 8314562, 8623381, 8939791, 9263723, 9595265, 9934670, 10281863, 10636932, 11000136, 11371394, 11750794, 12138601, 12534728}
)

const (
	HP PetAttr = iota
	Atk
	Def
	Mak
	Mdf
	Spd
)

func ExpLevelOf(expLevel uint) []int {
	switch expLevel {
	case 1:
		return ExpLevel1
	case 2:
		return ExpLevel2
	case 3:
		return ExpLevel3
	case 4:
		return ExpLevel4
	case 5:
		return ExpLevel5
	case 6:
		return ExpLevel6
	default:
		panic("unknown exp level")
	}
}

func PetAttrOf(attrName string) PetAttr {
	switch attrName {
	case "精力":
		return HP
	case "攻击":
		return Atk
	case "防御":
		return Def
	case "魔攻":
		return Mak
	case "魔抗", "魔防":
		return Mdf
	case "速度":
		return Spd
	default:
		panic("❓ 未知的属性名：" + attrName)
	}
}

type BattleAttributes struct {
	Hp     int
	Atk    int
	Def    int
	Mak    int
	Mdf    int
	Spd    int
	Crt    float64
	CriDmg float64
}

func EmojiLv(lv uint) string {
	if lv < 40 {
		return emoji.LvGreen
	} else if lv < 70 {
		return emoji.LvBlue
	} else if lv < 100 {
		return emoji.LvOrange
	} else {
		return emoji.LvRed
	}
}

func (p *UserPet) EmojiLv() string {
	return EmojiLv(p.Lv)
}

func (p *UserPet) NextLvExp() int {
	if p.Lv >= 100 {
		return 0
	}
	expLv := ExpLevelOf(p.Pet.ExpLevel)
	return expLv[p.Lv+1] - expLv[p.Lv]
}

func (p *UserPet) EmojiStatus() string {
	switch p.Status {
	case 1:
		return emoji.Lethargy
	case 2:
		if time.Now().Before(p.TreatEndTime) {
			seconds := int(p.TreatEndTime.Sub(time.Now()).Seconds())
			secondsStr := strconv.Itoa(seconds)
			return emoji.Treatment + secondsStr + "s"
		} else {
			return ""
		}
	default:
		return ""
	}
}

func (p *UserPet) Total() int {
	return p.Hp + p.Atk + p.Def + p.Mak + p.Mdf + p.Spd
}

func (p *UserPet) TotalStv() uint {
	return p.StvHp + p.StvAtk + p.StvDef + p.StvMak + p.StvMdf + p.StvSpd
}

func (p *UserPet) RandCharacter() *UserPet {
	newId := p.ChaId
	for newId == p.ChaId {
		newId = util.RandCharacterId()
	}
	p.ChaId = newId
	return p
}

func (p *UserPet) RandTalent() *UserPet {
	p.TaleHp = uint(util.RandInt(1, 31))
	p.TaleAtk = uint(util.RandInt(1, 31))
	p.TaleDef = uint(util.RandInt(1, 31))
	p.TaleMak = uint(util.RandInt(1, 31))
	p.TaleMdf = uint(util.RandInt(1, 31))
	p.TaleSpd = uint(util.RandInt(1, 31))
	return p
}

func (p *UserPet) ClearStv() *UserPet {
	p.StvHp = 0
	p.StvAtk = 0
	p.StvDef = 0
	p.StvMak = 0
	p.StvMdf = 0
	p.StvSpd = 0
	return p
}

func (p *UserPet) Evo() *UserPet {
	if p.Pet.EvoPetId == 0 {
		panic("not evo")
	}
	p.PetId = p.Pet.EvoPetId
	if p.Pet.EvoEx {
		p.Lv = 1
		p.Exp = 0
	}
	return p
}

func (p *UserPet) Calc() *UserPet {
	// 基础
	hp := float64(p.Pet.HP)
	atk := float64(p.Pet.Atk)
	def := float64(p.Pet.Def)
	mak := float64(p.Pet.Mak)
	mdf := float64(p.Pet.Mdf)
	spd := float64(p.Pet.Spd)
	crt := p.Pet.Crt
	criDmg := p.Pet.CriDmg

	// 天赋
	taleHp := float64(p.TaleHp)
	taleAtk := float64(p.TaleAtk)
	taleDef := float64(p.TaleDef)
	taleMak := float64(p.TaleMak)
	taleMdf := float64(p.TaleMdf)
	taleSpd := float64(p.TaleSpd)

	// 属性加成
	addHp := calc.NewNum(p.Pet.Feature.HpAdd).Add(p.Cha.HpAdd).Float64()
	addAtk := calc.NewNum(p.Pet.Feature.AtkAdd).Add(p.Cha.AtkAdd).Float64()
	addDef := calc.NewNum(p.Pet.Feature.DefAdd).Add(p.Cha.DefAdd).Float64()
	addMak := calc.NewNum(p.Pet.Feature.MakAdd).Add(p.Cha.MakAdd).Float64()
	addMdf := calc.NewNum(p.Pet.Feature.MdfAdd).Add(p.Cha.MdfAdd).Float64()
	addSpd := calc.NewNum(p.Pet.Feature.SpdAdd).Add(p.Cha.SpdAdd).Float64()
	addCrt := calc.NewNum(p.Pet.Feature.CrtAdd).Add(p.Cha.CrtAdd).Float64()
	addCriDmg := calc.NewNum(p.Pet.Feature.CriDmgAdd).Add(p.Cha.CriDmgAdd).Float64()

	// 努力值
	stvHp := calc.NewNum(float64(p.StvHp)).Div(4).Float64()
	stvAtk := calc.NewNum(float64(p.StvAtk)).Div(4).Float64()
	stvDef := calc.NewNum(float64(p.StvDef)).Div(4).Float64()
	stvMak := calc.NewNum(float64(p.StvMak)).Div(4).Float64()
	stvMdf := calc.NewNum(float64(p.StvMdf)).Div(4).Float64()
	stvSpd := calc.NewNum(float64(p.StvSpd)).Div(4).Float64()

	lv := float64(p.Lv)

	// 计算
	p.Hp = int(calc.NewNum(hp).Mul(2).Add(taleHp).Mul(lv).Div(100).Add(lv).Add(10).Mul(calc.NewNum(1).Add(addHp).Float64()).Add(stvHp).Int64())
	p.Atk = int(calc.NewNum(atk).Mul(2).Add(taleAtk).Mul(lv).Div(100).Add(5).Mul(calc.NewNum(1).Add(addAtk).Float64()).Add(stvAtk).Int64())
	p.Def = int(calc.NewNum(def).Mul(2).Add(taleDef).Mul(lv).Div(100).Add(5).Mul(calc.NewNum(1).Add(addDef).Float64()).Add(stvDef).Int64())
	p.Mak = int(calc.NewNum(mak).Mul(2).Add(taleMak).Mul(lv).Div(100).Add(5).Mul(calc.NewNum(1).Add(addMak).Float64()).Add(stvMak).Int64())
	p.Mdf = int(calc.NewNum(mdf).Mul(2).Add(taleMdf).Mul(lv).Div(100).Add(5).Mul(calc.NewNum(1).Add(addMdf).Float64()).Add(stvMdf).Int64())
	p.Spd = int(calc.NewNum(spd).Mul(2).Add(taleSpd).Mul(lv).Div(100).Add(5).Mul(calc.NewNum(1).Add(addSpd).Float64()).Add(stvSpd).Int64())
	p.Crt = calc.NewNum(crt).Add(addCrt).Float64()
	p.CriDmg = calc.NewNum(criDmg).Add(addCriDmg).Float64()
	return p
}

func (p *UserPet) View() string {
	p.Calc()
	view := fmt.Sprintf("%v %v\n", p.Pet.Feature.Emoji, p.Pet.Name)
	view += fmt.Sprintf("%v Lv.%v\n", p.EmojiLv(), p.Lv)

	var exp uint
	if p.Lv >= 100 {
		exp = 0
	} else {
		exp = p.Exp
	}
	view += fmt.Sprintf("%v %v/%v\n", emoji.Exp, exp, p.NextLvExp())

	view += fmt.Sprintf("%v 精力：%v ", emoji.HP, p.Hp)
	if p.TaleHp > 0 {
		view += fmt.Sprintf("%v %v ", emoji.Talent, p.TaleHp)
	}
	if p.StvHp > 0 {
		view += fmt.Sprintf("%v %v ", emoji.Strive, p.StvHp)
	}
	view = view[:len(view)-1] + "\n"

	view += fmt.Sprintf("%v 攻击：%v ", emoji.Atk, p.Atk)
	if p.TaleAtk > 0 {
		view += fmt.Sprintf("%v %v ", emoji.Talent, p.TaleAtk)
	}
	if p.StvAtk > 0 {
		view += fmt.Sprintf("%v %v ", emoji.Strive, p.StvAtk)
	}
	view = view[:len(view)-1] + "\n"

	view += fmt.Sprintf("%v 防御：%v ", emoji.Def, p.Def)
	if p.StvDef > 0 {
		view += fmt.Sprintf("%v %v ", emoji.Talent, p.TaleDef)
	}
	if p.StvDef > 0 {
		view += fmt.Sprintf("%v %v ", emoji.Strive, p.StvDef)
	}
	view = view[:len(view)-1] + "\n"

	view += fmt.Sprintf("%v 魔攻：%v ", emoji.Mak, p.Mak)
	if p.TaleMak > 0 {
		view += fmt.Sprintf("%v %v ", emoji.Talent, p.TaleMak)
	}
	if p.StvMak > 0 {
		view += fmt.Sprintf("%v %v ", emoji.Strive, p.StvMak)
	}
	view = view[:len(view)-1] + "\n"

	view += fmt.Sprintf("%v 魔抗：%v ", emoji.Mdf, p.Mdf)
	if p.TaleMdf > 0 {
		view += fmt.Sprintf("%v %v ", emoji.Talent, p.TaleMdf)
	}
	if p.StvMdf > 0 {
		view += fmt.Sprintf("%v %v ", emoji.Strive, p.StvMdf)
	}
	view = view[:len(view)-1] + "\n"

	view += fmt.Sprintf("%v 速度：%v ", emoji.Spd, p.Spd)
	if p.TaleSpd > 0 {
		view += fmt.Sprintf("%v %v ", emoji.Talent, p.TaleSpd)
	}
	if p.StvSpd > 0 {
		view += fmt.Sprintf("%v %v ", emoji.Strive, p.StvSpd)
	}
	view = view[:len(view)-1] + "\n"

	view += fmt.Sprintf("%v 暴击：%v%%\n", emoji.Crt, calc.NewNum(p.Crt).Mul(100).Float64())
	view += fmt.Sprintf("%v 爆伤：%v%%\n", emoji.CriDmg, calc.NewNum(p.CriDmg).Sub(1).Mul(100).Float64())
	view += fmt.Sprintf("%v %v\n%v\n", emoji.Character, p.Cha.Name, p.Cha.Effect)
	if p.Pet.EvoPetId != 0 {
		view += fmt.Sprintf("%v 下一进化：%v\n", p.Pet.EmojiEvo(), p.Pet.EvoPet.Name)
		view += fmt.Sprintf("%v 进化等级：%v\n", EmojiLv(p.Pet.EvoLv), p.Pet.EvoLv)
	}
	return view[:len(view)-1]
}
