package FSComparator

import (
	"bufio"
	"fmt"
	"os"
)

func ReadOldTXT(filePath string, data map[string]bool) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data[scanner.Text()] = false
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}

func CheckNewTXT(filePath string, data map[string]bool) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if _, ok := data[scanner.Text()]; ok {
			data[scanner.Text()] = true
		} else {
			fmt.Printf("ADDED %s\n", scanner.Text())
		}
	}

	for k, v := range data {
		if !v {
			fmt.Printf("REMOVED %s\n", k)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
