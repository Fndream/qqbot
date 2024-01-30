package view

import (
	"errors"
	"fmt"
	"qqbot/cmd"
	"qqbot/handle/dto"
	"qqbot/handle/service"
	"qqbot/handle/tb"
	"qqbot/handle/tb/emoji"
	"qqbot/handle/validated"
)

func UseItem(ctx *cmd.Context, name string, count *int, args ...interface{}) (res *cmd.MsgView, err error) {
	validated.Count(&count)

	// 判断用户物品是否足够
	userItem, err := service.GetUserItem(ctx.Data.Author.ID, name)
	if err != nil {
		return
	}
	if userItem.ID == "" {
		return nil, errors.New("🎒 背包中没有要使用的物品")
	}
	if userItem.Count < *count {
		return nil, fmt.Errorf("%v 背包中没有足够的%v", userItem.Item.Emoji, userItem.Item.Name)
	}

	// 获取对应的使用指令
	var command *cmd.Command
	if userItem.ItemId >= 10001 && userItem.ItemId <= 10011 {
		command, _ = cmd.GetPrivateCommand("use-exp-fruit")
	} else if userItem.ItemId >= 10101 && userItem.ItemId <= 10106 {
		command, _ = cmd.GetPrivateCommand("use-strive-fruit")
	} else if userItem.ItemId == 10107 {
		command, _ = cmd.GetPrivateCommand("use-forget-fruit")
	} else if userItem.ItemId == 10201 {
		command, _ = cmd.GetPrivateCommand("use-talent-fruit")
	} else if userItem.ItemId == 10202 {
		command, _ = cmd.GetPrivateCommand("use-evolve-fruit")
	} else if userItem.ItemId == 10203 {
		command, _ = cmd.GetPrivateCommand("use-character-fruit")
	}
	if command == nil {
		return nil, fmt.Errorf("%v %v无法通过该指令使用，或本就不可被使用", userItem.Item.Emoji, userItem.Item.Name)
	}
	ctx.Cmd = command
	ctx.Args = []interface{}{userItem, *count}
	ctx.Args = append(ctx.Args, args...)
	cmd.Run(ctx)
	return
}

// 经验果

func UseExpFruitByPetSerial(ctx *cmd.Context, userItem *tb.UserItem, count *int, serial int) (res *cmd.MsgView, err error) {
	var dto *dto.UseExpFruit
	if dto, err = service.UseExpFruitByPetSerial(ctx.Data.Author.ID, userItem, *count, uint(serial)); err != nil {
		return
	}
	res = &cmd.MsgView{Msg: useExpFruitView(dto)}
	return
}

func UseExpFruitByPetName(ctx *cmd.Context, userItem *tb.UserItem, count *int, petName string) (res *cmd.MsgView, err error) {
	var dto *dto.UseExpFruit
	if dto, err = service.UseExpFruitByPetName(ctx.Data.Author.ID, userItem, *count, petName); err != nil {
		return
	}
	res = &cmd.MsgView{Msg: useExpFruitView(dto)}
	return
}

func useExpFruitView(dto *dto.UseExpFruit) string {
	msg := fmt.Sprintf("%v %v吃下了%v个%v\n", dto.UserItem.Item.Emoji, dto.OldPet.Pet.Name, dto.UseCount, dto.UserItem.Item.Name)
	msg += fmt.Sprintf("🔵 Exp + %v\n", dto.AddExp)
	if dto.AddLv != 0 {
		msg += LevelUpView(dto.OldPet, dto.NewPet)
	}
	return msg[:len(msg)-1]
}

// 努力果

func UseStvFruitByPetSerial(ctx *cmd.Context, userItem *tb.UserItem, count *int, serial int) (res *cmd.MsgView, err error) {
	var dto *dto.UseStvFruit
	if dto, err = service.UseStvFruitByPetSerial(ctx.Data.Author.ID, userItem, *count, uint(serial)); err != nil {
		return
	}
	res = &cmd.MsgView{Msg: useStvFruitView(dto)}
	return
}

func UseStvFruitByPetName(ctx *cmd.Context, userItem *tb.UserItem, count *int, petName string) (res *cmd.MsgView, err error) {
	var dto *dto.UseStvFruit
	if dto, err = service.UseStvFruitByPetName(ctx.Data.Author.ID, userItem, *count, petName); err != nil {
		return
	}
	res = &cmd.MsgView{Msg: useStvFruitView(dto)}
	return
}

func useStvFruitView(dto *dto.UseStvFruit) string {
	dto.OldPet.Calc()
	dto.NewPet.Calc()
	msg := fmt.Sprintf("%v %v吃下了%v个%v\n", dto.UserItem.Item.Emoji, dto.OldPet.Pet.Name, dto.UseCount, dto.UserItem.Item.Name)
	if dto.NewPet.StvHp != dto.OldPet.StvHp {
		msg += fmt.Sprintf("🔥 精力努力值 + %v\n", dto.NewPet.StvHp-dto.OldPet.StvHp)
	}
	if dto.NewPet.StvAtk != dto.OldPet.StvAtk {
		msg += fmt.Sprintf("🔥 攻击努力值 + %v\n", dto.NewPet.StvAtk-dto.OldPet.StvAtk)
	}
	if dto.NewPet.StvDef != dto.OldPet.StvDef {
		msg += fmt.Sprintf("🔥 防御努力值 + %v\n", dto.NewPet.StvDef-dto.OldPet.StvDef)
	}
	if dto.NewPet.StvMak != dto.OldPet.StvMak {
		msg += fmt.Sprintf("🔥 魔攻努力值 + %v\n", dto.NewPet.StvMak-dto.OldPet.StvMak)
	}
	if dto.NewPet.StvMdf != dto.OldPet.StvMdf {
		msg += fmt.Sprintf("🔥 魔抗努力值 + %v\n", dto.NewPet.StvMdf-dto.OldPet.StvMdf)
	}
	if dto.NewPet.StvSpd != dto.OldPet.StvSpd {
		msg += fmt.Sprintf("🔥 速度努力值 + %v\n", dto.NewPet.StvSpd-dto.OldPet.StvSpd)
	}
	if dto.NewPet.Hp != dto.OldPet.Hp {
		msg += fmt.Sprintf("%v 精力 + %v\n", emoji.HP, dto.NewPet.Hp-dto.OldPet.Hp)
	}
	if dto.NewPet.Atk != dto.OldPet.Atk {
		msg += fmt.Sprintf("%v 攻击 + %v\n", emoji.Atk, dto.NewPet.Atk-dto.OldPet.Atk)
	}
	if dto.NewPet.Def != dto.OldPet.Def {
		msg += fmt.Sprintf("%v 防御 + %v\n", emoji.Def, dto.NewPet.Def-dto.OldPet.Def)
	}
	if dto.NewPet.Mak != dto.OldPet.Mak {
		msg += fmt.Sprintf("%v 魔攻 + %v\n", emoji.Mak, dto.NewPet.Mak-dto.OldPet.Mak)
	}
	if dto.NewPet.Mdf != dto.OldPet.Mdf {
		msg += fmt.Sprintf("%v 魔抗 + %v\n", emoji.Mdf, dto.NewPet.Mdf-dto.OldPet.Mdf)
	}
	if dto.NewPet.Spd != dto.OldPet.Spd {
		msg += fmt.Sprintf("%v 速度 + %v\n", emoji.Spd, dto.NewPet.Spd-dto.OldPet.Spd)
	}
	return msg[:len(msg)-1]
}

// 遗忘之果

func UseForgetFruitByPetSerial(ctx *cmd.Context, userItem *tb.UserItem, count *int, serial int) (res *cmd.MsgView, err error) {
	if *count > 1 {
		return nil, fmt.Errorf("%v %v每次最多只能使用1个", userItem.Item.Emoji, userItem.Item.Name)
	}
	var dto *dto.UseForgetFruit
	if dto, err = service.UseForgetFruitByPetSerial(ctx.Data.Author.ID, userItem, uint(serial)); err != nil {
		return
	}
	res = &cmd.MsgView{Msg: useForgetFruitView(dto)}
	return
}

func UseForgetFruitByPetName(ctx *cmd.Context, userItem *tb.UserItem, count *int, petName string) (res *cmd.MsgView, err error) {
	if *count > 1 {
		return nil, fmt.Errorf("%v %v每次最多只能使用1个", userItem.Item.Emoji, userItem.Item.Name)
	}
	var dto *dto.UseForgetFruit
	if dto, err = service.UseForgetFruitByPetName(ctx.Data.Author.ID, userItem, petName); err != nil {
		return
	}
	res = &cmd.MsgView{Msg: useForgetFruitView(dto)}
	return
}

func useForgetFruitView(dto *dto.UseForgetFruit) string {
	msg := fmt.Sprintf("%v %v吃掉了1个%v，遗忘了所有努力\n", dto.UserItem.Item.Emoji, dto.OldPet.Pet.Name, dto.UserItem.Item.Name)
	msg += fmt.Sprintf("🔥 获得%v努力点", dto.ReturnStrive)
	return msg
}

// 果儿糖

func UseTalentFruitByPetSerial(ctx *cmd.Context, userItem *tb.UserItem, count *int, serial int) (res *cmd.MsgView, err error) {
	if *count > 1 {
		return nil, fmt.Errorf("%v %v每次最多只能使用1个", userItem.Item.Emoji, userItem.Item.Name)
	}
	var dto *dto.UseTalentFruit
	if dto, err = service.UseTalentFruitByPetSerial(ctx.Data.Author.ID, userItem, uint(serial)); err != nil {
		return
	}
	res = &cmd.MsgView{Msg: useTalentFruitView(dto)}
	return
}

func UseTalentFruitByPetName(ctx *cmd.Context, userItem *tb.UserItem, count *int, petName string) (res *cmd.MsgView, err error) {
	if *count > 1 {
		return nil, fmt.Errorf("%v %v每次最多只能使用1个", userItem.Item.Emoji, userItem.Item.Name)
	}
	var dto *dto.UseTalentFruit
	if dto, err = service.UseTalentFruitByPetName(ctx.Data.Author.ID, userItem, petName); err != nil {
		return
	}
	res = &cmd.MsgView{Msg: useTalentFruitView(dto)}
	return
}

func useTalentFruitView(dto *dto.UseTalentFruit) string {
	msg := fmt.Sprintf("%v %v吃掉了1个%v，天赋得到洗炼\n", dto.UserItem.Item.Emoji, dto.OldPet.Pet.Name, dto.UserItem.Item.Name)
	msg += fmt.Sprintf("%v 精力：%v -> %v\n", emoji.HP, dto.OldPet.TaleHp, dto.NewPet.TaleHp)
	msg += fmt.Sprintf("%v 攻击：%v -> %v\n", emoji.Atk, dto.OldPet.TaleAtk, dto.NewPet.TaleAtk)
	msg += fmt.Sprintf("%v 防御：%v -> %v\n", emoji.Def, dto.OldPet.TaleDef, dto.NewPet.TaleDef)
	msg += fmt.Sprintf("%v 魔攻：%v -> %v\n", emoji.Mak, dto.OldPet.TaleMak, dto.NewPet.TaleMak)
	msg += fmt.Sprintf("%v 魔抗：%v -> %v\n", emoji.Mdf, dto.OldPet.TaleMdf, dto.NewPet.TaleMdf)
	msg += fmt.Sprintf("%v 速度：%v -> %v", emoji.Spd, dto.OldPet.TaleSpd, dto.NewPet.TaleSpd)
	return msg
}

// 进化果

func UseEvoFruitByPetSerial(ctx *cmd.Context, userItem *tb.UserItem, count *int, serial int) (res *cmd.MsgView, err error) {
	if *count > 1 {
		return nil, fmt.Errorf("%v %v每次最多只能使用1个", userItem.Item.Emoji, userItem.Item.Name)
	}
	var dto *dto.UseEvoFruit
	if dto, err = service.UseEvoFruitByPetSerial(ctx.Data.Author.ID, userItem, uint(serial)); err != nil {
		return
	}
	res = &cmd.MsgView{Msg: useEvoFruitView(dto)}
	return
}

func UseEvoFruitByPetName(ctx *cmd.Context, userItem *tb.UserItem, count *int, petName string) (res *cmd.MsgView, err error) {
	if *count > 1 {
		return nil, fmt.Errorf("%v %v每次最多只能使用1个", userItem.Item.Emoji, userItem.Item.Name)
	}
	var dto *dto.UseEvoFruit
	if dto, err = service.UseEvoFruitByPetName(ctx.Data.Author.ID, userItem, petName); err != nil {
		return
	}
	res = &cmd.MsgView{Msg: useEvoFruitView(dto)}
	return
}

func useEvoFruitView(dto *dto.UseEvoFruit) string {
	msg := fmt.Sprintf("%v %v吃掉了1个%v，进化为%v\n", dto.UserItem.Item.Emoji, dto.OldPet.Pet.Name, dto.UserItem.Item.Name, dto.NewPet.Pet.Name)
	msg += PetBattleAttrCompareView(dto.OldPet, dto.NewPet)
	return msg[:len(msg)-1]
}

// 归初果

func UseCharacterFruitByPetSerial(ctx *cmd.Context, userItem *tb.UserItem, count *int, serial int) (res *cmd.MsgView, err error) {
	if *count > 1 {
		return nil, fmt.Errorf("%v %v每次最多只能使用1个", userItem.Item.Emoji, userItem.Item.Name)
	}
	var dto *dto.UseCharacterFruit
	if dto, err = service.UseCharacterFruitByPetSerial(ctx.Data.Author.ID, userItem, uint(serial)); err != nil {
		return
	}
	res = &cmd.MsgView{Msg: useCharacterFruitView(dto)}
	return
}

func UseCharacterFruitByPetName(ctx *cmd.Context, userItem *tb.UserItem, count *int, petName string) (res *cmd.MsgView, err error) {
	if *count > 1 {
		return nil, fmt.Errorf("%v %v每次最多只能使用1个", userItem.Item.Emoji, userItem.Item.Name)
	}
	var dto *dto.UseCharacterFruit
	if dto, err = service.UseCharacterFruitByPetName(ctx.Data.Author.ID, userItem, petName); err != nil {
		return
	}
	res = &cmd.MsgView{Msg: useCharacterFruitView(dto)}
	return
}

func useCharacterFruitView(dto *dto.UseCharacterFruit) string {
	msg := fmt.Sprintf("%v %v吃掉了1个%v，性格发生了转变\n", dto.UserItem.Item.Emoji, dto.OldPet.Pet.Name, dto.UserItem.Item.Name)
	msg += fmt.Sprintf("%v %v\n%v", emoji.Character, dto.NewPet.Cha.Name, dto.NewPet.Cha.Effect)
	return msg
}
