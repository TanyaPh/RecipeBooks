package main

import (
	"day01/internal/FSComparator"
	"flag"
)

func main() {
	var oldFile, newFile string
	flag.StringVar(&oldFile, "old", "", "old snapshot.txt")
	flag.StringVar(&newFile, "new", "", "new snapshot.txt")
	flag.Parse()
	
	data := make(map[string]bool)
	FSComparator.ReadOldTXT(oldFile, data)
	FSComparator.CheckNewTXT(newFile, data)
}
