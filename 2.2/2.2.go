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

func proc(s string, dict []string) {
	for _, v := range dict {
		mismatch := 0
		pos := 0
		if len(s) != len(v) { continue }
		for i := 0; i < len(s); i++ {
			if s[i] != v[i] {
				mismatch++
				pos = i
			}
			if mismatch > 1 { break }
		}
		if mismatch == 1 {
			for i := 0; i < len(s); i++ {
				if i == pos { continue }
				fmt.Print(string(s[i]))
			}
			fmt.Println("")
		}
	}
}

func main() {
	dat, err := ioutil.ReadAll(os.Stdin)
	check(err)

	dict := strings.Split(string(dat), "\n")

	for _, v := range dict {
		proc(v, dict)
	}
}
