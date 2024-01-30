package util

import (
	"github.com/bwmarrin/snowflake"
	"math/rand"
)

var node *snowflake.Node

func init() {
	var err error
	node, err = snowflake.NewNode(1)
	if err != nil {
		panic(err)
	}
}

func NewId() string {
	return node.Generate().String()
}

func RandBool() bool {
	return rand.Intn(2) == 0
}

func RandIntn(n int) int {
	return rand.Intn(n)
}

func RandInt(min int, max int) int {
	return rand.Intn(max-min+1) + min
}

func RandString(s []string) string {
	return s[RandInt(0, len(s)-1)]
}

func Lottery(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	sum := 0
	for _, num := range nums {
		sum += num
	}

	r := rand.Intn(sum)

	for i, num := range nums {
		r -= num
		if r < 0 {
			return i
		}
	}
	return -1
}

func RandCharacterId() uint {
	return uint(RandInt(1, 56))
}
