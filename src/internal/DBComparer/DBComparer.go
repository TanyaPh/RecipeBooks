package DBComparer

import (
	"day01/internal/DBReader"
	"fmt"
)

func CheckCakes(oldBook *DBReader.RecipeBook, newBook *DBReader.RecipeBook) {
	oldMap := make(map[string]DBReader.Cake)
	for _, oldCake := range oldBook.Recipes {
		oldMap[oldCake.Name] = oldCake
	}

	newMap := make(map[string]DBReader.Cake)
	for _, newCake := range newBook.Recipes {
		newMap[newCake.Name] = newCake
	}

	var sameCake []string
	for name, newCake := range newMap {
		_, ok := oldMap[name]
		if !ok {
			fmt.Printf("ADDED cake %q\n", newCake.Name)
			continue
		}
		sameCake = append(sameCake, name)
	}

	for name, oldCake := range oldMap {
		if _, ok := newMap[name]; !ok {
			fmt.Printf("REMOVED cake %q\n", oldCake.Name)
		}
	}

	for _, cakeName := range sameCake {
		CheckTime(cakeName, oldMap[cakeName].Time, newMap[cakeName].Time)
		CheckIngredients(cakeName, oldMap[cakeName], newMap[cakeName])
	}
}

func CheckTime(cakeName string, oldTime string, newTime string) {
	if oldTime != newTime {
		fmt.Printf("CHANGED cooking time for cake %q - %q instead of %q\n", cakeName, newTime, oldTime)
	}
}

func CheckIngredients(cakeName string, oldCake DBReader.Cake, newCake DBReader.Cake) {
	oldList := make(map[string]DBReader.Item)
	for _, oldItem := range oldCake.Ingredients {
		oldList[oldItem.Name] = oldItem
	}

	newList := make(map[string]DBReader.Item)
	for _, newItem := range newCake.Ingredients {
		newList[newItem.Name] = newItem
	}

	var sameItems []string
	for name := range newList {
		_, ok := oldList[name]
		if !ok {
			fmt.Printf("ADDED ingredient %q for cake %q\n", name, cakeName)
			continue
		}
		sameItems = append(sameItems, name)
	}

	for name := range oldList {
		if _, ok := newList[name]; !ok {
			fmt.Printf("REMOVED ingredient %q for cake %q\n", name, cakeName)
		}
	}

	for _, itemName := range sameItems {
		CheckUnit(cakeName, oldList[itemName], newList[itemName])
	}
}

func CheckUnit(cakeName string, old DBReader.Item, new DBReader.Item) {
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

func CheckCount(cakeName string, old DBReader.Item, new DBReader.Item) {
	if old.Count != new.Count {
		fmt.Printf("CHANGED unit count for ingredient %q for cake  %q - %q instead of %q\n", new.Name, cakeName, new.Count, old.Count)
	}
}
