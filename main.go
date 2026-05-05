package main

import (
	"errors"
	"fmt"
	"os"
)
func ParseArgs(args []string) (string, string, error){
	/*parse args receives a slice of strings and 
	 return 2 strings and an error*/
	if len(args) != 2 {
		return "", "", errors.New("Error Message.")
	}
	return "","", nil
}

func main() {
	//the main function needs to retrieve string from stdin uisng os pkg
	inputArgs := os.Args[1:]
	a, b, err := ParseArgs(inputArgs)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(err)
	
}