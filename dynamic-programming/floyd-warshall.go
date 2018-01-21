package main

import (
	"fmt"
	"math"
)

func floydWashall(g [][]graph) [][]float64 {
	var results [][]float64 = make([][]float64, len(g))
	for i := range results {
		xs := make([]float64, len(g))
		for j := range xs {
			xs[j] = math.Inf(1)
		}
		xs[i] = 0
		results[i] = xs
	}

	for i, gv := range g {
		for _, value := range gv {
			results[i][value.to] = value.wt
		}
	}

	for k, dk := range results {
		for _, di := range results {
			for j, dij := range di {
				d := di[k] + dk[j]
				if dij > d {
					di[j] = d
				}
			}
		}
	}

	return results
}

type graph struct {
	to int
	wt float64
}

func main() {
	var g [][]graph = [][]graph{
		1: {{2, 3}, {3, 8}, {5, -4}},
		2: {{4, 1}, {5, 7}},
		3: {{2, 4}},
		4: {{1, 2}, {3, -5}},
		5: {{4, 6}},
	}

	result := floydWashall(g)
	for _, d := range result {
		fmt.Println(d)
	}
}
