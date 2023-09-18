package main

import (
	"fmt"
)

func moveDisks(n int, source, aux, target *Stack) {
	if n > 0 {

		moveDisks(n-1, source, target, aux)

		target.push(source.pop())
		fmt.Printf("Move disk %v from stick %v to stick %v \n", n, source.id, target.id)

		moveDisks(n-1, aux, source, target)
	}
}

type Stack struct {
	items []int
	id    string
}

func (s *Stack) push(item int) {
	s.items = append(s.items, item)
}

func (s *Stack) pop() int {
	if len(s.items) == 0 {
		return -1
	}

	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

func main() {
	var nDisks int
	_, err := fmt.Scanln(&nDisks)
	if err != nil {
		return
	}
	source := &Stack{id: "1"}
	for i := nDisks; i > 0; i-- {
		source.push(i)
	}
	aux := &Stack{id: "2"}
	target := &Stack{id: "3"}

	moveDisks(len(source.items), source, aux, target)

	fmt.Println("Source stack:", source.items)
	fmt.Println("Auxiliary stack:", aux.items)
	fmt.Println("Target stack:", target.items)
}
