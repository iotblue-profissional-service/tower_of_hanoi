package main

import (
	"fmt"
)

var rod = []byte{'1', '2', '3'}
var stacks = make([][]int, 3)

func moveDisk(a, b int) {
	if len(stacks[b]) == 0 || (len(stacks[a]) > 0 && stacks[a][len(stacks[a])-1] < stacks[b][len(stacks[b])-1]) {
		fmt.Printf("Move disk %d from rod %c to rod %c\n", stacks[a][len(stacks[a])-1], rod[a], rod[b])
		stacks[b] = append(stacks[b], stacks[a][len(stacks[a])-1])
		stacks[a] = stacks[a][:len(stacks[a])-1]
	} else {
		moveDisk(b, a)
	}
}

func towerOfHanoi(n int) {
	fmt.Printf("Tower of Hanoi for %d disks:\n", n)

	src, aux, dest := 0, 1, 2
	for i := n; i > 0; i-- {
		stacks[src] = append(stacks[src], i)
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
