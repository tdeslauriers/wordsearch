package main

import (
	"bufio"
	"regexp"
	"sort"

	"os"
	"testing"
)

var TEST_WORDS = [...]string{"JIUJITSU", "MANTRAS", "OVERSENSIBLENESS", "FORTYFIVE", "LASTJOB", "HAIR", "STUTSMAN", "HULME", "RECISSION", "ACCOLATED", "SUBGLABROUS", "OMOSTERNAL", "JNANAMARGA", "PULCIFER", "FLYOVER", "ANGLOMALTESE", "BITCHINESS", "SANGUINEOVASCULAR", "SCROOTCH", "SNIP"}

func TestReadDictionary(t *testing.T) {

	file, err := os.Open(ALL_WORDS)
	if err != nil {
		t.Log(err)
	}
	defer file.Close()

	inc := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inc += 1
	}
	t.Log(inc)

	if err = scanner.Err(); err != nil {
		t.Log(err)
	}
}

func TestRandom(t *testing.T) {

	// dictionary index
	for i := 0; i < 20; i++ {
		t.Log(generateRandom(466550, 1))
	}

	// orientation
	for i := 0; i < 20; i++ {
		t.Log(generateRandom(7, 1))
	}
}

func TestHandleHorizontal(t *testing.T) {

	sort.Slice(TEST_WORDS[:], func(i, j int) bool {
		return len(TEST_WORDS[i]) > len(TEST_WORDS[j])
	})
	w := len(TEST_WORDS[0]) // width
	table := make([][]string, 0, w)
	for i := 0; i < w; i++ {
		line := make([]string, w)
		table = append(table, line)
	}

	for _, v := range TEST_WORDS[0:] {

		_, err := handleHorizontal(v, table)
		if err != nil {
			t.Log(err)
		}

	}
	for i, v := range table {
		t.Logf("Row %d: %v", i, v)
	}

}

func TestHandleVertical(t *testing.T) {
	sort.Slice(TEST_WORDS[:], func(i, j int) bool {
		return len(TEST_WORDS[i]) > len(TEST_WORDS[j])
	})
	w := len(TEST_WORDS[0]) // width
	table := make([][]string, 0, w)
	for i := 0; i < w; i++ {
		line := make([]string, w)
		table = append(table, line)
	}

	// need to populate the rows to check
	for _, v := range TEST_WORDS[7:15] {

		_, err := handleHorizontal(v, table)
		if err != nil {
			t.Log(err)
		}
	}
	for _, v := range TEST_WORDS[0:] {

		_, err := handleVertical(v, table)
		if err != nil {
			t.Log(err)
		}
	}

	for i := 0; i < len(table); i++ {
		for j := 0; j < len(table[i]); j++ {
			if len(table[i][j]) == 0 {
				table[i][j] = "_"
			}
		}
	}
	for i, v := range table {
		t.Logf("%v - Row: %d", v, i)
	}

}

func TestHandleDiagonalDown(t *testing.T) {
	sort.Slice(TEST_WORDS[:], func(i, j int) bool {
		return len(TEST_WORDS[i]) > len(TEST_WORDS[j])
	})
	w := len(TEST_WORDS[0]) // width
	table := make([][]string, 0, w)
	for i := 0; i < w; i++ {
		line := make([]string, w)
		table = append(table, line)
	}

	// unit test
	word := TEST_WORDS[1]
	r := generateRandom(int64(len(table))-int64(len(word)-1), 0) // 17 - 15 => (2, 0) => 1 or 0
	c := generateRandom(int64(len(table))-int64(len(word)-1), 0)
	t.Logf("Random row index: %v", r)
	t.Logf("Random col index: %v", c)

	// allowed?
	// test data
	// need to populate the rows to check
	for _, v := range TEST_WORDS[19:] {

		_, err := handleHorizontal(v, table)
		if err != nil {
			t.Log(err)
		}
	}
	cur := make([]string, 0, len(word))
	y := int(c)
	for i := int(r); i < (len(word)-1)+int(r); i++ {
		cur = append(cur, table[i][y])
		y++

	}
	t.Logf("Confirm c val: %v", c)
	t.Logf("Word: %v", word)
	t.Logf("Control row: %v", len(cur))
	t.Logf("Is %s Allowed? %v", word, isAllowed(word, cur))

	// big test
	for _, v := range TEST_WORDS[0:] {

		_, err := handleDiagonalDown(v, table)
		if err != nil {
			t.Log(err)
		}
	}

	for i := 0; i < len(table); i++ {
		for j := 0; j < len(table[i]); j++ {
			if len(table[i][j]) == 0 {
				table[i][j] = "_"
			}
		}
	}
	for i, v := range table {
		t.Logf("%v - Row: %d", v, i)
	}

}

func TestHandleDiagonalUp(t *testing.T) {
	sort.Slice(TEST_WORDS[:], func(i, j int) bool {
		return len(TEST_WORDS[i]) > len(TEST_WORDS[j])
	})
	w := len(TEST_WORDS[0]) // width
	table := make([][]string, 0, w)
	for i := 0; i < w; i++ {
		line := make([]string, w)
		table = append(table, line)
	}

	for _, v := range TEST_WORDS[17:] {

		_, err := handleHorizontal(v, table)
		if err != nil {
			t.Log(err)
		}
	}

	for _, v := range TEST_WORDS[0:] {

		_, err := handleDiagonalUp(v, table)
		if err != nil {
			t.Log(err)
		}
	}

	for i := 0; i < len(table); i++ {
		for j := 0; j < len(table[i]); j++ {
			if len(table[i][j]) == 0 {
				table[i][j] = "_"
			}
		}
	}
	for i, v := range table {
		t.Logf("%v - Row: %d", v, i)
	}

}

func TestIsAllowed(t *testing.T) {

	control := make([]string, len(TEST_WORDS[0]))
	control[2] = string(TEST_WORDS[0][2])
	control[6] = string(TEST_WORDS[0][6])

	p1 := TEST_WORDS[0]
	if isAllowed(p1, control) {
		t.Logf("Test case 1: %s vs control: %s passed successfully", p1, control)
	}
	f1 := "SnoopDog"
	if !isAllowed(f1, control) {
		t.Logf("Test case 2: %s vs control: %s failed successfully", f1, control)
	}
}

func TestReverser(t *testing.T) {

	for _, v := range TEST_WORDS {
		t.Logf("Reverse of %s = %s", v, reverser(v))
		if v[len(v)-1] != reverser(v)[0] {
			t.Errorf("First letter != last leter: %s: %s", string(v[0]), string(reverser(v)[0]))
		}
	}
}

func TestInsertWordInPuzzle(t *testing.T) {
	sort.Slice(TEST_WORDS[:], func(i, j int) bool {
		return len(TEST_WORDS[i]) > len(TEST_WORDS[j])
	})
	w := len(TEST_WORDS[0]) // width
	table := make([][]string, 0, w)
	for i := 0; i < w; i++ {
		line := make([]string, w)
		table = append(table, line)
	}

	for _, v := range TEST_WORDS {
		insertWordInPuzzle(v, table)

	}

	sort.Strings(TEST_WORDS[:])
	for _, v := range TEST_WORDS {
		t.Log(v)
	}
	for i := 0; i < len(table); i++ {
		for j := 0; j < len(table[i]); j++ {
			if len(table[i][j]) == 0 {
				table[i][j] = "_"
			}
		}
	}
	for i, v := range table {
		t.Logf("%v - Row: %d", v, i)
	}
}

func TestRegex(t *testing.T) {
	pass := "aTOMic"
	fail := "12;-fail's test"

	ex := "^[a-zA-Z]+$"
	r, _ := regexp.Compile(ex)

	if !r.MatchString(pass) {
		t.Logf("test val %s does not satisfy Regex %s", pass, ex)
	}

	if r.MatchString(fail) {
		t.Logf("test val %s did not successfully fail/trip Regex %s", pass, ex)
	}

}

func TestGetValidPuzzleWord(t *testing.T) {

	rgx := []string{"atomic", "fail!", "1-2-3-fail!!", "You blew it!"}
	ex := "^[a-zA-Z]+$"
	r, _ := regexp.Compile(ex)

	word := getValidPuzzleWord(rgx, r)
	t.Logf("%s must equal 'atomic'", word)
}
