package operations

import (
	"fmt"
	"sort"

	s "github.com/gkiryaziev/go-wordlist-cleaner/service"
)

// DoSorting read source file, sort it alphabetically and write to new file
func DoSorting(srcFile, newFile string) error {

	total, err := s.CalculateLines(srcFile)
	if err != nil {
		return err
	}

	// read file
	fmt.Println("\nReading...")
	source, err := s.ReadLine(srcFile)
	if err != nil {
		return err
	}

	// sorting
	fmt.Println("Sorting...")
	sort.Strings(source)

	// write file
	fmt.Println("Saving...")
	err = s.WriteLine(source, newFile)
	if err != nil {
		return err
	}

	fmt.Println("\nResult:", srcFile)
	fmt.Println("-------------------------------------------")
	fmt.Printf("|%-20s|%20d|\n", "Total", total)
	fmt.Println("-------------------------------------------")

	return nil
}
