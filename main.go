package main

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"os"
	"regexp"
	"sort"
	"strings"
)

const ALL_WORDS = "./data/allDictionaryWords.txt"
const ALPHABET = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func main() {

	// read in dictionary
	file, err := os.Open(ALL_WORDS)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	all := make([]string, 0, 466550) // capacity == wordcount from dict
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		all = append(all, scanner.Text())
	}

	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// get 20 random words
	size := 27
	words := make([]string, 0, size)
	r, _ := regexp.Compile("^[a-zA-Z]+$") // regex: no numbers or punctuation
	for i := 0; i < size; i++ {
		word := getValidPuzzleWord(all, r) // also gets rid of two letter words
		words = append(words, strings.ToUpper(word))
	}

	// build matrix for puzzle
	w := len(words[0]) // width
	if w < 20 {
		w = 20 // --> original challenge allowed 20 rows.
	}
	table := make([][]string, 0, w)
	for i := 0; i < w; i++ {
		line := make([]string, w)
		table = append(table, line)
	}

	// populate words of puzzle
	for _, v := range words {
		insertWordInPuzzle(v, table)
	}

	// leftover with random letters
	for i := 0; i < len(table); i++ {
		for j := 0; j < len(table[i]); j++ {
			if len(table[i][j]) == 0 {

				table[i][j] = string(ALPHABET[generateRandom(26, 0)])
			}
		}
	}

	// output: words + puzzle
	sort.Strings(words)
	for i, v := range words {
		log.Printf("%d.) %s", i+1, v)
	}
	for _, v := range table {
		log.Println(v)
	}
}

func insertWordInPuzzle(word string, table [][]string) [][]string {

	orient := generateRandom(9, 1)
	switch orient {
	// 1: horizontal
	case 1:
		t, err := handleHorizontal(word, table)
		if err != nil {
			log.Println(err)
			insertWordInPuzzle(word, table) // recurse
		}
		return t

	// 2: horizontal backwards
	case 2:
		t, err := handleHorizontal(reverser(word), table)
		if err != nil {
			log.Println(err)
			insertWordInPuzzle(word, table)
		}
		return t

	// 3: vertical
	case 3:
		t, err := handleVertical(word, table)
		if err != nil {
			log.Println(err)
			insertWordInPuzzle(word, table)
		}
		return t

	// 4: vertical backwards
	case 4:
		t, err := handleVertical(reverser(word), table)
		if err != nil {
			log.Println(err)
			insertWordInPuzzle(word, table)
		}
		return t

	// 5: diagonal up
	case 5:
		t, err := handleDiagonalUp(word, table)
		if err != nil {
			log.Println(err)
			insertWordInPuzzle(word, table)
		}
		return t

	// 5: diagonal up backwards
	case 6:
		t, err := handleDiagonalUp(reverser(word), table)
		if err != nil {
			log.Println(err)
			insertWordInPuzzle(word, table)
		}
		return t

	// 5: diagonal down
	case 7:
		t, err := handleDiagonalDown(word, table)
		if err != nil {
			log.Println(err)
			insertWordInPuzzle(word, table)
		}
		return t

	// 6: diagonal down backwards
	default:
		t, err := handleDiagonalDown(reverser(word), table)
		if err != nil {
			log.Println(err)
			insertWordInPuzzle(word, table)
		}
		return t
	}
}

func generateRandom(max, min int64) int64 {

	// between 1 and 6
	b := big.NewInt(max - min) // top-end exclusive
	r, err := rand.Int(rand.Reader, b)
	if err != nil {
		panic(err)
	}
	return r.Int64() + min // no zeros
}

func handleHorizontal(word string, table [][]string) (t [][]string, e error) {

	// select row and column indices
	r := generateRandom(int64(len(table)), 0)                    // 17, 0
	c := generateRandom(int64(len(table))-int64(len(word)-1), 0) //17 - 11 => 6, 0

	if !isAllowed(word, table[r][c:int(c)+len(word)]) {
		e = fmt.Errorf("'%s' is not allowed horizontally beginning at row %d, column index %d", word, r, c)
	} else {
		for _, v := range word {
			table[r][c] = string(v)
			c++
		}
	}
	return table, e
}

func handleVertical(word string, table [][]string) (t [][]string, e error) {

	// select row and column indices
	c := generateRandom(int64(len(table)), 0)
	r := generateRandom(int64(len(table))-int64(len(word)-1), 0)

	// build control row from vertical
	cur := make([]string, 0, len(word))
	for i := int(r); i < len(word)+int(r); i++ {
		cur = append(cur, table[i][c])
	}

	if !isAllowed(word, cur) {
		e = fmt.Errorf("'%s' is not allowed vertically beginning at row %d, column index %d", word, r, c)
	} else {
		for _, v := range word {
			table[r][c] = string(v)
			r++
		}
	}

	return table, e
}

func handleDiagonalDown(word string, table [][]string) (t [][]string, e error) {

	r := generateRandom(int64(len(table))-int64(len(word)-1), 0)
	c := generateRandom(int64(len(table))-int64(len(word)-1), 0)

	// control row
	cur := make([]string, 0, len(word))
	y := int(c)
	for i := int(r); i < (len(word))+int(r); i++ {
		cur = append(cur, table[i][y])
		y++

	}

	if !isAllowed(word, cur) {
		e = fmt.Errorf("'%s' is not allowed diagonally-down beginning at row index %d, column index %d", word, r, c)

	} else {
		for i, v := range word {

			table[int(r)+i][int(c)] = string(v)
			c++
		}
	}

	return table, e
}

func handleDiagonalUp(word string, table [][]string) (t [][]string, e error) {

	r := generateRandom(int64(len(table)), int64(len(word)-1))
	c := generateRandom(int64(len(table))-int64(len(word)-1), 0)

	// control row
	cur := make([]string, 0, len(word))
	x := int(r)
	for i := int(c); i < (len(word))+int(c); i++ {
		cur = append(cur, table[x][i])
		x--

	}

	if !isAllowed(word, cur) {
		e = fmt.Errorf("'%s' is not allowed diagonally-up beginning at row index %d, column index %d", word, r, c)

	} else {
		for i, v := range word {

			table[int(r)][int(c)+i] = string(v)
			r--
		}
	}

	return table, e
}

func isAllowed(chal string, cur []string) bool {

	for i, v := range cur {

		if len(v) != 0 && string(chal[i]) != v {
			return false
		}
	}
	return true
}

func reverser(fw string) (bw string) {
	for _, v := range fw {
		bw = string(v) + bw
	}
	return
}

func getValidPuzzleWord(list []string, r *regexp.Regexp) string {

	index := generateRandom(int64(len(list)), 0)
	word := list[index]
	if !r.MatchString(word) || len(word) < 3 {
		word = getValidPuzzleWord(list, r)
	}
	return word
}
