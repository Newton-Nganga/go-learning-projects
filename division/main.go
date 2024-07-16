package main

import (
	"errors"
	"fmt"
)
//the divide function takes in y and x as float64 and 
//returns a float64 and error
func divide(x,y float64) (float64,error){
  
	//check if the denominator is zero
	if y == 0{
		return 0,errors.New("cannot divide by zero")
	}
    //if not return the values
	return x / y ,nil
}


func main() {

// we are destructuring response from the divide function
  result,err := divide(10,0)

//check if error is nil in other words null
if err != nil {
	fmt.Println("Error :",err)
	return
}

//else if there is no error
  fmt.Println("Result :",result)
}