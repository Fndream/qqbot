package view

import (
	"fmt"
	"qqbot/cmd"
	"qqbot/handle/dto"
	"qqbot/handle/service"
	"qqbot/handle/validated"
)

func Fishing(ctx *cmd.Context) (res *cmd.MsgView, err error) {
	var dto *dto.Fishing
	if dto, err = service.Fishing(ctx.Data.Author.ID); err != nil {
		return
	}

	msg := fmt.Sprintf("🎣 开始钓鱼\n预言家说在%v-%v秒内会有鱼咬钩，记得及时【拉竿】哦~", dto.BiteSecond, dto.MaxSecond)
	res = &cmd.MsgView{Msg: msg}
	return
}

func Draw(ctx *cmd.Context) (res *cmd.MsgView, err error) {
	var dto *dto.Draw
	if dto, err = service.Draw(ctx.Data.Author.ID); err != nil {
		return
	}

	msg := ""
	switch dto.Type {
	case service.DrawGold:
		switch dto.GoldType {
		case service.DrawGoldFish:
			msg += fmt.Sprintf("🐟 鱼上钩了\n💰 卖了%v枚金币\n", dto.Gold)
		case service.DrawGoldBigFish:
			msg += fmt.Sprintf("🐟 一条大鱼上钩了！\n💰 卖了%v枚金币\n", dto.Gold)
		case service.DrawGoldMiniFish:
			msg += fmt.Sprintf("🦐 emm...钓上来一只小虾米\n💰 卖了%v枚金币\n", dto.Gold)
		}
	case service.DrawDiamond:
		msg += "🐠 一条美丽的鱼跃出水面\n💎 将1颗钻石抛到了你的手里\n"
	case service.DrawGoldTreasure:
		msg += fmt.Sprintf("🐋 这是！深海的宝藏！！！\n💰 获得了%v枚金币！\n", dto.Gold)
	case service.DrawDiamondTreasure:
		msg += fmt.Sprintf("🐋 这是！深海的宝藏！！！\n💎 获得了%v颗钻石！！\n", dto.Diamond)
	case service.DrawGoldGod:
		msg += fmt.Sprintf("☀ God！\n💰 + %v\n", dto.Gold)
	case service.DrawDiamondGod:
		msg += fmt.Sprintf("☀ God！\n💎 + %v\n", dto.Diamond)
	case service.DrawExpFruit:
		msg += fmt.Sprintf("%v 哇！钓上来%v个%v！\n", dto.Item.Emoji, dto.ItemCount, dto.Item.Name)
	}
	msg = msg[:len(msg)-1]
	res = &cmd.MsgView{Msg: msg}
	return
}

func Gambling(ctx *cmd.Context, count int) (res *cmd.MsgView, err error) {
	validated.Gt0(count)

	var dto *dto.Gambling
	if dto, err = service.Gambling(ctx.Data.Author.ID, count); err != nil {
		return
	}

	msg := ""
	if dto.Success {
		msg += "🍱 聚宝成功\n"
	} else {
		msg += "🍱 聚宝失败\n"
	}
	if dto.Gold > 0 {
		if dto.Success {
			msg += fmt.Sprintf("💰 金币 + %v\n", dto.Gold)
		} else {
			msg += fmt.Sprintf("💰 金币 - %v\n", dto.Gold)
		}
	}
	if dto.Diamond > 0 {
		msg += fmt.Sprintf("💎 钻石 + %v\n", dto.Diamond)
	}

	msg = msg[:len(msg)-1]
	res = &cmd.MsgView{Msg: msg}
	return
}
