package main

import (
	"fmt"

	"github.com/hassanjawwad12/betterj"
)

func main() {

	fmt.Println("This is a sample program to use the JSON Minify and Beautify module")
	fmt.Printf("\n")
	jsonData := `{
		"name": "Hassan Jawwad",
		"age": 22,
		"city": "Lahore"
	}`
	response, err := betterj.MinifyJ(jsonData)
	if err != nil {
		fmt.Println("Error has occured", err)
		return
	}
	fmt.Println("The output json after Minify is: ", response)

	fmt.Printf("\n")

	res, err := betterj.BeautifyJ(jsonData, " ")
	if err != nil {
		fmt.Println("Error has occured", err)
		return
	}
	fmt.Println("The output json after Beautify is: ", res)

}
