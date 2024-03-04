package main

import (
    "fmt"
	"flag"
	"day01/internal/DBReader"
)

// func checkUnit() {
	
// }

func checkCakes(old *DBReader.RecipeBook, new *DBReader.RecipeBook) {
	same := false
	for _, new_cake := range new.Recipes {
        for _, old_cake := range old.Recipes {
			if new_cake.Name == old_cake.Name {
				checkTime(old_cake.Name, old_cake.Time, new_cake.Time)
				checkIngredients(&old_cake, &new_cake)
				same = true
				break
			}
		}
		if !same {
			fmt.Printf("ADDED cake %q\n", new_cake.Name)
		} else {
			same = false
		}
    }

	for _, old_cake := range old.Recipes {
        for _, new_cake := range new.Recipes {
			if old_cake.Name == new_cake.Name {
				same = true
				break
			}
		}
		if !same {
			fmt.Printf("REMOVED cake %q\n", old_cake.Name)
		} else {
			same = false
		}
    }
}

func checkTime(cakeName string, oldTime string, newTime string) {
	if oldTime != newTime {
		fmt.Printf("CHANGED cooking time for cake %q - %q instead of %q\n", cakeName, newTime, oldTime)
    }
}

func checkIngredients(old *DBReader.Cake, new *DBReader.Cake) {
	same := false
	for _, new_item := range new.Ingredients {
        for _, old_item := range old.Ingredients {
			if new_item.Name == old_item.Name {
				checkUnit(old.Name, &old_item, &new_item)
				same = true
				break
			}
		}
		if !same {
			fmt.Printf("ADDED ingredient %q for cake %q\n", new_item.Name, new.Name)
		} else {
			same = false
		}
    }

	for _, old_item := range old.Ingredients {
        for _, new_item := range new.Ingredients {
			if old_item.Name == new_item.Name {
				same = true
				break
			}
		}
		if !same {
			fmt.Printf("REMOVED ingredient %q for cake %q\n", old_item.Name, old.Name)
		} else {
			same = false
		}
    }
}


func checkUnit(cakeName string, old *DBReader.Item, new *DBReader.Item) {
	if old.Unit != new.Unit {
		switch {
		case old.Unit == "":
				fmt.Printf("ADDED unit %q for ingredient %q for cake  %q\n", new.Unit, new.Name, cakeName)
		case new.Unit == "":
				fmt.Printf("ADDED unit %q for ingredient %q for cake  %q\n", old.Unit, new.Name, cakeName)
		default:
			fmt.Printf("CHANGED unit for ingredient %q for cake  %q - %q instead of %q\n", new.Name, cakeName, new.Unit, old.Unit)
		}
    }
	if old.Count != new.Count {
		fmt.Printf("CHANGED unit count for ingredient %q for cake  %q - %q instead of %q\n", new.Name, cakeName, new.Count, old.Count)
    }
}

func main() {
	// check Args
	var oldFile string
    flag.StringVar(&oldFile, "old", "", "a string var")
	var newFile string
    flag.StringVar(&newFile, "new", "", "a string var")
	flag.Parse()
	// fmt.Println("old file:", oldFile)
	// fmt.Println("new file:", newFile)

	// var reader DBReader
	// reader = &readFromXML {BookName: oldFile}
	reader := DBReader.New(oldFile)
	oldRecipeBook := reader.Read()

	reader = DBReader.New(newFile)
	newRecipeBook := reader.Read()

	checkCakes(&oldRecipeBook, &newRecipeBook)
}
