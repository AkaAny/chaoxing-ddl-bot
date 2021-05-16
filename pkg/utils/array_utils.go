package utils

import (
	"sync"
	"sync/atomic"
)

func Contains(u []string, sub string) bool {
	if u == nil {
		return false
	}
	for _, uItem := range u {
		if uItem == sub {
			return true
		}
	}
	return false
}

func ListContains(u []string, sub []string) bool {
	for _, subItem := range sub {
		if !Contains(u, subItem) {
			return false
		}
	}
	return true
}

func InMutableRemoveIf(u []string, predictor func(item string) bool) []string {
	var result []string
	for _, item := range u {
		if !predictor(item) {
			result = append(result, item)
		}
	}
	return result
}

func Merge(arrays ...[]string) []string {
	var resultWithRepeat []string
	for _, arr := range arrays {
		resultWithRepeat = append(resultWithRepeat, arr...)
	}
	var result []string
	var repeatMap = make(map[string]struct{})
	for _, item := range resultWithRepeat {
		if _, ok := repeatMap[item]; !ok {
			repeatMap[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

func ForEachAsync(u []string, consumer func(item string) error) error {
	var err error
	var wg sync.WaitGroup
	var waitCount = len(u)
	var remaining int64 = int64(len(u))
	wg.Add(waitCount)
	for _, item := range u {
		var itemCaptured = item
		go func() {
			err = consumer(itemCaptured)
			if err != nil {
				wg.Add(int(-remaining)) //提前结束
				return
			}
			atomic.AddInt64(&remaining, -1) //原子-1
			wg.Done()
		}()
	}
	wg.Wait()
	return err
}

func SingletonArray(item string) []string {
	return []string{item}
}
