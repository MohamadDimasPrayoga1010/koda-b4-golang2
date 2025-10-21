package main

import "fmt"

func main() {
	scores := []int{50, 75, 60, 20, 32, 90}

	inputScore := []int{}
	// inputScore = append(inputScore, scores[:2]...)
	// inputScore = append(inputScore, 88)
	// inputScore = append(inputScore, scores[2:]...)

	for _, score := range scores {
		if score == 66 {
			inputScore = append(inputScore, 88)
		}
		inputScore = append(inputScore, score)

	}

	fmt.Println(inputScore)

	// for x := range len(inputScore){
	// 	fmt.Println(inputScore[x])
	// }

}
