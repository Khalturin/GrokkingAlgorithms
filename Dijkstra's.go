package main

import (
	"fmt"
	"math"
	"strconv"
)

func printGraph(graph map[string]map[string]int) {
	for i, val := range graph {
		fmt.Print(i + ": ")
		for j, k := range val {
			fmt.Print(j + "(" + strconv.Itoa(k) + ") ")
		}
		fmt.Println("")
	}
}

func printExampleCosts(costs map[string]int) {
	fmt.Print("COSTS: plate: ", costs["plate"], " ")
	fmt.Print("poster: ", costs["poster"], " ")
	fmt.Print("guitar: ", costs["guitar"], " ")
	fmt.Print("drum: ", costs["drum"], " ")
	fmt.Print("piano: ", costs["piano"], "\n")
}

func printExampleParents(parents map[string]string) {
	fmt.Print("PARENTS: [plate: ", parents["plate"], "] ")
	fmt.Print("[poster: ", parents["poster"], "] ")
	fmt.Print("[guitar: ", parents["guitar"], "] ")
	fmt.Print("[drum: ", parents["drum"], "] ")
	fmt.Print("[piano: ", parents["piano"], "]\n")
}

func checkInProcessed(processed []string, key string) bool {
	for _, val := range processed {
		if val == key {
			return true
		}
	}
	return false
}

func find_lowes_cost_node(costs map[string]int, processed []string) string {
	lowest_cost := math.MaxInt
	lowest_cost_node := "None"
	for node := range costs {
		cost := costs[node]
		if cost < lowest_cost && !checkInProcessed(processed, node) {
			lowest_cost = cost
			lowest_cost_node = node
		}
	}
	return lowest_cost_node
}

func AlgorythmDijkstra(graph map[string]map[string]int) int {
	//costs := make(map[string]map[string]int)
	//costs["book"] = make(map[string]int)
	costs := make(map[string]int)
	costs["plate"] = 5
	costs["poster"] = 0
	costs["guitar"] = math.MaxInt
	costs["drum"] = math.MaxInt
	costs["piano"] = math.MaxInt

	parents := make(map[string]string)
	parents["plate"] = "book"
	parents["poster"] = "book"
	parents["guitar"] = ""
	parents["drum"] = ""
	parents["piano"] = ""

	var processed []string

	node := find_lowes_cost_node(costs, processed)
	for node != "None" {
		cost := costs[node]
		neighbors := graph[node]

		for i, _ := range neighbors {
			new_cost := cost + neighbors[i]
			if costs[i] > new_cost {
				costs[i] = new_cost
				parents[i] = node
			}
		}
		processed = append(processed, node)
		node = find_lowes_cost_node(costs, processed)

		printExampleCosts(costs)
		printExampleParents(parents)
		fmt.Println("PROCESSED:", processed)
		fmt.Println()
	}

	return costs["piano"]
	//printGraph(costs)
}

/*
     plate -15-> guitar
  5 /     \    30/      \20
  book       \/      piano
  0 \       /     \20   /10
    poster -35-> drum
*/
func main() {
	graph := make(map[string]map[string]int)
	graph["book"] = make(map[string]int)
	graph["book"]["plate"] = 5
	graph["book"]["poster"] = 0
	graph["plate"] = make(map[string]int)
	graph["plate"]["guitar"] = 15
	graph["plate"]["drum"] = 20
	graph["poster"] = make(map[string]int)
	graph["poster"]["guitar"] = 30
	graph["poster"]["drum"] = 35
	graph["guitar"] = make(map[string]int)
	graph["guitar"]["piano"] = 20
	graph["drum"] = make(map[string]int)
	graph["drum"]["piano"] = 10
	graph["piano"] = make(map[string]int)

	fmt.Println(AlgorythmDijkstra(graph))

	//printGraph(graph)
}