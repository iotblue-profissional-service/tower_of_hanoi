// package main
//
// import "fmt"
//
// type Stack []int
//
//	func (s *Stack) push(item int) {
//		*s = append(*s, item)
//	}
//
//	func (s *Stack) pop() int {
//		if len(*s) == 0 {
//			return -1
//		}
//		top := (*s)[len(*s)-1]
//		*s = (*s)[:len(*s)-1]
//		return top
//	}
//
//	func towerOfHanoi(n int) {
//		fromRod := Stack{}
//		remRod := Stack{}
//		toRod := Stack{}
//
//		for i := n; i >= 1; i-- {
//			fromRod.push(i)
//		}
//
//		if n%2 == 0 {
//			remRod, toRod = toRod, remRod
//		}
//
//		totalMoves := (1 << uint(n)) - 1
//
//		for move := 1; move <= totalMoves; move++ {
//			if move%3 == 1 {
//				moveDisk(&fromRod, &toRod)
//			} else if move%3 == 2 {
//				moveDisk(&fromRod, &remRod)
//			} else {
//				moveDisk(&remRod, &toRod)
//			}
//		}
//
// }
//
//	func moveDisk(from, to *Stack) {
//		if from == nil || to == nil {
//			return
//		}
//		disk := from.pop()
//		if disk != -1 {
//			to.push(disk)
//			fmt.Printf("Move disk %d from source to target", disk)
//		}
//	}
//
//	func main() {
//		// start
//		numDisks := 3
//		fmt.Printf("The number of disks is %d , Solve the problem with steps", numDisks)
//		towerOfHanoi(numDisks)
//	}
package main

import (
	"fmt"
)

func moveDisks(n int, source, aux, target *Stack) {
	if n > 0 {
		// Move n-1 disks from source to aux using target as the auxiliary stack
		moveDisks(n-1, source, target, aux)

		// Move the nth disk from source to target
		target.push(source.pop())
		fmt.Printf("Move disk %v from stick %v to stick %v \n", n, source.id, target.id)
		// Move the n-1 disks from aux to target using source as the auxiliary stack
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
		return -1 // Stack is empty
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
