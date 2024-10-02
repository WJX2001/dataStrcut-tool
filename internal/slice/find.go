package slice

// Find 查找元素
// 如果没有找到，第二个返回值返回 false

func Find[T any](src []T, match matchFunc[T]) (T, bool) {
	for _, v := range src {
		if match(v) {
			return v, true
		}
	}

	var t T
	return t, false
}

// FindAll 查找所有符合条件的元素
// 永远不会返回 nil

func FindAll[T any](src []T, match matchFunc[T]) []T {
	// 我们认为符合条件元素应该是少数
	// 所以会除以8
	// 也就是触发扩容的情况下，最多三次就会和原来的容量一样
	res := make([]T, 0, len(src)>>3+1)
	for _, val := range src {
		if match(val) {
			res = append(res, val)
		}
	}
	return res
}
