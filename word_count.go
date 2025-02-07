package main

import(
	"fmt"
	"strings"
	"bufio"
	"os"
	"unicode"
)

func getInput() string{
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a sentence in which you want to count the words: ")
	sentence,err := reader.ReadString('\n')
	if err != nil{
		fmt.Println("Error reading input: ",err)
	}
	sentence = strings.TrimSpace(sentence)
	return sentence
}

func wordCount(s string) map[string]int{
	frequency := make(map[string]int)
	words := strings.Fields(s)

	for _, word := range words{
		if (len(word) == 1) && (unicode.IsPunct(rune(word[0]))){
			continue
		}
		word = strings.ToLower(word)
		frequency[word] += 1
	}
	return frequency
}

func printMap(m map[string]int){
	for key, value := range m{
		fmt.Println(key, " : ", value)
	}
}

func main(){
	sentence := getInput()
	printMap(wordCount(sentence))
	//fmt.Println(checkPalindrome("abba"))
}

func checkPalindrome(s string) bool{
	l,r := 0 , len(s)-1
	for l < r{
		if s[l] != s[r]{
			return false
		}
		l++
		r--
	}
	return true
}