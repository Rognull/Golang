package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

var words = []string{
	"helloz",
	"buzy",
	"zombie",
	"zoom",
	"mangoz",
	"anyaz",
}

func main() {
	rand.Seed(time.Now().UnixNano())
	word := words[rand.Intn(len(words))]
	mistakes := 0
	var ch rune
	//var wordMap map[rune]bool
	wordMap := map[rune]bool{}

	for _, w := range word {
		if w == ' ' {
			wordMap[w] = true
		} else {
			wordMap[w] = false
		}
	}
	fmt.Scan(&ch)
	fmt.Println(ch)
	getChar(wordMap, &mistakes, ch)
	print(mistakes, wordMap, word)

}

func print(mistakes int, wordMap map[rune]bool, word string) {
	fmt.Print("Your word is : ")
	for _, b := range word {
		if !wordMap[b] {
			fmt.Print("_ ")
		} else {
			fmt.Printf("%c ", b)
		}
	}
	fmt.Print("\n")
	path := "c:/Projects For Prog/Golang/hello world/hungman/2"
	//path += strconv.Itoa(mistakes)
	/*f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			panic(err)
		}
	}()*/
	b, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	fmt.Print(string(b))
}

func getChar(wordMap map[rune]bool, mistakes *int, ch rune) {
	fmt.Println("Write your letter: ", ch)
	flag := 0
	for c, b := range wordMap {
		if c == ch && !b {
			wordMap[c] = true
			flag = 1
		}
	}
	if flag == 0 {
		fmt.Println("I'ts misstake !")
		*mistakes++
	} else {
		fmt.Println("Okey, we have this letter in word.")
	}
}
