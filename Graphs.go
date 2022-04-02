package main

import (
	"fmt"
	"github.com/gammazero/deque"
)

func Search(deq deque.Deque, name string) {

}

func PersonIsSeller(person string) bool {
	if person == "Thom" {
		return true
	}
	return false
}

func PushBack(q * deque.Deque, persons []string){
	for _, person := range persons{
		q.PushBack(person)
	}
}

func PopFront(q * deque.Deque) string{
	val := q.PopFront()
	var person string
	if v, ok := val.(string); ok{
		person = v
	}else{
		fmt.Println("Ooooops")
	}
	return person
}

func isHaveInSlice(slice []string, name string) bool{
	for _, person := range slice{
		if name == person{
			return true
		}
	}
	return false
}

func SearchPerson(q * deque.Deque, graph map[string][]string) bool{
	searched := make([]string, 0, 0)
	for q.Len() != 0 {
		person := PopFront(q)
		if !isHaveInSlice(searched, person) {
			if PersonIsSeller(person) {
				fmt.Println(person + " is a mango seller!")
				return true
			} else {
				fmt.Println(person + " is NOT a mango seller!")
				searched = append(searched, person)
				PushBack(q, graph[person])
			}
		}
	}
	return false
}

/*
	Поиск в ширину
 */
func main() {
	graph := make(map[string][]string)

	graph["You"] = append(graph["You"], "Alice", "Bob", "Claire")
	graph["Bob"] = append(graph["Bob"], "Anuj", "Peggy")
	graph["Alice"] = append(graph["Alice"], "Peggy")
	graph["Claire"] = append(graph["Claire"], "Thom", "Jonny")
	graph["Anuj"] = append(graph["Anuj"], "")
	graph["Peggy"] = append(graph["Peggy"], "")
	graph["Thom"] = append(graph["Thom"], "")
	graph["Jonny"] = append(graph["Jonny"], "")

	var q deque.Deque
	PushBack(&q, graph["You"])

	if !SearchPerson(&q, graph) {
		fmt.Println("Ua ua ua")
	}
}
