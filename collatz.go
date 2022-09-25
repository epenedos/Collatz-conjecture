package main

import "fmt"

//    "os"

//    "math/rand"

//    "github.com/go-echarts/go-echarts/v2/types"

func coll(r int) (res int) {

	if r%2 == 0 {
		res = r / 2
	} else {
		res = r*3 + 1
	}

	return res
}

func main() {

	var i int = 1000

	fmt.Println(i)
	for {
		n := coll(i)
		fmt.Println(n)

		if n == 1 {
			break
		}
		i = n
	}

}
