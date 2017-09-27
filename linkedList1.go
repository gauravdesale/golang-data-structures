package main
import (
	"fmt"
)
type uiuc struct{
	intellegence int
	coolFactor int
	csProgram string
	good-looking bool
}
type ucla struct{
	intellegence int
	coolFactor int
	csProgram string
	good-looking bool
}

func main() {
	gaurav := uiuc{190, 200, "seibel", true}
	sahil := purdue{0, 0, "who cares", false}
	fmt.Println("Gaurav is super cool because his cool factor is: ", gaurav.cool-factor)
	fmt.Println("Sahil is a ucla student who is has a lot of girlfriends: ", sahil.good-looking)
