package main

// STRING
import (
	"fmt"
	"strings"
)

const (
	str    = "something"
	substr = "abc"
)

func main() {
	// 1) len string, menghitung panjang string
	sentence := "Hello"
	lenSentence := len(sentence)
	fmt.Println(lenSentence)

	// 2) Compare string
	str1 := "abcd"
	str2 := "abcd"
	fmt.Println(str1 == str2)

	// 3) Contain
	res := strings.Contains(str, substr)
	fmt.Println(res)

	// 4) get substring
	value := "cat;dog"
	substring1 := value[4:]
	substring2 := value[4:len(value)]
	fmt.Println(substring1)
	fmt.Println(substring2)

	// 5) Replace
	strAwal := "katak katak katak"
	// indeks -1 untuk mengubah semua huruf
	replacedStr := strings.Replace(strAwal, "k", "R", 3)
	fmt.Println(replacedStr)

	// 6) Insert
	strAsli := "green"
	index := 2
	insertChar := strAsli[:index] + "HAY" + strAsli[index:]
	fmt.Println(strAsli, insertChar)
}
