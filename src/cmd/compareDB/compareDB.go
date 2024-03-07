package main

import (
	"day01/internal/DBComparer"
	"day01/internal/DBReader"
	"flag"
)

func main() {
	var oldFile, newFile string
	flag.StringVar(&oldFile, "old", "", "path to file")
	flag.StringVar(&newFile, "new", "", "path to file")
	flag.Parse()

	reader := DBReader.New(oldFile)
	oldRecipeBook := reader.Read()

	reader = DBReader.New(newFile)
	newRecipeBook := reader.Read()

	DBComparer.CheckCakes(&oldRecipeBook, &newRecipeBook)
}
