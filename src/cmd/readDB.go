package main

import (
    "fmt"
	"os"
	"DBReader"
)

func main() {
	// if len(os.Args) < 3 {
    //     fmt.Println("no file specified")
	// 	os.Exit(0)
    // }
	// filePath := os.Args[2]

	var file string
    flag.StringVar(&file, "f", "", "a string var")
	flag.Parse()

	reader := DBReader.New(file)

	recipeBook := reader.Read()
	fmt.Println(recipeBook)

	resJSON := recipeBook.WriteToJSON()
	fmt.Println(string(resJSON))

	resXML := recipeBook.WriteToXML()
	fmt.Println(string(resXML))
}
