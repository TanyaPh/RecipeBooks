package DBReader

import (
	"fmt"
	"os"
	"path/filepath"
)

type DBReader interface {
	Read() RecipeBook
}

func New(filePath string) DBReader {
	fileExtension := filepath.Ext(filePath)
	switch fileExtension {
	case ".xml":
		return &readFromXML{BookName: filePath}
	case ".json":
		return &readFromJSON{BookName: filePath}
	default:
		fmt.Println("Check file Extension")
		os.Exit(0)
	}
	return nil
}
