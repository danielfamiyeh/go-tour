package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
  matrix := [][] uint8{}
  
  for i := 0; i < dy; i++ {
    slice := [] uint8{}
	
	for j := 0; j < dx; j++ {
	  slice = append(slice, uint8(j * i))	
	}
	matrix = append(matrix, slice)
  }
  
  return matrix
}

func main() {
	 pic.Show(Pic)
}

