package utils

import (
	"fmt"
	"strings"
)

func MapDemos() {
	m := make(map[string]int)
	m["foo"] = 10
	m["bar"] = 20

	fmt.Println(m)
	fmt.Println(m["foo"])

	delete(m, "foo")
	fmt.Println(m)

	colors := map[string]int{"red": 1, "green": 2, "blue": 3}
	colors["yellow"] = 4
	fmt.Println(colors)
	fmt.Println(len(colors))

	/* _, found := colors["purple"] //=> val, bool
	fmt.Println(found) */

	colors["purple"] = 5
	if _, found := colors["purple"]; found {
		fmt.Println("purple exits")
	} else {
		fmt.Println("purple doesnot exist")
	}
}

func WordStats() {
	str := "Culpa pariatur commodo cillum nulla duis veniam esse incididunt officia ex deserunt Occaecat sunt do elit dolor mollit Ipsum ullamco cupidatat amet nostrud enim excepteur duis do ex quis labore ut dolore Excepteur ut nostrud elit ut consectetur Elit adipisicing Lorem sit amet sit reprehenderit Lorem aliquip dolore id commodo ut Id quis irure ut eiusmod aliqua labore laborum amet magna aliqua mollit mollit laboris"
	words := strings.Split(str, " ")
	stats := make(map[int]int)

	/* for index := 0; index < len(words); index++ {
		word := words[index]
		wordLength := len(word)
		if _, found := stats[wordLength]; !found {
			stats[wordLength] = 0
		}
		stats[wordLength] += 1
	} */

	for _, word := range words {
		wordLength := len(word)
		if _, found := stats[wordLength]; !found {
			stats[wordLength] = 0
		}
		stats[wordLength]++
	}

	for wordLength, wordCount := range stats {
		fmt.Printf("WordLength : %d\t WordCount : %d\n", wordLength, wordCount)
	}

}
