package main

import (
	"fmt"

	"strings"

	"github.com/texttheater/golang-levenshtein/levenshtein"
)

// func mhash(b []byte) uint64 { return metro.Hash64(b, 0) }
// func testMinHash() {
// 	m1 := minhash.NewMinWise(spooky.Hash64, mhash, 10)
// 	m2 := minhash.NewMinWise(spooky.Hash64, mhash, 10)

// 	s1 := strings.ToLower("kitten")
// 	b1 := []byte(s1)
// 	s2 := strings.ToLower("sitten")
// 	b2 := []byte(s2)

// 	startMs := time.Now().UnixNano() / 1e6
// 	for i := 0; i < 1; i++ {
// 		for _, c := range b1 {
// 			d := make([]byte, 1)
// 			d[0] = c
// 			fmt.Printf("%s, %d\n", d, c)
// 			m1.Push(d)
// 		}
// 		fmt.Println("next one")
// 		for _, c := range b2 {
// 			d := make([]byte, 1)
// 			d[0] = c
// 			fmt.Printf("%s, %d\n", d, c)
// 			m2.Push(d)
// 		}

// 		fmt.Printf("sim:%f\n", m1.Similarity(m2))
// 	}
// 	endMs := time.Now().UnixNano() / 1e6
// 	fmt.Printf("cost %d ms\n", endMs-startMs)
// }

func testEditDist() {
	str1 := "Bili哔哩哔哩"
	str2 := "BiLI哔哩哔哩"
	ratio := levenshtein.RatioForStrings(
		[]rune(strings.ToLower(str1)),
		[]rune(strings.ToLower(str2)),
		levenshtein.DefaultOptionsWithSub)
	dist := levenshtein.DistanceForStrings(
		[]rune(strings.ToLower(str1)),
		[]rune(strings.ToLower(str2)),
		levenshtein.DefaultOptionsWithSub)
	fmt.Printf("ratio:%f, dist:%d", ratio, dist)

}

func main() {
	testEditDist()
	//testMinHash()
}
