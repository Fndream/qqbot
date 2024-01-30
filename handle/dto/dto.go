package dto

import "qqbot/handle/tb"

// Main

type Sign struct {
	Count         int
	Continuous    int
	RewardGold    int
	RewardDiamond int
}

type Query struct {
	*tb.UserInfo
	*tb.UserSign
}

type Bag struct {
	PageNo    int
	PageTotal int
	Items     []*tb.UserItem
}

// Good

type Shop struct {
	PageNo    int
	PageTotal int
	Goods     []*tb.Good
}

type BuyGood struct {
	Good       *tb.Good
	BuyCount   int
	SubGold    int
	SubDiamond int
}

// Game

type Fishing struct {
	BiteSecond int
	MaxSecond  int
}

type Draw struct {
	Type      int
	GoldType  int
	Gold      int
	Diamond   int
	Item      *tb.Item
	ItemCount int
}

type Gambling struct {
	Success bool
	Gold    int
	Diamond int
}

// Pet

type PetBag struct {
	PageNo    int
	PageTotal int
	Pets      []*tb.UserPet
}

// Use

type Use struct {
	UserItem *tb.UserItem
	UseCount int
	OldPet   *tb.UserPet
	NewPet   *tb.UserPet
}

type UseExpFruit struct {
	*Use
	AddExp int
	AddLv  int
}

type UseStvFruit struct {
	*Use
}

type UseForgetFruit struct {
	*Use
	ReturnStrive int
}

type UseTalentFruit struct {
	*Use
}

type UseEvoFruit struct {
	*Use
}

type UseCharacterFruit struct {
	*Use
}
