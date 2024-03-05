package DBReader

import (
	"encoding/xml"
	"fmt"
	// "io/ioutil"
	"os"
)

type readFromXML struct {
	BookName string
}

func (baker1 *readFromXML) Read() RecipeBook {
	var origin RecipeBook
	err := xml.Unmarshal(ReadFile(baker1.BookName), &origin)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	// fmt.Println(origin)
	return origin
}

// func ReadFile(Path string) []byte {
// 	file, err := os.Open(Path)
// 	if err != nil {
// 		fmt.Println(err)
// 		os.Exit(0)
// 	}
// 	defer file.Close()
// 	data, err := ioutil.ReadAll(file)
// 	if err != nil {
// 		fmt.Println(err)
// 		os.Exit(0)
// 	}
// 	return data
// }

// func main() {
// 	var reader DBReader
// 	reader = &readFromXML {BookName: "original_database.xml"}

// 	recipeBook := reader.Read()
// 	fmt.Println(recipeBook)

// 	resJSON := recipeBook.WriteToJSON()
// 	fmt.Println(string(resJSON))

// 	resXML := recipeBook.WriteToXML()
// 	fmt.Println(string(resXML))
// }
