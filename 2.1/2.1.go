package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func check(e error) {
	if e != nil { panic(e) }
}

func proc(s string) [4]int {
	var hits [4]int

	for i := 0; i < len(s); i++ {
		num := strings.Count(s, string(s[i]))
		if num < 4 {
			hits[num] = 1
		}
	}

	return hits
}

func main() {
	twos := 0
	threes := 0
	dat, err := ioutil.ReadAll(os.Stdin)
	check(err)
	

	s := strings.Split(string(dat), "\n")

	for _, v := range s {
		hits := proc(v)
		if hits[2] > 0 { twos++ }
		if hits[3] > 0 { threes++ }
	}

	checksum := twos * threes
	fmt.Println("checksum = ", checksum)
}
