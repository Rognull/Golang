package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var words = []string{
	"helloz",
	"buzzy",
	"zombie",
	"zoom",
	"mangoz",
	"anyaz",
}

func main() {
	rand.Seed(time.Now().UnixNano())
	word := words[rand.Intn(len(words))]
	mistakes := 0
	wordMap := map[string]bool{}
	for mistakes < 8 {
		getChar(wordMap, word, &mistakes)
		print(&mistakes, wordMap, word)
	}
	if mistakes == 8 {
		println("GAMEOVER YOU LOSE")
	} else {
		println("YOU WIN!!")
	}
}

func print(mistakes *int, wordMap map[string]bool, word string) {
	comparestr := ""
	fmt.Print("Your word is : ")
	for _, b := range word {
		if b == ' ' {
			fmt.Print(" ")
		} else if wordMap[string(b)] {
			fmt.Print(string(b))
			comparestr += string(b)

		} else {
			fmt.Print("_")
		}
		fmt.Print(" ")
	}
	fmt.Print("\n")
	path := "hungman/"
	path += strconv.Itoa(*mistakes)

	b, err := os.ReadFile(path)

	if err != nil {
		panic(err)
	}
	fmt.Print(string(b))
	if comparestr == word {
		*mistakes = 9
	}
}

func getChar(wordMap map[string]bool, word string, mistakes *int) {
	fmt.Println("Write your letter: ")
	r := bufio.NewReader(os.Stdin)
	ch, err := r.ReadString('\n')
	if err != nil {
		panic(err)
	}
	ch = strings.TrimSpace(ch)
	if len(ch) != 1 || ch == " " {
		for len(ch) != 1 {
			fmt.Println("Error, Write only 1 letter")
			ch, err = r.ReadString('\n')
			if err != nil {
				panic(err)
			}
			ch = strings.TrimSpace(ch)
		}
	}

	flag := 0
	for _, b := range word {
		if ch == string(b) {
			wordMap[ch] = true
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
