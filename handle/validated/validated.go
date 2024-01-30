package validated

func PageNo(i **int) {
	if *i == nil {
		*i = new(int)
		**i = 1
	}
	if **i < 1 {
		panic("🔖 无效的页码")
	}
}

func Count(i **int) {
	if *i == nil {
		*i = new(int)
		**i = 1
	}
	if **i < 1 {
		panic("🔖 非法的数量")
	}
}

func Gt0(i int) {
	if i < 1 {
		panic("🔖 非法的数量")
	}
}

func UintId(i int) {
	if i < 1 {
		panic("🔖 非法的ID")
	}
}
