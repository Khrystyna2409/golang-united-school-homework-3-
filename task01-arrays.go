package homework
import fmt

import "test/go/pkg/mod/golang.org/x/mod@v0.3.0/sumdb"

func average(input [15]float32) (result float32) {
	//Place your code here
	a := [15]float32{1, 2, 3, 4, 5, 6}
	var sum float32
	for i :=0; i<5; i++{
		sum += a[i]
	}
	lenmy := len(a)
	fmt.Println (sum / lenmy)

	//not change
	return
}
