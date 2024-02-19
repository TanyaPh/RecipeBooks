package main

import (
	"encoding/json"
	"encoding/xml"
    "fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// type Ingredient struct {
// 	XMLName xml.Name 	`xml:"item"`
// 	Name string			`xml:"itemname"`		`json:"ingredient_name"`
// 	Count string		`xml:"itemcount"`		`json:"ingredient_count"`
// 	Unit string			`xml:"itemunit"`		`json:"ingredient_unit"`
// }

// type Cake struct {
// 	XMLName xml.Name 	`xml:"cake"`
// 	Name string			`xml:"name"`			`json:"name"`
// 	Time string			`xml:"stovetime"`		`json:"time"`
// 	Ingredients []Ingredient	`xml:"ingredients"`		`json:"ingredients"`
// }			

// type RecipeBook struct {
// 	Recipes []Cake		`xml:"recipes"`			`json:"cake"`
// }

// type Item struct {
// 	// XMLName xml.Name 	`xml:"item"`
// 	Name string			`xml:"itemname"`
// 	Count string		`xml:"itemcount"`
// 	Unit string			`xml:"itemunit"`
// }

// type Cake struct {
// 	// XMLName xml.Name 	`xml:"cake"`
// 	Name string			`xml:"name"`
// 	Time string			`xml:"stovetime"`
// 	Ingredients []Item	`xml:"ingredients"`
// }

// type RecipeBook struct {
// 	XMLName xml.Name 	`xml:"recipes"`
// 	Recipes []Cake		`xml:"cake"`
// }

func (rb *RecipeBook) WriteToJSON() []byte {
	json_data, err := json.Marshal(rb)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	err = ioutil.WriteFile("data.json", json_data, 0644)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	return json_data
}

func (rb *RecipeBook) WriteToXML() []byte {
	xml_data, err := xml.Marshal(rb)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	err = ioutil.WriteFile("data.xml", xml_data, 0644)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	return xml_data
}

// type Ingredient struct {
// 	Name string			`json:"ingredient_name"`
// 	Count string		`json:"ingredient_count"`
// 	Unit string			`json:"ingredient_unit"`
// }

// type Cake struct {
// 	Name string			`json:"name"`
// 	Time string			`json:"time"`
// 	Ingredients []Ingredient	`json:"ingredients"`
// }			

// type RecipeBook struct {
// 	Recipes []Cake		`json:"cake"`
// }

// type Ingredient struct {
// 	Name string
// 	Count string
// 	Unit string
// }

// type Cake struct {
// 	Name string
// 	Time string
// 	Ingredients []Ingredient
// }			

// type RecipeBook struct {
// 	Recipes []Cake
// }


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

type Baker1 struct {
	BookName string
}

func (b1 *Baker1) Read(Path string) RecipeBook {
	var origin RecipeBook	//xml
	err := xml.Unmarshal(ReadFile(Path), &origin);
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	fmt.Println(origin)
	return origin
}

type Baker2 struct {
	BookName string
}

func (b2 *Baker2) Read(Path string) RecipeBook {
	var stolen RecipeBook	//json
	err := json.Unmarshal(ReadFile(Path), &stolen)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	fmt.Println(stolen)
	return stolen
}


type DBReader interface {
	Read(Path string) RecipeBook
}


func main() {
	if len(os.Args) < 3 {
        fmt.Println("no file specified")
		os.Exit(0)
    }
	filePath := os.Args[2]

	var reader DBReader 
	fileExtension := filepath.Ext(filePath)
	switch fileExtension {
	case ".xml":
		reader = &Baker1 {BookName: filePath}
	case ".json":
		reader = &Baker2 {BookName: filePath}
	default:
		fmt.Println("Check file Extension")
		os.Exit(0)
	}

	recipeBook := reader.Read(filePath)
	fmt.Println(recipeBook)

	resJSON := recipeBook.WriteToJSON()
	fmt.Println(string(resJSON))

	resXML := recipeBook.WriteToXML()
	fmt.Println(string(resXML))

}
