package main

import "fmt"

const (
	START int = 1
	END   int = 2
	WALL  int = 3
)

func bts(city []int, widht int) int {
	weidht := len(city) / widht
	// do vrcholu se vstupuje z různých stran, zahrnout
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
		city[Sy*M+Sx] = START
		city[Cy*M+Cx] = END
		for ; K > 0; K-- {
			var Wx, Wy int
			fmt.Scanf("%d%d", &Wx, &Wy)
			city[Wy*M+Wx] = WALL
		}
	}
}
