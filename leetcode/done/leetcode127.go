package main

import "fmt"

func main() {
	fmt.Println(ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log"}))
}
func ladderLength(beginWord string, endWord string, wordList []string) int {
	var findWord func(string)
	ch := make(chan string, 10000)
	ch <- beginWord

	findWord = func(word string) {
		for i := 0; i < len(word); i++ {
			wordTemp := make([]byte, len(word))
			copy(wordTemp, word)
			for j := 0; j < 26; j++ {
				wordTemp[i] = byte(int('a') + j)
				strWordTemp := string(wordTemp)
				for k, v := range wordList {
					if v == strWordTemp {
						ch <- v
						copy(wordList[k:], wordList[k+1:])
						wordList = wordList[:len(wordList)-1]
					}
				}
			}
		}
	}

	findWord(<-ch)
	cnt := 2
	for len(ch) != 0 {
		chCnt := len(ch)
		for num := 0; num < chCnt; num++ {
			temp := <-ch
			if temp == endWord {
				return cnt
			}
			findWord(temp)
		}
		cnt++
	}

	return 0
}
