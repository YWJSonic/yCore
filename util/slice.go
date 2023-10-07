package util

import (
	"math/rand"
	"sort"
)

// 取得目標存在於來源陣列的索引值
//
// @params []T 索引來源
//
// @params T 索引目標
//
// @return int 索引值
func Index[T comparable](source []T, target T) int {
	for i := range source {
		if target == source[i] {
			return i
		}
	}
	return -1
}

// 檢查站列是否包含目標
//
// @params []T 比對來源
//
// @params T 比對目標
//
// @return bool 是否包含物件
func Contains[T comparable](source []T, target T) bool {
	return Index(source, target) > -1
}

// 檢查目標陣列是否有任一物件存在來源陣列內
//
// @params []T 比對來源
//
// @params []T 比對目標
//
// @return bool 是否包含所有物件
func ContainAnyoneMulti[T comparable](source []T, targets []T) bool {
	for _, target := range targets {
		if Contains(source, target) {
			return true
		}
	}
	return false
}

// 檢查目標陣列是否全部物件存在來源陣列內
// (排除重複計算ex: source:[2], target:[2,2], return: false)
//
// @params []T 比對來源
//
// @params []T 比對目標
//
// @return bool 是否包含所有物件
func ContainAllMulti[T comparable](source []T, targets []T) bool {
	var idx int
	clone := make([]T, len(source))
	copy(clone, source)
	for _, target := range targets {
		if idx = Index(clone, target); idx == -1 {
			return false
		}

		clone = append(clone[:idx], clone[idx+1:]...)
	}
	return true
}

// 陣列使否完全相同
//
// @params []T 比對陣列
func EquarList[T comparable](a ...T) bool {
	for i, count := 0, len(a)-1; i < count; i++ {
		if a[i] != a[i+1] {
			return false
		}
	}
	return true
}

// 兩陣列是否相等
//
// @params []T 比對來源
//
// @params []T 比對目標
//
// @return bool 是否相等
func EquarTargetList[T comparable](source, target []T) bool {
	if len(source) != len(target) {
		return false
	}

	for i, count := 0, len(source); i < count; i++ {
		if source[i] != target[i] {
			return false
		}
	}

	return true
}

// 刪除第一次出現的目標物件
//
// @params []T 索引列表
//
// @params T 目標物件
//
// @return []T 複製結果
func RemoveFirst[T comparable](source []T, target T) []T {
	clone := make([]T, len(source))
	copy(clone, source)
	for i, count := 0, len(clone); i < count; i++ {
		if clone[i] == target {
			clone = append(clone[:i], clone[i+1:]...)
			return clone
		}
	}
	return clone
}

// 刪除所有目標第一次出現的物件
//
// @params []T 索引列表
//
// @params T 目標物件
//
// @return []T 複製結果
func RemoveFirstArray[T comparable](source []T, targets []T) []T {
	idxs := []int{}
	for idx, s := range source {
		if Contains(targets, s) {
			idxs = append(idxs, idx)
		}
	}

	return RemoveIndexMulti(source, idxs)
}

// 刪除索引直來源 0~n
//
// @params source 索引列表
// @params int 索引直範圍 0~n
func RemoveIndex[T comparable](source []T, idx int) []T {
	if idx < 0 || len(source) <= idx {
		return source
	}
	return append(source[:idx], source[idx+1:]...)
}

// 刪除索引直來源 1~n
//
// @params source 索引列表
// @params int 索引直範圍 1~n
func RemoveIndexCount[T comparable](source []T, idx int) []T {
	if idx < 1 || len(source) < idx {
		return source
	}
	return append(source[:idx-1], source[idx:]...)
}

// 刪除索引直來源 1~n
//
// @params source 索引列表
// @params []int 索引直範圍 0~n
func RemoveIndexMulti[T comparable](source []T, idxs []int) []T {
	if len(idxs) == 0 {
		return source
	}

	for _, idx := range idxs {
		if idx < 0 || len(source) <= idx {
			return source
		}
	}

	newSource := make([]T, len(source))
	for idx, v := range source {
		if !Contains(idxs, idx) {
			newSource = append(newSource, v)
		}
	}

	return newSource
}

func Shuffle[T int64](source []T) []T {
	count := len(source)
	if count < 2 {
		return source
	}

	indexs := make([]int, count)
	for i, count := 0, len(source); i < count; i++ {
		indexs[i] = i
	}
	s := rand.NewSource(ServerTimeNow().Unix())
	r := rand.New(s)
	r.Shuffle(count, func(i, j int) {
		indexs[i], indexs[j] = indexs[j], indexs[i]
	})

	newArr := []T{}
	for i := 0; i < len(indexs); i++ {
		newArr = append(newArr, source[indexs[i]])
	}
	return newArr
}

func ShufflePoint[T int64](source []T) {
	count := len(source)
	if count < 2 {
		return
	}

	s := rand.NewSource(ServerTimeNow().Unix())
	r := rand.New(s)
	r.Shuffle(count, func(i, j int) {
		source[i], source[j] = source[j], source[i]
	})

}

func Sort[T int64](source []T, desc ...bool) []T {
	count := len(source)
	if count < 2 {
		return source
	}

	indexs := make([]int, count)
	for i, count := 0, len(source); i < count; i++ {
		indexs[i] = i
	}

	if len(desc) > 0 && desc[0] {
		sort.Slice(indexs, func(i, j int) bool {
			return source[indexs[i]] > source[indexs[j]]
		})
	} else {
		sort.Slice(indexs, func(i, j int) bool {
			return source[indexs[i]] < source[indexs[j]]
		})
	}

	newArr := []T{}
	for i := 0; i < len(indexs); i++ {
		newArr = append(newArr, source[indexs[i]])
	}
	return newArr
}

func SortPoint[T int64 | float64](source []T, desc ...bool) {
	count := len(source)
	if count < 2 {
		return
	}

	if len(desc) > 0 && desc[0] {
		sort.Slice(source, func(i, j int) bool {
			return source[i] > source[j]
		})
	} else {
		sort.Slice(source, func(i, j int) bool {
			return source[i] < source[j]
		})
	}
}
