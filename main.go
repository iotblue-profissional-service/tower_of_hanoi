package main

/*ok hena ha expain el algorithm, ana 3andi 3 towers, fa ha3mel le kol wa7da feehom stack w keda a3raf
eni a3mel push w pop zay ma ana 3ayz bzbt le ko tower w ma3 kol movement ha2ool el movement deeh 7aslet men feen le feen
*/
import (
	"container/list"
	"fmt"
)

type Stack struct {
	stack *list.List
	name  string
}

func NewStack(name string) *Stack {
	return &Stack{list.New(), name}
}

func (s *Stack) Push(value interface{}) {
	s.stack.PushBack(value)
}

func (s *Stack) Pop() interface{} {
	err := s.stack.Back()
	//law el stack empty
	if err != nil {
		s.stack.Remove(err)
		return err.Value
	}
	return nil
}

func (s *Stack) Peek() interface{} {
	err := s.stack.Back()
	if err != nil {
		return err.Value
	}
	return nil
}

func (s *Stack) Len() int {
	return s.stack.Len()
}

func (s *Stack) IsEmpty() bool {
	return s.stack.Len() == 0
}

func hanoi(n int, source, helper, target *Stack) {
	if n > 0 {
		hanoi(n-1, source, target, helper)
		target.Push(source.Pop())
		fmt.Printf("Move disk %v from rod %v to rod %v\n", n, source.name, target.name)
		hanoi(n-1, helper, source, target)
	}
}

func main() {
	numDisks := 3
	source := NewStack("1")
	target := NewStack("3")
	helper := NewStack("2")
	fmt.Printf("How many disks do you want\n Enter \n A for 3 \n B for 5 \n C for 7 \n?\n")

	var letter string
	_, err := fmt.Scanln(&letter)
	if err != nil {
		return
	}

	if letter == "A" {
		numDisks = 3
	} else if letter == "B" {
		numDisks = 5
	} else if letter == "C" {
		numDisks = 7
	} else {
		fmt.Printf("Invalid Input, Running Default case 3 :)")
	}

	for i := numDisks; i > 0; i-- {
		source.Push(i)
	}

	hanoi(numDisks, source, helper, target)
}
