package main

import (
	"context"
	"fmt"
	bot "github.com/tencent-connect/botgo"
	"github.com/tencent-connect/botgo/dto"
	"github.com/tencent-connect/botgo/event"
	"github.com/tencent-connect/botgo/log"
	"github.com/tencent-connect/botgo/token"
	"github.com/tencent-connect/botgo/websocket"
	"path"
	"qqbot/cmd"
	"qqbot/handle/view"
	"qqbot/pkg/db"
	"runtime"
	"time"
)

func init() {
	db.SetUp()
}

func main() {
	ctx := context.Background()
	tok := token.New(token.TypeBot)
	if err := tok.LoadFromConfig(getConfigPath("config.yaml")); err != nil {
		log.Error(err)
		return
	}

	api := bot.NewOpenAPI(tok).WithTimeout(3 * time.Second)
	cmd.SetApi(api)

	wsInfo, err := api.WS(ctx, nil, "")
	if err != nil {
		log.Error(err)
		return
	}

	regCmd()

	intent := websocket.RegisterHandlers(
		// at 机器人事件，目前是在这个事件处理中有逻辑，会回消息，其他的回调处理都只把数据打印出来，不做任何处理
		ATMessageEventHandler(),
		// 如果想要捕获到连接成功的事件，可以实现这个回调
		ReadyHandler(),
		// 连接关闭回调
		ErrorNotifyHandler(),
		//// 频道事件
		//GuildEventHandler(),
		//// 成员事件
		//MemberEventHandler(),
		//// 子频道事件
		//ChannelEventHandler(),
		//// 私信，目前只有私域才能够收到这个，如果你的机器人不是私域机器人，会导致连接报错，那么启动 example 就需要注释掉这个回调
		DirectMessageHandler(),
		//// 频道消息，只有私域才能够收到这个，如果你的机器人不是私域机器人，会导致连接报错，那么启动 example 就需要注释掉这个回调
		CreateMessageHandler(),
		//// 互动事件
		//InteractionHandler(),
		//// 发帖事件
		//ThreadEventHandler(),
	)

	// 指定需要启动的分片数为 2 的话可以手动修改 wsInfo
	if err = bot.NewSessionManager().Start(wsInfo, tok, &intent); err != nil {
		log.Error(err)
	}
}

func regCmd() {
	// main
	cmd.Register(&cmd.Config{
		ID:    "sign",
		Name:  "签到",
		Emoji: "📅",
	}, view.Sign)
	cmd.Register(&cmd.Config{
		ID:    "query",
		Name:  "查询",
		Emoji: "🔍",
	}, view.Query)
	cmd.Register(&cmd.Config{
		ID:    "bag",
		Name:  "背包",
		Alias: []string{"我的背包"},
		Emoji: "🎒",
	}, view.Bag)
	cmd.Register(&cmd.Config{
		ID:    "use-item",
		Name:  "使用物品",
		Alias: []string{"使用"},
		Emoji: "🍰",
	}, view.UseItem)

	// Use
	cmd.Register(&cmd.Config{
		ID:      "use-exp-fruit",
		Emoji:   "🍏",
		Private: true,
	}, view.UseExpFruitByPetSerial, view.UseExpFruitByPetName)
	cmd.Register(&cmd.Config{
		ID:      "use-strive-fruit",
		Emoji:   "🌸",
		Private: true,
	}, view.UseStvFruitByPetSerial, view.UseStvFruitByPetName)
	cmd.Register(&cmd.Config{
		ID:      "use-forget-fruit",
		Emoji:   "💙",
		Private: true,
	}, view.UseForgetFruitByPetSerial, view.UseForgetFruitByPetName)
	cmd.Register(&cmd.Config{
		ID:      "use-evolve-fruit",
		Emoji:   "🍊",
		Private: true,
	}, view.UseEvoFruitByPetSerial, view.UseEvoFruitByPetName)
	cmd.Register(&cmd.Config{
		ID:      "use-talent-fruit",
		Emoji:   "🍬",
		Private: true,
	}, view.UseTalentFruitByPetSerial, view.UseTalentFruitByPetName)
	cmd.Register(&cmd.Config{
		ID:      "use-character-fruit",
		Emoji:   "💟",
		Private: true,
	}, view.UseCharacterFruitByPetSerial, view.UseCharacterFruitByPetName)

	// shop
	cmd.Register(&cmd.Config{
		ID:    "shop",
		Name:  "商城",
		Alias: []string{"商店", "商场"},
		Emoji: "🛍",
	}, view.Shop)
	cmd.Register(&cmd.Config{
		ID:    "trading",
		Name:  "交易所",
		Alias: []string{"交易"},
		Emoji: "🥝",
	}, view.Trading)
	cmd.Register(&cmd.Config{
		ID:    "buy-good",
		Name:  "购买商品",
		Alias: []string{"购买物品", "购买"},
		Emoji: "🛍",
	}, view.BuyGood)

	// game
	cmd.Register(&cmd.Config{
		ID:    "fishing",
		Name:  "钓鱼",
		Emoji: "🎣",
	}, view.Fishing)
	cmd.Register(&cmd.Config{
		ID:    "draw",
		Name:  "拉竿",
		Alias: []string{"收竿", "拉杆"},
		Emoji: "🎣",
	}, view.Draw)
	cmd.Register(&cmd.Config{
		ID:    "gambling",
		Name:  "聚宝",
		Emoji: "🍱",
	}, view.Gambling)

	// pet
	cmd.Register(&cmd.Config{
		ID:    "receive-pet",
		Name:  "领取宠物",
		Emoji: "🐱",
	}, view.ReceivePet)
	cmd.Register(&cmd.Config{
		ID:    "pet-bag",
		Name:  "宠物背包",
		Alias: []string{"宠物", "我的宠物"},
		Emoji: "🐱",
	}, view.PetBag)
	cmd.Register(&cmd.Config{
		ID:    "query-pet-handbook",
		Name:  "查询宠物图鉴",
		Alias: []string{"查询图鉴", "查询宠物", "查宠"},
		Emoji: "📓",
	}, view.QueryPetById, view.QueryPetByName)
	cmd.Register(&cmd.Config{
		ID:    "query-userpet",
		Name:  "查看宠物",
		Alias: []string{"查看"},
		Emoji: "🐱",
	}, view.QueryUserPetById, view.QueryUserPetByName)
}

// ATMessageEventHandler 实现处理 at 消息的回调
func ATMessageEventHandler() event.ATMessageEventHandler {
	return func(event *dto.WSPayload, data *dto.WSATMessageData) error {
		//cmd.Process((*dto.Message)(data))
		return nil
	}
}

// DirectMessageHandler 处理私信事件
func DirectMessageHandler() event.DirectMessageEventHandler {
	return func(event *dto.WSPayload, data *dto.WSDirectMessageData) error {
		cmd.Process((*dto.Message)(data))
		return nil
	}
}

// ReadyHandler 自定义 ReadyHandler 感知连接成功事件
func ReadyHandler() event.ReadyHandler {
	return func(event *dto.WSPayload, data *dto.WSReadyData) {
		log.Info("ready event receive: ", data)
	}
}

func ErrorNotifyHandler() event.ErrorNotifyHandler {
	return func(err error) {
		log.Info("error notify receive: ", err)
	}
}

// GuildEventHandler 处理频道事件
func GuildEventHandler() event.GuildEventHandler {
	return func(event *dto.WSPayload, data *dto.WSGuildData) error {
		fmt.Println(data)
		return nil
	}
}

// ChannelEventHandler 处理子频道事件
func ChannelEventHandler() event.ChannelEventHandler {
	return func(event *dto.WSPayload, data *dto.WSChannelData) error {
		fmt.Println(data)
		return nil
	}
}

// MemberEventHandler 处理成员变更事件
func MemberEventHandler() event.GuildMemberEventHandler {
	return func(event *dto.WSPayload, data *dto.WSGuildMemberData) error {
		fmt.Println(data)
		return nil
	}
}

// CreateMessageHandler 处理消息事件
func CreateMessageHandler() event.MessageEventHandler {
	return func(event *dto.WSPayload, data *dto.WSMessageData) error {
		cmd.Process((*dto.Message)(data))
		return nil
	}
}

// InteractionHandler 处理内联交互事件
//func InteractionHandler() event.InteractionEventHandler {
//	return func(event *dto.WSPayload, data *dto.WSInteractionData) error {
//		fmt.Println(data)
//		return processor.ProcessInlineSearch(data)
//	}
//}

func getConfigPath(name string) string {
	_, filename, _, ok := runtime.Caller(1)
	if ok {
		return fmt.Sprintf("%s/%s", path.Dir(filename), name)
	}
	return ""
}
