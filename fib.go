package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
  vals := [] int{}
  
  return func() int {
    if len(vals) == 0 {
	  vals = append(vals, 0)
	  return 0;
	}
	
	if len(vals) == 1 || len(vals) == 2 {
	  vals = append(vals, 1)
	  return 1;
	}
	
	vals = vals[1:]
	vals = append(vals, vals[0] + vals[1])
	
	return vals[2]
  }
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

