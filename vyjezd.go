package main

import "fmt"

const (
	END  int = -2
	WALL int = -3
)

type vertex struct {
	depth   int
	parents []int
}

func bts(city []int, weight int, Sx, Sy int) int {
	height := len(city) / weight
	var graph []vertex
	var queueA, queueF []int
	graph = make([]vertex, weight*height)
	queueA = make([]int, 0)
	queueF = make([]int, 0)

	for i := range graph {
		graph[i].depth = -1
		graph[i].parents = make([]int, 0)
	}
	graph[Sy*weight+Sx].depth = 0
	queueA = append(queueA, Sy*weight+Sx)

	for {
		for _, v := range queueA {
		}
	}
	return 0
}

func main() {
	var T int
	fmt.Scanf("%d", &T)
	for ; T > 0; T-- {
		var M, N, K int
		var Sx, Sy, Cx, Cy int
		var city []int
		fmt.Scanf("%d%d", &N, &M, &K)
		N++ // indexuje se od 1
		M++
		city = make([]int, M*N)
		fmt.Scanf("%d%d%d%d", &Sx, &Sy, &Cx, &Cy)
		city[Cy*M+Cx] = END
		for ; K > 0; K-- {
			var Wx, Wy int
			fmt.Scanf("%d%d", &Wx, &Wy)
			city[Wy*M+Wx] = WALL
		}
	}
}
