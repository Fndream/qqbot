package cmd

import (
	"github.com/tencent-connect/botgo/dto"
	"github.com/tencent-connect/botgo/dto/message"
	"github.com/tencent-connect/botgo/log"
	"qqbot/cmd/cache"
	"strconv"
	"time"
)

type MsgView struct {
	Msg   string
	Image string
	NotAt bool
	MDID  int
}

var imageCache = cache.New(10*time.Minute, 5*time.Minute)

func SendReply(ctx *Context, msg *MsgView) {
	imgUrl := ""
	haveCache := false
	if msg.Image != "" {
		cv, ok := imageCache.Get(msg.Image)
		haveCache = ok
		if ok {
			imgUrl = cv.(string)
		} else {
			imgUrl = msg.Image
		}
	}

	var msgCreate *dto.MessageToCreate

	if msg.MDID == 0 {
		msgCreate = &dto.MessageToCreate{
			Image:   imgUrl,
			Content: msg.Msg,
			MsgID:   ctx.Data.ID,
		}
	} else {
		msgCreate = &dto.MessageToCreate{
			Markdown: &dto.Markdown{
				TemplateID: msg.MDID,
				Params:     []*dto.MarkdownParams{{Key: "content", Values: []string{msg.Msg}}},
			},
			MsgID: ctx.Data.ID,
		}
		if imgUrl != "" {
			msgCreate.Markdown.Params = append(msgCreate.Markdown.Params, &dto.MarkdownParams{Key: "image", Values: []string{imgUrl}})
		}
	}

	if ctx.Data.DirectMessage {
		directMessage := dto.DirectMessage{
			GuildID:    ctx.Data.GuildID,
			ChannelID:  ctx.Data.ChannelID,
			CreateTime: strconv.FormatInt(time.Now().Unix(), 10),
		}
		_, err := ctx.Api.PostDirectMessage(ctx, &directMessage, msgCreate)
		if err != nil {
			log.Error(err)
			return
		}
	} else {
		if !msg.NotAt {
			msgCreate.Content = message.MentionUser(ctx.Data.Author.ID) + "\n" + msgCreate.Content
		}
		rep, err := ctx.Api.PostMessage(ctx, ctx.Data.ChannelID, msgCreate)
		if err != nil {
			log.Error(err)
			return
		}
		// 只能获取子频道中的消息
		if imgUrl != "" && !haveCache {
			m, err := ctx.Api.Message(ctx, rep.ChannelID, rep.ID)
			if err != nil {
				log.Error(err)
				return
			}
			imageCache.Set(msg.Image, "https://"+m.Attachments[0].URL)
		}
	}
}

func SendReplyS(ctx *Context, msg string) {
	SendReply(ctx, &MsgView{Msg: msg})
}
