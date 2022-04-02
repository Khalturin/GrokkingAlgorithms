package main

import (
	"fmt"
)

//func Union(setOne, setTwo map[string]string) map[string]string {
//	res := make(map[string]string)
//	for key, val := range setOne {
//		res[key] = val
//	}
//	for key, val := range setTwo {
//		res[key] = val
//	}
//	return res
//}
//
//func Intersection(setOne, setTwo map[string]string) map[string]string {
//	res := make(map[string]string)
//	if len(setOne) > len(setTwo) {
//		setOne, setTwo = setTwo, setOne
//	}
//	for key, val := range setOne {
//		if _, ok := setTwo[key]; ok {
//			res[key] = val
//		}
//	}
//
//	return res
//}

func Union(setOne, setTwo []string) []string {
	res := make([]string,0,0)
	for _, val := range setOne {
		if !isHaveValue(res, val) {
			res = append(res, val)
		}else{
			continue
		}
	}
	for _, val := range setTwo{
		if !isHaveValue(res, val) {
			res = append(res, val)
		}else{
			continue
		}
	}
	return res
}

func isHaveValue(mass []string, searchingVal string) bool{
	for _, val := range mass{
		if searchingVal == val{
			return true
		}
	}
	return false
}

func Intersection(setOne, setTwo []string) []string {
	res := make([]string,0,0)
	if len(setOne) > len(setTwo) {
		setOne, setTwo = setTwo, setOne
	}
	for _, val := range setOne {
		if isHaveValue(setTwo, val) {
			res = append(res, val)
		}
	}

	return res
}

func Difference(From, To []string)[]string {
	res := make([]string,0,0)
	for _, val := range From{
		isHave := false
		for _, val2 := range To{
			if val == val2{
				isHave = true
				break
			}
		}
		if !isHave{
			res = append(res, val)
		}
	}
	return res
}

//func CheckSet() {
//	setOne := make(map[string]string)
//	setTwo := make(map[string]string)
//
//	setOne["1"] = "one"
//	setOne["2"] = "two"
//	setOne["3"] = "three"
//
//	setTwo["2"] = "two"
//	setTwo["3"] = "three"
//	setTwo["4"] = "fore"
//
//	fmt.Println("Source: setOne: ", setOne)
//	fmt.Println("        setTwo: ", setTwo)
//	fmt.Println("________________________________________________")
//	fmt.Println("Union:          ", Union(setOne, setTwo))
//	fmt.Println("Intersection:   ", Intersection(setOne, setTwo))
//}

func CheckSet2() {
	setOne := []string{"one", "two", "three"}
	setTwo := []string{"two", "three", "fore"}

	fmt.Println("Source: setOne: ", setOne)
	fmt.Println("        setTwo: ", setTwo)
	fmt.Println("________________________________________________")
	fmt.Println("Union:          ", Union(setOne, setTwo))
	fmt.Println("Intersection:   ", Intersection(setOne, setTwo))
	fmt.Println("Difference:     ", Difference(setOne, setTwo))
}


func Greedy(stations map[string][]string, statesNeeded []string) []string {
	finalStations := []string{}

	for len(statesNeeded) > 0 {
		statesCovered := []string{}
		bestStation := ""
		for station, states := range stations{
			covered := Intersection(statesNeeded, states)
			if len(covered) > len(statesCovered){
				bestStation = station
				statesCovered = covered
			}
		}
		//statesNeeded -= statesCovered
		statesNeeded = Difference(statesNeeded, statesCovered)
		finalStations = append(finalStations, bestStation)
	}


	return finalStations
}

func main() {
	//CheckSet2()

	statesNeeded := []string{"mt", "wa", "or", "id", "nv", "ut", "ca", "az"}
	stations := make(map[string][]string)
	stations["kone"] = []string{"id", "nv", "ut"} //make([]string, 0, 0)
	stations["ktwo"] = []string{"wa", "id", "mt"}
	stations["kthree"] = []string{"or", "nv", "ca"}
	stations["kfore"] = []string{"nv", "ut"}
	stations["kfive"] = []string{"ca", "az"}

	fmt.Println(Greedy(stations, statesNeeded))
}
