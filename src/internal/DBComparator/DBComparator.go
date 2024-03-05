package DBComparator

import (
	"day01/internal/DBReader"
	"fmt"
)

func CheckCakes(old *DBReader.RecipeBook, new *DBReader.RecipeBook) {
	same := false
	for _, new_cake := range new.Recipes {
		for _, old_cake := range old.Recipes {
			if new_cake.Name == old_cake.Name {
				CheckTime(old_cake.Name, old_cake.Time, new_cake.Time)
				CheckIngredients(&old_cake, &new_cake)
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

func CheckTime(cakeName string, oldTime string, newTime string) {
	if oldTime != newTime {
		fmt.Printf("CHANGED cooking time for cake %q - %q instead of %q\n", cakeName, newTime, oldTime)
	}
}

func CheckIngredients(old *DBReader.Cake, new *DBReader.Cake) {
	same := false
	for _, new_item := range new.Ingredients {
		for _, old_item := range old.Ingredients {
			if new_item.Name == old_item.Name {
				CheckUnit(old.Name, &old_item, &new_item)
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

func CheckUnit(cakeName string, old *DBReader.Item, new *DBReader.Item) {
	switch {
	case old.Unit == new.Unit:
		CheckCount(cakeName, old, new)
	case old.Unit == "":
		fmt.Printf("ADDED unit %q for ingredient %q for cake  %q\n", new.Unit, new.Name, cakeName)
	case new.Unit == "":
		fmt.Printf("REMOVED unit %q for ingredient %q for cake  %q\n", old.Unit, new.Name, cakeName)
	default:
		fmt.Printf("CHANGED unit for ingredient %q for cake  %q - %q instead of %q\n", new.Name, cakeName, new.Unit, old.Unit)
	}
}

func CheckCount(cakeName string, old *DBReader.Item, new *DBReader.Item) {
	if old.Count != new.Count {
		fmt.Printf("CHANGED unit count for ingredient %q for cake  %q - %q instead of %q\n", new.Name, cakeName, new.Count, old.Count)
	}
}
