package main

import "fmt"

func main() {

	fmt.Printf("%.2f litres needed", calcPaint(4.2, 5.6))
}

func calcPaint(width float64, height float64) float64 {
	area := width * height
	return area / 10t
}
