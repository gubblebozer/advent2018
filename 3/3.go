package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"regexp"
	"strconv"
//	"container/list"
)

var WID = 1000
var HGT = 1000

type claimtype struct {
	claim int
	x int
	y int
	width int
	height int
}

func check(e error) {
	if e != nil { panic(e) }
}

func re_named_capture(re *regexp.Regexp, line string) map[string]string {
	match := re.FindStringSubmatch(line)
	capmap := make(map[string]string)
	if len(match) == 0 { return capmap }
	for i, name := range re.SubexpNames() {
		if i != 0 && name != "" {
			capmap[name] = match[i]
		}
	}
	return capmap
}

func dump(conflicts [][][]int) {
	for x := 0; x < WID; x++ {
		for y := 0; y < HGT; y++ {
			if len(conflicts[x][y]) > 0 {
				fmt.Printf("x:%d y:%d ", x, y)
				for _, v := range conflicts[x][y] {
					fmt.Printf("%d ", v)
				}
				fmt.Printf("\n")
			}
		}
	}
}

func square_inches_in_n(conflicts [][][]int, n int) int {
	sqin := 0

	for x := 0; x < WID; x++ {
		for y := 0; y < HGT; y++ {
			if len(conflicts[x][y]) >= n {
				sqin++
			}
		}
	}

	return sqin
}

func conflict_set(conflicts [][][]int, c claimtype) {
	for x := c.x; x < c.x + c.width; x++ {
		for y := c.y; y < c.y + c.height; y++ {
			conflicts[x][y] = append(conflicts[x][y], c.claim)
		}
	}
}

func claim_has_conflicts(conflicts [][][]int, c claimtype) bool {
	for x := c.x; x < c.x + c.width; x++ {
		for y := c.y; y < c.y + c.height; y++ {
			if len(conflicts[x][y]) > 1 {
				return true
			}
		}
	}
	return false
}

// #1337 @ 775,762: 10x18
// #<claim> @ <x>,<y>: <width>x<height>
func line_parse(line string, conflicts [][][]int) (claimtype, bool) {
	c := claimtype{}
	re := regexp.MustCompile(`\#(?P<claim>\d+)\s*\@\s*(?P<x>\d+),(?P<y>\d+)\s*:\s*(?P<width>\d+)x(?P<height>\d+)`)
	m := re_named_capture(re, line)
	if len(m) > 0 {
		claim, err := strconv.Atoi(m["claim"])
		if err != nil { return c, false }
		x, err := strconv.Atoi(m["x"])
		if err != nil { return c, false }
		y, err := strconv.Atoi(m["y"])
		if err != nil { return c, false }
		width, err := strconv.Atoi(m["width"])
		if err != nil { return c, false }
		height, err := strconv.Atoi(m["height"])
		if err != nil { return c, false }
		c := claimtype { x: x, y: y, width: width, height: height, claim: claim }
		conflict_set(conflicts, c)
		fmt.Printf("claim:%d x:%d y:%d width:%d height:%d\n", c.claim, c.x, c.y, c.width, c.height)
		return c, true
	}
	return c, false
}

func main() {
	conflicts := make([][][]int, WID)
	for x := range conflicts {
		conflicts[x] = make([][]int, HGT)
		for y := range conflicts[x] {
			conflicts[x][y] = make([]int, 0)
		}
	}
	
	dat, err := ioutil.ReadAll(os.Stdin)
	check(err)

	lines := strings.Split(string(dat), "\n")
	
	claims := make([]claimtype, 0)
	for _, v := range lines {
		c, ok := line_parse(v, conflicts)
		if ok {
			claims = append(claims, c)
		}
	}

	sqin := square_inches_in_n(conflicts, 2)

	fmt.Printf("square inches with conflicts: %d\n", sqin)

	fmt.Printf("claims: %d\n", len(claims))
	for _, c := range claims {
	  	if !claim_has_conflicts(conflicts, c) {
			fmt.Printf("claim %d has no conflicts\n", c.claim)
		}
	}
}
