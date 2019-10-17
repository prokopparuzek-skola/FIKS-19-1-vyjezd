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

func bts(city *[]int, weight int, Sx, Sy int) int {
	var graph []vertex
	var queueA, queueF []int
	graph = make([]vertex, len(*city))
	queueA = make([]int, 0)
	queueF = make([]int, 0)

	for i := range graph {
		if (*city)[i] < 0 {
			graph[i].depth = (*city)[i]
		} else {
			graph[i].depth = -1
		}
		graph[i].parents = make([]int, 0)
	}
	graph[Sy*weight+Sx].depth = 0
	graph[Sy*weight+Sx].parents = append(graph[Sy*weight+Sx].parents, Sy*weight+Sx+weight)
	queueA = append(queueA, Sy*weight+Sx)

	for {
		for _, v := range queueA {
			stat := makeStep(&graph, v, &queueF, weight)
			switch stat {
			case 1:
				queueA = queueF
				queueF = make([]int, 0)
			case 2:
				index := queueF[len(queueF)-1]
				length := graph[index].depth
				fmt.Printf("%d\n", length)
				return 0
			case 3:
				fmt.Printf("No solution\n")
				return 0
			}
		}
	}
}

func makeStep(graph *[]vertex, v int, queueF *[]int, weight int) int {
	//fmt.Printf("%d: ", v)
	var canGo bool = false
	x := v % weight
	y := v / weight
	for _, p := range (*graph)[v].parents {
		Px := p % weight
		Py := p / weight
		var sons []int = make([]int, 0)

		if y < Py {
			if y > 0 {
				sons = append(sons, v-weight)
			}
			if x < weight-1 {
				sons = append(sons, v+1)
			}
		} else if y == Py {
			if x < Px {
				if y > 0 {
					sons = append(sons, v-weight)
				}
				if x > 0 {
					sons = append(sons, v-1)
				}
			} else {
				if y < len(*graph)/weight-1 {
					sons = append(sons, v+weight)
				}
				if x < weight-1 {
					sons = append(sons, v+1)
				}
			}
		} else {
			if y < len(*graph)/weight-1 {
				sons = append(sons, v+weight)
			}
			if x > 0 {
				sons = append(sons, v-1)
			}
		}
		for _, s := range sons {
			if !isIn(v, &(*graph)[s].parents) && (*graph)[s].depth != WALL {
				*queueF = append(*queueF, s)
				if (*graph)[s].depth == END {
					(*graph)[s].depth = (*graph)[v].depth + 1
					return 2
				}
				(*graph)[s].depth = (*graph)[v].depth + 1
				(*graph)[s].parents = append((*graph)[s].parents, v)
				canGo = true
			}
		}
		/*
			for _, x := range sons {
				fmt.Printf("%d(%d) ", x, (*graph)[x].depth)
			}
			println()
		*/
	}
	if (*graph)[v].depth == 0 {
		(*graph)[v].parents = make([]int, 0)
	}
	if canGo {
		return 1
	} else {
		return 3
	}
}

func isIn(what int, where *[]int) bool {
	for _, x := range *where {
		if x == what {
			return true
		}
	}
	return false
}

func main() {
	var T int
	fmt.Scanf("%d", &T)
	for ; T > 0; T-- {
		var M, N, K int
		var Sx, Sy, Cx, Cy int
		var city []int
		fmt.Scanf("%d%d%d", &N, &M, &K) // M šířka
		city = make([]int, M*N)
		fmt.Scanf("%d %d %d %d", &Sy, &Sx, &Cy, &Cx)
		//fmt.Printf(" %d %d %d %d %d %d %d", N, M, K, Sx, Sy, Cx, Cy)
		Sx--
		Sy--
		Cx--
		Cy--
		city[Cy*M+Cx] = END
		for ; K > 0; K-- {
			var Wx, Wy int
			fmt.Scanf("%d%d", &Wy, &Wx)
			Wx--
			Wy--
			city[Wy*M+Wx] = WALL
		}
		if Sy == Cy && Sx == Cx {
			fmt.Printf("0\n")
			continue
		}
		bts(&city, M, Sx, Sy)
	}
}
