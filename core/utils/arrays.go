package utils

import (
	"math/rand/v2"
	"slices"
)

func InArray[T comparable](needle T, haystack []T) bool {
	for _, elem := range haystack {
		if needle == elem {
			return true
		}
	}
	return false
}

func RemoveDuplicates[T comparable](sliceList []T) []T {
	allKeys := make(map[T]bool)
	list := []T{}
	for _, item := range sliceList {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func FindDuplicates[T comparable](sliceList []T) []T {
	seen := make(map[T]bool)
	duplicates := []T{}

	for _, item := range sliceList {
		if seen[item] {
			duplicates = append(duplicates, item)
		} else {
			seen[item] = true
		}
	}

	return duplicates
}

func Intersection[T comparable](s1 []T, s2 []T) (result []T) {
	hash := make(map[T]bool)
	for _, e := range s1 {
		hash[e] = true
	}
	for _, e := range s2 {
		if hash[e] {
			result = append(result, e)
		}
	}
	return
}

// PrependToArrayWithLimitSize вставляет элемент в начало массива, при этом массив обрезается, чтобы его размер не был больше maxSize
func PrependToArrayWithLimitSize[T comparable](arr []T, element T, maxSize int) []T {
	arr = append([]T{element}, arr...)

	if len(arr) > maxSize {
		arr = arr[:maxSize]
	}

	return arr
}

type IDer interface {
	GetId() int64
}

func SortArrayByIds[T IDer](arr []T, ids []int64) {
	idsMap := map[int64]int{}
	for i, gameId := range ids {
		idsMap[gameId] = i
	}
	slices.SortFunc(arr, func(a, b T) int {
		positionA := idsMap[a.GetId()]
		positionB := idsMap[b.GetId()]
		if positionA < positionB {
			return -1
		} else if positionA > positionB {
			return 1
		}
		return 0
	})
}

type shufflePairs struct {
	Key, Value string
}

func ShuffleMap(m map[string]string) []shufflePairs {
	items := make([]shufflePairs, 0, len(m))
	for k, v := range m {
		items = append(items, shufflePairs{k, v})
	}

	rand.Shuffle(len(items), func(i, j int) {
		items[i], items[j] = items[j], items[i]
	})
	return items
}
