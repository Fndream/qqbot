package view

import (
	"fmt"
	"qqbot/cmd"
	"qqbot/handle/dto"
	"qqbot/handle/service"
	"qqbot/handle/validated"
	"time"
)

// Sign 签到
func Sign(ctx *cmd.Context) (res *cmd.MsgView, err error) {
	var dto *dto.Sign
	if dto, err = service.Sign(ctx.Data.Author.ID); err != nil {
		return
	}

	msg := "🎉 签到成功！\n"
	msg += fmt.Sprintf("☀ 已累计签到%v天\n", dto.Count)
	msg += fmt.Sprintf("🌙 已连续签到%v天\n", dto.Continuous)
	msg += fmt.Sprintf("🎁 获得奖励\n")
	if dto.RewardGold > 0 {
		msg += fmt.Sprintf("💰 金币 + %v\n", dto.RewardGold)
	}
	if dto.RewardDiamond > 0 {
		msg += fmt.Sprintf("💎 钻石 + %v\n", dto.RewardDiamond)
	}
	msg = msg[:len(msg)-1]
	res = &cmd.MsgView{Msg: msg}
	return
}

// Query 查询
func Query(ctx *cmd.Context) (res *cmd.MsgView, err error) {
	var dto *dto.Query
	if dto, err = service.Query(ctx.Data.Author.ID); err != nil {
		return
	}

	now := time.Now().Format(time.DateOnly)
	sd := dto.UpdatedAt.Format(time.DateOnly)
	msg := ""
	if now != sd {
		msg += fmt.Sprintf("📅 今日未签到！\n")
	}
	if !dto.UpdatedAt.IsZero() {
		msg += fmt.Sprintf("📅 签到日期：%v\n", sd)
	}
	msg += fmt.Sprintf("☀ 累计签到：%v天\n", dto.Count)
	msg += fmt.Sprintf("🌙 连续签到：%v天\n", dto.Continuous)
	msg += fmt.Sprintf("💰 金币数量：%v枚\n", dto.Gold)
	msg += fmt.Sprintf("💎 钻石数量：%v颗\n", dto.Diamond)
	if dto.Strive > 0 {
		msg += fmt.Sprintf("🔥 努力点数: %v点\n", dto.Strive)
	}
	msg = msg[:len(msg)-1]
	res = &cmd.MsgView{Msg: msg}
	return
}

// Bag 背包
func Bag(ctx *cmd.Context, pageNo *int) (res *cmd.MsgView, err error) {
	validated.PageNo(&pageNo)

	var dto *dto.Bag
	if dto, err = service.Bag(ctx.Data.Author.ID, *pageNo); err != nil {
		return
	}

	msg := fmt.Sprintf("🎒 %v/%v页\n", dto.PageNo, dto.PageTotal)
	for _, item := range dto.Items {
		msg += fmt.Sprintf("%v %v %v\n", item.Item.Emoji, item.Item.Name, item.Count)
	}
	msg = msg[:len(msg)-1]
	res = &cmd.MsgView{Msg: msg}
	return
}
