package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const bagSplit = " contain "
const containsNone = "no other bags"
const myBag = "shiny gold bag"

type Bag struct {
	color string
	count int
	bags  []*Bag
}

func (b *Bag) addChild(str string, num int) {
	b.bags = append(b.bags, &Bag{
		color: str,
		count: num,
		bags:  nil,
	})
}

func newBag(str string) *Bag {
	return &Bag{
		color: str,
		count: 1,
		bags:  nil,
	}
}

func findBag(search string, dict map[string]*Bag) int {
	ctr := 0

	tmp := dict[search]
	if tmp == nil {
		return ctr
	}

	for i := 0; i < len(tmp.bags); i++ {
		ctr += tmp.bags[i].count*findBag(tmp.bags[i].color, dict) + tmp.bags[i].count
	}

	return ctr
}

func Evaluate() int {

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("failed to open file")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	dict := make(map[string]*Bag)
	//
	// Parse all the rules
	//
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), bagSplit)

		bagName := line[0]
		bagName = bagName[:len(bagName)-1]
		rules := strings.Split(line[1], ",")

		tmp := newBag(bagName)
		for i := 0; i < len(rules); i++ {
			rule := strings.Trim(rules[i], " .")
			numOfBags, _ := strconv.Atoi(string(rule[0]))

			if rule == containsNone {
				break
			} else {
				// Takes the 's' out of the bag
				if rule[len(rule)-1] == 's' {
					tmp.addChild(rule[2:len(rule)-1], numOfBags)
				} else {
					tmp.addChild(rule[2:], numOfBags)
				}
			}
		}

		dict[bagName] = tmp

	}

	ans := findBag(myBag, dict)

	return ans
}

func main() {

	ans := Evaluate()

	fmt.Println(ans)
}
