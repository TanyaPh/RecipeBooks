package main

import (
	"day01/internal/DBReader"
	"day01/internal/DBComparator"
	"flag"
)

func main() {
	// check Args
	var oldFile, newFile string
	flag.StringVar(&oldFile, "old", "", "a string var")
	flag.StringVar(&newFile, "new", "", "a string var")
	flag.Parse()

	reader := DBReader.New(oldFile)
	oldRecipeBook := reader.Read()

	reader = DBReader.New(newFile)
	newRecipeBook := reader.Read()

	DBComparator.CheckCakes(&oldRecipeBook, &newRecipeBook)
}
