package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"strings"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	var x int = 0

	var reached bool
	var hits map[int]int
	hits = make(map[int]int)

	dat, err := ioutil.ReadAll(os.Stdin)
	check(err)

	s := strings.Split(string(dat), "\n")

	for !reached {
		for _, v := range s {
			y, err := strconv.Atoi(v)
			if err == nil {
				x += y
				_, present := hits[x]
				if present {
					hits[x]++
					reached = true
					break
				} else {
					hits[x] = 1
				}
			}
		}
		fmt.Println("x: ", x)
	}

	fmt.Println("first reached: ", x)
}
