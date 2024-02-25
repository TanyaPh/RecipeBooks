package DBReader

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type readFromJSON struct {
	BookName string
}

func (baker2 *readFromJSON) Read() RecipeBook {
	var stolen RecipeBook
	err := json.Unmarshal(ReadFile(baker2.BookName), &stolen)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	fmt.Println(stolen)
	return stolen
}

func ReadFile(Path string) []byte {
	file, err := os.Open(Path)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	return data
}

// func main() {
// 	var reader DBReader
// 	reader = &readFromJSON {BookName: "stolen_database.json"}

// 	recipeBook := reader.Read()
// 	fmt.Println(recipeBook)

// 	resJSON := recipeBook.WriteToJSON()
// 	fmt.Println(string(resJSON))

// 	resXML := recipeBook.WriteToXML()
// 	fmt.Println(string(resXML))
// }
