package slice

// UnionSet 并集，只支持 comparable
// 已去重
// 返回值的元素顺序是不定的

func UnionSet[T comparable](src, dst []T) []T {
	srcMap, dspMap := toMap[T](src), toMap[T](dst)
	for key := range srcMap {
		dspMap[key] = struct{}{}
	}

	var res = make([]T, 0, len(dspMap))
	for key := range dspMap {
		res = append(res, key)
	}
	return res
}

// UnionSetFunc 并集，支持任意类型
// 优先使用 UnionSet
// 已去重

func UnionSetFunc[T any](src, dst []T, equal equalFunc[T]) []T {
	var ret = make([]T, 0, len(src)+len(dst))
	ret = append(ret, dst...)
	ret = append(ret, src...)

	return deduplicateFunc[T](ret, equal)
}
