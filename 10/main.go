package main

import "fmt"

func main() {
	temps := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	res := make(map[int64][]float64)
	for _, v := range temps {
		grp := int64(v/10) * 10
		res[grp] = append(res[grp], v)
	}
	fmt.Println(res)
}
