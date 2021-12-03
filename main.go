package main

import (
	"fmt"
	"math/rand"
)

var lyric = map[int]string{
	1: "Twinkle, twinkle, little star",
	2: "How I wonder what you are",
	3: "Up above the world so high",
	4: "Like a diamond in the sky",
}

var result string

func main() {
	resultRandom := random(0 , 4)
	fullLyric := algoFindLyric(1)
	fmt.Println(fullLyric)
	fmt.Println(resultRandom)
}

func random(min int, max int) int {
	return rand.Intn(max-min) + min
}


