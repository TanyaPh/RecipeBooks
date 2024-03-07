package main

import (
	"day01/internal/DBReader"
	"flag"
	"fmt"
)

func main() {
	var file string
	flag.StringVar(&file, "f", "", "path to file")
	flag.Parse()

	reader := DBReader.New(file)

	recipeBook := reader.Read()
	fmt.Println(recipeBook)

	resJSON := recipeBook.WriteToJSON()
	fmt.Println(string(resJSON))

	resXML := recipeBook.WriteToXML()
	fmt.Println(string(resXML))
}
