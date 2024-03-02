package main

import (
    "fmt"
	"flag"
	"day01/internal/DBReader"
)

// func checkTime(compareCakes *[]string, old *RecipeBook, new *RecipeBook) {
// 	// for _, real_cake := range compareCakes  {

// 	// }
// 	for _, new_cake := range new.Recipes {
//         for _, old_cake := range old.Recipes {
// 			if new_cake.Name != old_cake.Name {
// 				continue
// 			}
// 			if new_cake.Time != old_cake.Time {
// 				fmt.Printf("CHANGED cooking time for cake %q - %q instead of %q\n", new_cake.Name, old_cake.Time, new_cake.Time)
// 			}
// 			checkIngredients(&new_cake.Ingredients, &old_cake.Ingredients)
// 		}
//     }
// }

// func checkIngredients(new *[]Item, old *[]Item) {
// 	for _, new_item := range new {
//         for _, old_item := range old {
// 			if new_item.Name == old_item.Name {
// 				// sameCake = append(sameCake, old_cake.Name)
// 				continue
// 			}
// 			fmt.Printf("ADDED ingredient %q for cake %q\n", new_item.Name, new_cake.Name)
// 		}
//     }

// 	for _, old_cake := range old.Recipes {
//         for _, real_cake := range sameCake {
// 			if old_cake.Name == real_cake {
// 				continue
// 			}
// 			fmt.Printf("REMOVED cake %q\n", old_cake.Name)
// 		}
//     }
// }

// func checkCount() {
	
// }

// func checkUnit() {
	
// }

func checkCakes(old RecipeBook, new RecipeBook) []string {
	var sameCake []string
	for _, new_cake := range new.Recipes {
        for _, old_cake := range old.Recipes {
			if new_cake.Name == old_cake.Name {
				sameCake = append(sameCake, old_cake.Name)
				break
			}
		}
		fmt.Printf("ADDED cake %q\n", new_cake.Name)
    }

	for _, old_cake := range old.Recipes {
        for _, real_cake := range sameCake {
			if old_cake.Name == real_cake {
				continue
			}
			fmt.Printf("REMOVED cake %q\n", old_cake.Name)
		}
    }
	return sameCake
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

	compareCakes := checkCakes(oldRecipeBook, newRecipeBook)

}
