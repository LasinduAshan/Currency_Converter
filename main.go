package main

import "fmt"

func main() {

	var lkrRuppes float32
	var usdAmount float32

	fmt.Print("Enter LKR Amount : ")
	fmt.Scanf("%f", &usdAmount)
	fmt.Printf("LKR Amount : %f  \n", usdAmount * 0.0051)

	fmt.Print("Enter USD Amount : ")
	fmt.Scanf("%f", &lkrRuppes)
	fmt.Printf("USD Amount : %f  \n", lkrRuppes * 187.99)
}
