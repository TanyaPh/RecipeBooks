package DBReader

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Item struct {
	Name  string `json:"ingredient_name" xml:"itemname"`
	Count string `json:"ingredient_count" xml:"itemcount"`
	Unit  string `json:"ingredient_unit" xml:"itemunit"`
}

type Cake struct {
	Name        string `json:"name" xml:"name"`
	Time        string `json:"time" xml:"stovetime"`
	Ingredients []Item `json:"ingredients" xml:"ingredients>item"`
}

type RecipeBook struct {
	Recipes []Cake `json:"cake" xml:"cake"`
}

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
