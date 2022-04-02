package main

import (
	"fmt"
)

/// Принцип разделяй и влавствуй

/*
	Задача: найти наибольший квадрат на который можно разбить заданный прямоугольник
	рекурсия
*/

func DivideAndRule(width, height int) (result int) {
	if height > width {
		width, height = height, width
	}
	if width%height == 0 {
		return height
	} else {
		return DivideAndRule(width%height, height)
	}
}

/*
	Задача: сумма через рекурсию
*/
func SumMass(mass []int, sum int) int {
	if len(mass) >= 1 {
		sum += mass[0]
		return SumMass(mass[1:], sum)
	} else {
		return sum
	}
}

/*
	Задача: бинарный поиск и снова рекусия
*/
func BinarySearch(mass []int, num int) int {
	sizeMass := len(mass)
	halfMass := sizeMass / 2
	if mass[halfMass] > num {
		newMass := mass[:halfMass]
		return BinarySearch(newMass, num)
	} else if mass[halfMass] < num {
		newMass := mass[halfMass:]
		return BinarySearch(newMass, num)
	} else {
		return mass[halfMass]
	}
}

/*
	Задача: быстрая сортировка и снова рекусия
*/
func QuickSort(mass []int) (resMass []int) {
	if len(mass) < 2 {
		return mass
	}
	halfMass := len(mass) / 2
	pivot := mass[halfMass]
	less := make([]int, 0, 0)
	greater := make([]int, 0, 0)
	for idx, a := range mass {
		if idx == halfMass {
			continue
		}
		if a <= pivot {
			less = append(less, a)
		} else {
			greater = append(greater, a)
		}
	}

	resMass = append(append(append(resMass, QuickSort(less)...), pivot), QuickSort(greater)...)
	return
}

func main() {
	var res int

	// Поиск наибольшего квадрата
	//res = DivideAndRule(1680, 640)

	// Сумма массива рекурсией
	//mass := []int{2, 4, 6}
	//res = SumMass(mass, 0)

	// Бинарный поиск
	//s1 := rand.NewSource(time.Now().UnixNano())
	//r1 := rand.New(s1)
	//randNum := r1.Intn(100)
	//mass := make([]int, 100, 100)
	//for idx := range mass{
	//	mass[idx] = idx
	//}
	//fmt.Println("guessed: ", randNum);
	//res = BinarySearch(mass, randNum)

	// Быстрая сортировка
	//s1 := rand.NewSource(time.Now().UnixNano())
	//r1 := rand.New(s1)
	//mass := make([]int, 10, 10)
	//for idx := range mass {
	//	randNum := r1.Intn(100)
	//	mass[idx] = randNum
	//}
	//fmt.Println("source mass: ", mass)
	//resMass := QuickSort(mass)
	//fmt.Println(resMass)

	fmt.Println(res)
}
