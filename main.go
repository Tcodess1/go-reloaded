package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)
var errorMessage error = errors.New("Error: Please provide exactly one input file and one output file")
func ParseArgs(args []string) ( inputFile string, outputFile string, err error){
	/*parse args receives a slice of strings and 
	 return 2 strings and an error*/
	if len(args) != 2 {
		return "", "", errorMessage
	}
	return args[0],args[1], nil
}

func main() {
	//the main function needs to retrieve string from stdin uisng os pkg
	inputArgs := os.Args[1:]
	inputFile, outputFile, err := ParseArgs(inputArgs)
	if err != nil {
		log.Fatalf("%v", err)
	}
	fmt.Println(inputFile)
	fmt.Println(outputFile)
	fmt.Println(err)
	
}