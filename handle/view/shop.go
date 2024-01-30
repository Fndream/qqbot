package view

import (
	"fmt"
	"qqbot/cmd"
	"qqbot/handle/dto"
	"qqbot/handle/service"
	"qqbot/handle/validated"
)

// Shop 商店
func Shop(ctx *cmd.Context, pageNo *int) (res *cmd.MsgView, err error) {
	validated.PageNo(&pageNo)

	var dto *dto.Shop
	if dto, err = service.Shop(*pageNo); err != nil {
		return
	}

	msg := fmt.Sprintf("🛍 %v/%v页\n", dto.PageNo, dto.PageTotal)
	for _, good := range dto.Goods {
		msg += fmt.Sprintf("%v %v\n", good.Item.Emoji, good.Item.Name)
		if good.BuyGold > 0 || good.BuyDiamond > 0 {
			msg += "👛 价格："
			if good.BuyGold > 0 {
				msg += fmt.Sprintf("💰%v金币  ", good.BuyGold)
			}
			if good.BuyDiamond > 0 {
				msg += fmt.Sprintf("💎%v钻石  ", good.BuyDiamond)
			}
			msg = msg[:len(msg)-2] + "\n"
		}
	}
	msg = msg[:len(msg)-1]
	res = &cmd.MsgView{Msg: msg}
	return
}

// Trading 交易所
func Trading(ctx *cmd.Context, pageNo *int) (res *cmd.MsgView, err error) {
	validated.PageNo(&pageNo)

	var dto *dto.Shop
	if dto, err = service.Trading(*pageNo); err != nil {
		return
	}

	msg := fmt.Sprintf("🥝 %v/%v页\n", dto.PageNo, dto.PageTotal)
	for _, good := range dto.Goods {
		msg += fmt.Sprintf("%v %v\n", good.Item.Emoji, good.Item.Name)
		if good.SellGold > 0 || good.SellDiamond > 0 {
			msg += "🥢 出售可得："
			if good.SellGold > 0 {
				msg += fmt.Sprintf("💰%v金币 ", good.SellGold)
			}
			if good.SellDiamond > 0 {
				msg += fmt.Sprintf("💎%v钻石 ", good.SellDiamond)
			}
			msg = msg[:len(msg)-1] + "\n"
		}
	}
	msg = msg[:len(msg)-1]
	res = &cmd.MsgView{Msg: msg}
	return
}

func BuyGood(ctx *cmd.Context, buyName string, buyCount *int) (res *cmd.MsgView, err error) {
	validated.Count(&buyCount)

	var dto *dto.BuyGood
	if dto, err = service.BuyGood(ctx.Data.Author.ID, buyName, *buyCount); err != nil {
		return
	}

	var msg string
	msg += fmt.Sprintf("%v %v × %v 已放入背包\n", dto.Good.Item.Emoji, dto.Good.Item.Name, dto.BuyCount)
	if dto.SubGold > 0 {
		msg += fmt.Sprintf("💰 金币 - %v\n", dto.SubGold)
	}
	if dto.SubDiamond > 0 {
		msg += fmt.Sprintf("💎 钻石 - %v\n", dto.SubDiamond)
	}
	msg = msg[:len(msg)-1]
	res = &cmd.MsgView{Msg: msg}
	return
}
