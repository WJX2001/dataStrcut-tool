package slice

// DiffSet 差集，只支持 comparable类型
// 已去重
// 并且返回的顺序是不确定的

func DiffSet[T comparable](src, dst []T) []T {
	srcMap := toMap[T](src)
	for _, val := range dst {
		delete(srcMap, val)
	}

	var ret = make([]T, 0, len(srcMap))
	for key := range srcMap {
		ret = append(ret, key)
	}
	return ret
}

// DiffSetFunc 差集，已去重
func DiffSetFunc[T any](src, dst []T, equal equalFunc[T]) []T {
	var res = make([]T, 0, len(src))
	for _, val := range src {
		if !ContainsFunc[T](dst, func(src T) bool {
			return equal(src, val)
		}) {
			res = append(res, val)
		}
	}
	return deduplicateFunc[T](res, equal)

}
