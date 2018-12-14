package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	var x int = 0
	dat, err := ioutil.ReadAll(os.Stdin)
	check(err)

	s := strings.Split(string(dat), "\n")

	for _, v := range s {
		fmt.Println(v)
		y, err := strconv.Atoi(v)
		if err == nil {
			x += y
		}
	}

	fmt.Println("end: ", x)
}
