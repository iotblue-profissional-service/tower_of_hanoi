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
func (s *Stack) push(size Disk) {
	s.items = append(s.items, size)
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
		fmt.Println(item)
	}
	fmt.Println()
}

// the problem
func tower_of_hanoi(n int, s, a, e *Stack) {
	// Todo: add print statements to be able to see code progress
	if n == 1 {
		e.push(s.top())
		s.pop()
	} else {
		tower_of_hanoi(n-1, s, e, a)
		e.push(s.top())
		s.pop()
		tower_of_hanoi(n-1, a, s, e)
	}
}

// main function
func main() {
	s := Stack{}
	a := Stack{}
	e := Stack{}

	var x int
	fmt.Println("Enter the number of disks:")
	fmt.Scan(&x)

	if x == 3 || x == 5 || x == 7 {
		for i := x; i > 0; i-- {
			s.push(Disk{size: i})
		}
	} else {
		fmt.Println("wrong value, enter either 3, 5 or 7: ")
		fmt.Scan(&x)
		for i := x; i > 0; i-- {
			s.push(Disk{size: i})
		}
	}
	tower_of_hanoi(x, &s, &a, &e)

	fmt.Println("Source:")
	s.Print()
	fmt.Println("Destination:")
	e.Print()
}
