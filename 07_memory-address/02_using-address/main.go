//package main
//
//import "fmt"
//
//const metersToYards float64 = 1.09361
//
//func main() {
//	var meters float64
//	fmt.Print("Enter meters swam: ")
//	fmt.Scan(&meters)
//	yards := meters * metersToYards
//	fmt.Println(meters, " meters is ", yards, " yards.")
//}
package main

import "fmt"

const metersToYards = 1.09361

func main(){
	var meters float64
	fmt.Print("Enter how many meters you ran: ")
	fmt.Scan(&meters)
	yards := meters * metersToYards
	fmt.Printf("%v meters is %.2f yards \n",meters, yards)
}
