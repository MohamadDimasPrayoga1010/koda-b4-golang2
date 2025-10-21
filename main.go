package main

import "fmt"

func main(){
	scores := []int{50, 75, 66, 20, 32, 90}

	inputScore := [] int{}
	inputScore = append(inputScore, scores[:3]...)
	inputScore = append(inputScore, 88)
	inputScore = append(inputScore, scores[4:]...)

	for x := range len(inputScore){
		fmt.Println(inputScore[x])
	}
}