package main

import (
	"day01/internal/FSComparator"
	"flag"
)

// func ReadOldTXT(filePath string, data map[string]bool) {
// 	// file, err := os.ReadFile(filePath)
// 	// if err != nil {
// 	// 	fmt.Println(err)
// 	// 	return
// 	// }
// 	// fmt.Println(string(file))
// 	file, err := os.Open(filePath)
//     if err != nil {
//         fmt.Println(err)
//         return
//     }
//     defer file.Close()

//     scanner := bufio.NewScanner(file)

//     for scanner.Scan() {
//         data[scanner.Text()] = false
//     }

//     if err := scanner.Err(); err != nil {
//         fmt.Println(err)
//     }
// }

// func CheckNewTXT(filePath string, data map[string]bool) {
// 	file, err := os.Open(filePath)
//     if err != nil {
//         fmt.Println(err)
//         return
//     }
//     defer file.Close()

//     scanner := bufio.NewScanner(file)

//     for scanner.Scan() {
//         // fmt.Println(scanner.Text())
// 		if _, ok := data[scanner.Text()]; ok {
// 			data[scanner.Text()] = true
// 		} else {
// 			fmt.Printf("ADDED %s\n", scanner.Text())
// 		}
//     }

// 	for k, v := range data {
// 		if !v {
// 			fmt.Printf("REMOVED %s\n", k)
// 		}
// 	}

//     if err := scanner.Err(); err != nil {
//         fmt.Println(err)
//     }
// }

func main() {
	data := make(map[string]bool)
	var oldFile, newFile string
	flag.StringVar(&oldFile, "old", "", "old snapshot.txt")
	flag.StringVar(&newFile, "new", "", "new snapshot.txt")
	flag.Parse()

	FSComparator.ReadOldTXT(oldFile, data)
	FSComparator.CheckNewTXT(newFile, data)
}
