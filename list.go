package main

import (
  "strings"
  "fmt"
)

type List[T comparable] struct {
	next *List[T]
	val  T
}

func (l *List[T]) String() string {
	items := []string {}
	
	
	for curr := l; curr != nil; curr = curr.next {
		items = append(items, fmt.Sprint(curr.val))
	}
	
	str := strings.Join(items, " -> ")
	
	return str
}


func main() {
	l := &List[int] { val: 1, next: &List[int] { val: 3, next: &List[int] { val: 5, next: nil }}}
	fmt.Println(l)
}

