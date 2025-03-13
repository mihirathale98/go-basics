package main

import ("fmt")

func add(x float64, y float64) float64 {
	return x+y 
}

func multiple(a, b string) (string, string) {
	return a, b
}

func main() {
	var num1, num2 float64 = 5.6, 9.5

	fmt.Println(add(num1, num2))

	w1, w2 := "Hey", "there"

	
	fmt.Println(multiple(w1, w2))
}





