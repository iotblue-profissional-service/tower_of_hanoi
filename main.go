package main

import (
	"fmt"
)

var rod = []byte{'1', '2', '3'}
var stacks = make([][]int, 3)

func Push(stack *[]int, item int) {
	*stack = append(*stack, item)
}

func Pop(stack *[]int) int {
	if len(*stack) > 0 {
		index := len(*stack) - 1
		item := (*stack)[index]
		*stack = (*stack)[:index]
		return item
	}
	return -1 // Return -1 for an empty stack
}

func moveDisk(a, b int) {
	if len(stacks[b]) == 0 || (len(stacks[a]) > 0 && stacks[a][len(stacks[a])-1] < stacks[b][len(stacks[b])-1]) {
		disk := Pop(&stacks[a])
		fmt.Printf("Move disk %d from rod %c to rod %c\n", disk, rod[a], rod[b])
		Push(&stacks[b], disk)
	} else {
		moveDisk(b, a)
	}
}

func towerOfHanoi(n int) {
	fmt.Printf("Tower of Hanoi for %d disks:\n", n)

	src, aux, dest := 0, 1, 2
	for i := n; i > 0; i-- {
		Push(&stacks[src], i)
	}

	totalMoves := (1 << n) - 1
	if n%2 == 0 {
		aux, dest = dest, aux
	}

	for i := 1; i <= totalMoves; i++ {
		if i%3 == 0 {
			moveDisk(aux, dest)
		} else if i%3 == 1 {
			moveDisk(src, dest)
		} else {
			moveDisk(src, aux)
		}
	}
}

func main() {
	n := 3 // number of disks
	towerOfHanoi(n)
}
