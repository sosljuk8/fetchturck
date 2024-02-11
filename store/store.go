package store

import (
	"fmt"
	"os"
	"strings"
)

func SaveInFile(c string, n string) (bool, error) {

	// split url by "/" and get last part
	urlParts := strings.Split(n, "/")
	filename := urlParts[len(urlParts)-1]

	// write output to file
	file, err := os.Create("parse/pages/turck/" + filename + ".html")
	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
		return false, err
	}
	defer file.Close()
	file.WriteString(c)

	return true, nil
}
