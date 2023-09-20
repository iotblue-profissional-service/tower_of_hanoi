package main

import "fmt"

// structs
type Disk struct {
	size int
}
type Stack struct {
	items []Disk
}

// stack basic functions
func (s *Stack) push(data Disk) {
	s.items = append(s.items, data)
}
func (s *Stack) IsEmpty() bool {
	if len(s.items) == 0 {
		return true
	}
	return false
}
func (s *Stack) pop() {
	if s.IsEmpty() {
		return
	}
	s.items = s.items[:len(s.items)-1]
}
func (s *Stack) top() Disk {
	if s.IsEmpty() {
		return Disk{}
	}
	return s.items[len(s.items)-1]
}
func (s *Stack) Print() {
	for _, item := range s.items {
		fmt.Print(item, " ")
	}
	fmt.Println()
}

// the problem
func tower_of_hanoi(n int, s Stack, a Stack, e Stack) {
	if n == 1 {
		e.push(s.top())
		s.pop()
	} else {
		tower_of_hanoi(n-1, s, a, e)
		e.push(s.top())
		s.pop()
		tower_of_hanoi(n-1, a, e, s)
	}
}

// main function
func main() {
	s := Stack{}
	a := Stack{}
	e := Stack{}

	s.push(Disk{size: 3})
	s.push(Disk{size: 2})
	s.push(Disk{size: 1})

	var x int
	fmt.Println("enter the number of disks: ")
	fmt.Scan(&x)
	tower_of_hanoi(x, s, a, e)

	//s.Print()
	fmt.Println("---------------------------")
	e.Print()
}
