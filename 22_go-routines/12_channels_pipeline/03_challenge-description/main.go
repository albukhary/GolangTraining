//package main
//
//import (
//	"fmt"
//)
//
//func main() {
//	c := factorial(4,5,2,6,7)
//	for n := range c {
//		fmt.Println(n)
//	}
//}
//
//func factorial(numbers ...int) chan int {
//	out := make(chan int)
//
//		go func() {
//			for n := range numbers {
//				total := 1
//				for i := n; i > 0; i-- {
//					total *= i
//				}
//				out <- total
//			}
//			close(out)
//		}()
//
//	return out
//}

/*
CHALLENGE #1:
-- Change the code above to execute 100 factorial computations concurrently and in parallel.
-- Use the "pipeline" pattern to accomplish this

POST TO DISCUSSION:
-- What realizations did you have about working with concurrent code when building your solution?
-- eg, what insights did you have which helped you understand working with concurrency?
-- Post your answer, and read other answers, here: https://goo.gl/uJa99G
*/
package main

import "fmt"

func main(){
	palindrome(123454321)
}
func palindrome(n int) bool {
	//12321
	var digits []int
	i:=0
	for n>0 {
		digits= append(digits, n%10)
		n=n/10
		i++
	}
	length := len(digits)
	fmt.Println(digits)

	if length%2 ==0{
		fmt.Println("It is not a palindrome")
		return false
	}

	for i:=0;i<length/2;i++{
		if digits[i]==digits[length-1-i]{
			fmt.Println("It is a palindrome !")
			return true
		}
	}
	return false
}
/*

PostGres
Gin
Swagger
Redis

Hammasini ishlatib CRUD

 */