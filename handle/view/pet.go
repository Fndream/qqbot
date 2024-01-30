package view

import (
	"fmt"
	"qqbot/cmd"
	"qqbot/handle/dto"
	"qqbot/handle/service"
	"qqbot/handle/tb"
	"qqbot/handle/validated"
)

func ReceivePet(ctx *cmd.Context) (res *cmd.MsgView, err error) {
	var pet *tb.Pet
	if pet, err = service.ReceivePet(ctx.Data.Author.ID); err != nil {
		return
	}

	res = &cmd.MsgView{Msg: fmt.Sprintf("%v %v已放入宠物背包", pet.Feature.Emoji, pet.Name)}
	return
}

func PetBag(ctx *cmd.Context, pageNo *int) (res *cmd.MsgView, err error) {
	validated.PageNo(&pageNo)

	var dto *dto.PetBag
	if dto, err = service.PetBag(ctx.Data.Author.ID, *pageNo); err != nil {
		return
	}

	msg := fmt.Sprintf("🐱 %v/%v页\n", dto.PageNo, dto.PageTotal)
	for _, pet := range dto.Pets {
		emojiStatus := pet.EmojiStatus()
		if emojiStatus != "" {
			emojiStatus = " " + emojiStatus
		}
		msg += fmt.Sprintf("%v ID.%v %v Lv.%v%v %v %v\n", pet.EmojiLv(), pet.Serial, pet.Pet.Name, pet.Lv, emojiStatus, pet.Cha.Emoji, pet.Cha.Name)
	}
	msg = msg[:len(msg)-1]
	res = &cmd.MsgView{Msg: msg}
	return
}

func QueryPetById(ctx *cmd.Context, id int) (res *cmd.MsgView, err error) {
	validated.UintId(id)
	var pet *tb.Pet
	if pet, err = service.QueryPetById(uint(id)); err != nil {
		return
	}
	res = &cmd.MsgView{Msg: pet.View(), Image: "https://gitee.com/fndream/mg-rep/raw/master/mg/001.png", MDID: 1705978299}
	return
}

func QueryPetByName(ctx *cmd.Context, name string) (res *cmd.MsgView, err error) {
	var pet *tb.Pet
	if pet, err = service.QueryPetByName(name); err != nil {
		return
	}
	res = &cmd.MsgView{Msg: pet.View()}
	return
}

func QueryUserPetById(ctx *cmd.Context, id int) (res *cmd.MsgView, err error) {
	validated.UintId(id)
	var pet *tb.UserPet
	if pet, err = service.QueryUserPetBySerial(ctx.Data.Author.ID, uint(id)); err != nil {
		return
	}
	res = &cmd.MsgView{Msg: pet.View()}
	return
}

func QueryUserPetByName(ctx *cmd.Context, name string) (res *cmd.MsgView, err error) {
	var pet *tb.UserPet
	if pet, err = service.QueryUserPetByName(ctx.Data.Author.ID, name); err != nil {
		return
	}
	res = &cmd.MsgView{Msg: pet.View()}
	return
}
