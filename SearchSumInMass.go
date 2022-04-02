package main

import "fmt"

func BruteForce(mass []int, k int) (result []int) {
	for idx1, val1 := range mass {
		for idx2, val2 := range mass {
			if idx1 == idx2 {
				continue
			}
			summ := val1 + val2
			if summ == k {
				result = append(result, val1, val2)
				return
			}
		}
	}
	return
}

func CacheSet(mass []int, k int) (result []int) {
	mp := make(map[int]int)

	for idx, val := range mass {
		if idx == 0 {
			mp[val] = idx
		}
		searchingNum := k - val
		if _, inMap := mp[searchingNum]; inMap {
			result = append(result, searchingNum, val)
			return
		} else {
			mp[val] = idx
		}
	}
	return
}

func Binary(mass []int, k int) (result []int) {
	for idx, val := range mass {
		searchingNum := k - val
		l := idx + 1
		r := len(mass) - 1
		for l <= r {
			mid := l + (r-l)/2
			if mass[mid] == searchingNum {
				result = append(result, val, searchingNum)
				return
			}
			if searchingNum < mass[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
	}
	return
}

func TwoPoiners(mass []int, k int) (result []int) {
	l := 0
	r := len(mass) - 1
	for l < r {
		sum := mass[l] + mass[r]
		if sum == k {
			result = append(result, mass[l], mass[r])
			return
		}
		if sum < k{
			l++
		}else{
			r--
		}
	}
	return
}

func main() {
	//mass := []int{-1, 2, 5, 8}
	//k := 7
	//mass := []int{-3, -1, 0, 2, 6}
	//k := 6
	//mass := []int{2, 4, 5}
	//k := 8
	//mass := []int{-2, -1, 1, 2}
	//k := 0
	//mass := []int{-3, 0, 1, 3, 4}
	//k := 5
	mass := []int{-9, -5, -2, -1, 1, 4, 9, 11}
	k := 3

	result := BruteForce(mass, k)
	fmt.Printf("BruteForce: %v\n", result)
	result = CacheSet(mass, k)
	fmt.Printf("CacheSet: %v\n", result)
	result = Binary(mass, k)
	fmt.Printf("Binary: %v\n", result)
	result = TwoPoiners(mass, k)
	fmt.Printf("TwoPoiners: %v\n", result)
}
