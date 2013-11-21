package main

import (
	//"code.google.com/p/go-tour/wc"
	"fmt"
	"strings"
)

//WordCount returns a map of each word
//that appears in the sentence as key
//the number of the word as value.
func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	fmt.Println(words)
	result := make(map[string]int)
	for _, word := range words {
		//_, ok := result[word]
		if _, ok := result[word]; ok {
			result[word]++
		} else {
			result[word] = 1
		}
		//fmt.Println(result[word])
	}
	return result
	//return map[string]int{"x": 1}
}

func main() {
	//wc.Test(WordCount)
	fmt.Println(WordCount("henryxian is a stupid boy boy boy!"))
	fmt.Println(WordCount("This is another another another horrible test!"))
}
